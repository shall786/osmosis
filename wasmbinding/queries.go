package wasmbinding

import (
	"fmt"
	"time"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/osmosis-labs/osmosis/v12/wasmbinding/bindings"
	gammkeeper "github.com/osmosis-labs/osmosis/v12/x/gamm/keeper"
	gammtypes "github.com/osmosis-labs/osmosis/v12/x/gamm/types"
	tokenfactorykeeper "github.com/osmosis-labs/osmosis/v12/x/tokenfactory/keeper"
	twapkeeper "github.com/osmosis-labs/osmosis/v12/x/twap"
)

type QueryPlugin struct {
	gammKeeper         *gammkeeper.Keeper
	twapKeeper         *twapkeeper.Keeper
	tokenFactoryKeeper *tokenfactorykeeper.Keeper
}

// NewQueryPlugin returns a reference to a new QueryPlugin.
func NewQueryPlugin(gk *gammkeeper.Keeper, tk *twapkeeper.Keeper, tfk *tokenfactorykeeper.Keeper) *QueryPlugin {
	return &QueryPlugin{
		gammKeeper:         gk,
		twapKeeper:         tk,
		tokenFactoryKeeper: tfk,
	}
}

// GetDenomAdmin is a query to get denom admin.
func (qp QueryPlugin) GetDenomAdmin(ctx sdk.Context, denom string) (*bindings.DenomAdminResponse, error) {
	metadata, err := qp.tokenFactoryKeeper.GetAuthorityMetadata(ctx, denom)
	if err != nil {
		return nil, fmt.Errorf("failed to get admin for denom: %s", denom)
	}

	return &bindings.DenomAdminResponse{Admin: metadata.Admin}, nil
}

// GetPoolState is a query to get pool liquidity and amount of each denoms' pool shares.
func (qp QueryPlugin) GetPoolState(ctx sdk.Context, poolID uint64) (*bindings.PoolAssets, error) {
	poolData, err := qp.gammKeeper.GetPoolAndPoke(ctx, poolID)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "gamm get pool")
	}

	return &bindings.PoolAssets{
		Assets: poolData.GetTotalPoolLiquidity(ctx),
		Shares: sdk.Coin{
			Denom:  gammtypes.GetPoolShareDenom(poolID),
			Amount: poolData.GetTotalShares(),
		},
	}, nil
}

// GetSpotPrice is a query to get spot price of denoms.
func (qp QueryPlugin) GetSpotPrice(ctx sdk.Context, spotPrice *bindings.SpotPrice) (*sdk.Dec, error) {
	if spotPrice == nil {
		return nil, wasmvmtypes.InvalidRequest{Err: "gamm spot price null"}
	}

	poolId := spotPrice.Swap.PoolId
	denomIn := spotPrice.Swap.DenomIn
	denomOut := spotPrice.Swap.DenomOut
	withSwapFee := spotPrice.WithSwapFee

	price, err := qp.gammKeeper.CalculateSpotPrice(ctx, poolId, denomIn, denomOut)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "gamm get spot price")
	}

	if withSwapFee {
		poolData, err := qp.gammKeeper.GetPoolAndPoke(ctx, poolId)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "gamm get pool")
		}

		price = price.Mul(sdk.OneDec().Sub(poolData.GetSwapFee(ctx)))
	}

	return &price, nil
}

// EstimateSwap validates each denom (in / out) and performs a swap.
func (qp QueryPlugin) EstimateSwap(ctx sdk.Context, estimateSwap *bindings.EstimateSwap) (*bindings.SwapAmount, error) {
	if estimateSwap == nil {
		return nil, wasmvmtypes.InvalidRequest{Err: "gamm estimate swap null"}
	}
	if err := sdk.ValidateDenom(estimateSwap.First.DenomIn); err != nil {
		return nil, sdkerrors.Wrap(err, "gamm estimate swap denom in")
	}
	if err := sdk.ValidateDenom(estimateSwap.First.DenomOut); err != nil {
		return nil, sdkerrors.Wrap(err, "gamm estimate swap denom out")
	}

	if estimateSwap.Amount == (bindings.SwapAmount{}) {
		return nil, wasmvmtypes.InvalidRequest{Err: "gamm estimate swap empty swap"}
	}

	subCtx, _ := ctx.CacheContext()
	estimate, err := EstimatePerformSwap(qp.gammKeeper, subCtx, estimateSwap.ToSwapMsg())
	return estimate, err
}

