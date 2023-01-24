package keeper

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	cl "github.com/osmosis-labs/osmosis/v14/x/concentrated-liquidity"
	cltypes "github.com/osmosis-labs/osmosis/v14/x/concentrated-liquidity/types"
	gammtypes "github.com/osmosis-labs/osmosis/v14/x/gamm/types"
	"github.com/osmosis-labs/osmosis/v14/x/lockup/types"
)

// Returns a list of newly created lockIDs, or an error.
func (k Keeper) Migrate(ctx sdk.Context, sender sdk.AccAddress, poolId uint64, lockId uint64, sharesToMigrate sdk.Coin, poolIdEntering uint64) (amount0, amount1 sdk.Int, liquidity sdk.Dec, poolIdLeaving uint64, err error) {
	// Steps for unpooling for a (sender, poolID, lockID) triplet.
	// 1) Check if its for a whitelisted unpooling poolID
	// 2) Consistency check that lockID corresponds to sender, and contains correct LP shares. (Should also be validated by caller)
	// 3) Get remaining duration on the lock.
	// 4) If superfluid delegated, superfluid undelegate
	// 5) Break underlying lock. This will clear any metadata if things are superfluid unbonding
	// 6) ExitPool with these unlocked LP shares
	// 7) Make 1 new lock for every asset in collateral. Many code paths need 1 coin / lock assumption to hold
	// 8) Make new lock begin unlocking
	// Get the balancer poolId by parsing the gamm share denom.
	poolIdLeaving, err = getPoolIdFromSharesDenom(sharesToMigrate.Denom)
	if err != nil {
		return sdk.Int{}, sdk.Int{}, sdk.Dec{}, 0, err
	}

	// 1) check if pool is whitelisted for unpool
	migrationInfo := k.gk.GetMigrationInfo(ctx)
	matchFound := false
	for _, info := range migrationInfo.BalancerToConcentratedPoolLinks {
		if info.BalancerPoolId == poolIdLeaving {
			if info.ClPoolId != poolIdEntering {
				return sdk.Int{}, sdk.Int{}, sdk.Dec{}, 0, gammtypes.InvalidPoolMigrationLinkError{PoolIdEntering: poolIdEntering, CanonicalId: info.ClPoolId}
			}
			matchFound = true
			break
		}
	}
	if !matchFound {
		return sdk.Int{}, sdk.Int{}, sdk.Dec{}, 0, gammtypes.PoolMigrationLinkNotFoundError{PoolIdLeaving: poolIdLeaving}
	}

	// 2) Consistency check that lockID corresponds to sender, and contains correct LP shares.
	// These are expected to be true by the caller, but good to double check
	// TODO: Try to minimize dependence on lock here
	lock, err := k.validateLockForUnpool(ctx, sender, poolId, lockId)
	if err != nil {
		return sdk.Int{}, sdk.Int{}, sdk.Dec{}, 0, err
	}
	//gammSharesInLock := lock.Coins[0]

	// 3) Get remaining duration on the lock. Handle if the lock was unbonding.
	//lockRemainingDuration := k.getExistingLockRemainingDuration(ctx, lock)

	// 4) If superfluid delegated, superfluid undelegate
	err = k.unbondSuperfluidIfExists(ctx, sender, lockId)
	if err != nil {
		return sdk.Int{}, sdk.Int{}, sdk.Dec{}, 0, err
	}

	// 5) finish unlocking directly for locked locks
	// this also unlocks locks that were in the unlocking queue
	err = k.lk.ForceUnlock(ctx, *lock)
	if err != nil {
		return sdk.Int{}, sdk.Int{}, sdk.Dec{}, 0, err
	}

	// Get the concentrated pool from the message and type cast it to ConcentratedPoolExtension.
	poolI, err := k.clk.GetPool(ctx, poolIdEntering)
	if err != nil {
		return sdk.Int{}, sdk.Int{}, sdk.Dec{}, 0, err
	}
	concentratedPool, ok := poolI.(cltypes.ConcentratedPoolExtension)
	if !ok {
		// If the conversion fails, return an error.
		return sdk.Int{}, sdk.Int{}, sdk.Dec{}, 0, fmt.Errorf("given pool does not implement ConcentratedPoolExtension, implements %T", poolI)
	}

	// Exit the balancer pool position.
	exitCoins, err := k.gk.ExitPool(ctx, sender, poolIdLeaving, sharesToMigrate.Amount, sdk.NewCoins())
	if err != nil {
		return sdk.Int{}, sdk.Int{}, sdk.Dec{}, 0, err
	}

	// Determine the max and min ticks for the concentrated pool we are migrating to.
	minTick, maxTick := cl.GetMinAndMaxTicksFromExponentAtPriceOne(concentratedPool.GetPrecisionFactorAtPriceOne())

	// Determine remaining lock time and freeze position for that duration.
	freezeUntil := remainingLockTime(ctx, *lock)

	// Create a full range (min to max tick) concentrated liquidity position.
	amount0, amount1, liquidity, err = k.clk.CreatePosition(ctx, poolIdEntering, sender, exitCoins.AmountOf(concentratedPool.GetToken0()), exitCoins.AmountOf(concentratedPool.GetToken1()), sdk.ZeroInt(), sdk.ZeroInt(), minTick, maxTick, freezeUntil)
	if err != nil {
		return sdk.Int{}, sdk.Int{}, sdk.Dec{}, 0, err
	}

	return amount0, amount1, liquidity, poolIdLeaving, nil
}

// getPoolIdFromSharesDenom takes in a string representing a pool share denom and extracts the pool ID.
// It returns the pool ID as a uint64 and an error if the denom is invalid.
func getPoolIdFromSharesDenom(denom string) (uint64, error) {
	if !strings.HasPrefix(denom, "gamm/pool/") {
		return 0, fmt.Errorf("invalid pool share denom %s", denom)
	}
	return strconv.ParseUint(denom[len("gamm/pool/"):], 10, 64)
}

func remainingLockTime(ctx sdk.Context, lock types.PeriodLock) (lockedUntil time.Time) {
	if lock.EndTime.IsZero() {
		return ctx.BlockTime().Add(lock.Duration)
	} else {
		return lock.EndTime
	}
}