// EstimatePerformSwap can be used to either estimate a exact amount in swap, or an exact amount out swap.
// This function requires requires either swap.Amount.ExactIn or swap.Amount.ExactOut to be non nil, or else it will error.
func EstimatePerformSwap(keeper *gammkeeper.Keeper, ctx sdk.Context, swap *bindings.SwapMsg) (*bindings.SwapAmount, error) {
	if swap == nil {
		return nil, wasmvmtypes.InvalidRequest{Err: "gamm perform swap null swap"}
	}
	if swap.Amount.ExactIn != nil {
		routes, tokenIn, tokenOutMinAmount, err := getSwapExactAmountInParams(swap)
		if err != nil {
			return nil, err
		}

		tokenOutAmount, err := keeper.EstimateMultihopSwapExactAmountIn(ctx, routes, tokenIn, tokenOutMinAmount)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "gamm perform swap exact amount in")
		}
		return &bindings.SwapAmount{Out: &tokenOutAmount}, nil
	} else if swap.Amount.ExactOut != nil {
		routes, tokenOut, tokenInMaxAmount, err := getSwapAmountOutParams(swap)
		if err != nil {
			return nil, err
		}

		tokenInAmount, err := keeper.EstimateMultihopSwapExactAmountOut(ctx, routes, tokenInMaxAmount, tokenOut)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "gamm perform swap exact amount out")
		}
		return &bindings.SwapAmount{In: &tokenInAmount}, nil
	} else {
		return nil, wasmvmtypes.UnsupportedRequest{Kind: "must support either Swap.ExactIn or Swap.ExactOut"}
	}
}

func (qp QueryPlugin) ArithmeticTwap(ctx sdk.Context, arithmeticTwap *bindings.ArithmeticTwap) (*sdk.Dec, error) {
	if arithmeticTwap == nil {
		return nil, wasmvmtypes.InvalidRequest{Err: "gamm arithmetic twap null"}
	}

	poolId := arithmeticTwap.PoolId
	quoteAssetDenom := arithmeticTwap.QuoteAssetDenom
	baseAssetDenom := arithmeticTwap.BaseAssetDenom
	startTime := time.UnixMilli(arithmeticTwap.StartTime)
	endTime := time.UnixMilli(arithmeticTwap.EndTime)

	twap, err := qp.twapKeeper.GetArithmeticTwap(ctx, poolId, quoteAssetDenom, baseAssetDenom, startTime, endTime)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "gamm arithmetic twap")
	}

	return &twap, nil
}

func (qp QueryPlugin) ArithmeticTwapToNow(ctx sdk.Context, arithmeticTwap *bindings.ArithmeticTwapToNow) (*sdk.Dec, error) {
	if arithmeticTwap == nil {
		return nil, wasmvmtypes.InvalidRequest{Err: "gamm arithmetic twap null"}
	}

	poolId := arithmeticTwap.PoolId
	quoteAssetDenom := arithmeticTwap.QuoteAssetDenom
	baseAssetDenom := arithmeticTwap.BaseAssetDenom
	startTime := time.UnixMilli(arithmeticTwap.StartTime)

	twap, err := qp.twapKeeper.GetArithmeticTwapToNow(ctx, poolId, quoteAssetDenom, baseAssetDenom, startTime)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "gamm arithmetic twap")
	}

	return &twap, nil
}
