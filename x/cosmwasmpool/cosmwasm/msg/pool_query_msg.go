// This file defines requests and responses for querying
// the cosmwasm pool contract from the cosmwasm pool model.
package msg

import sdk "github.com/cosmos/cosmos-sdk/types"

// GetSwapFeeQueryMsg
type GetSwapFeeQueryMsg struct {
	GetSwapFee EmptyStruct `json:"get_swap_fee"`
}

type GetSwapFeeQueryMsgResponse struct {
	SwapFee sdk.Dec `json:"swap_fee"`
}

// GetExitFeeQueryMsg
type GetExitFeeQueryMsg struct {
	GetExitFee EmptyStruct `json:"get_exit_fee"`
}

type GetExitFeeQueryMsgResponse struct {
	ExitFee sdk.Dec `json:"exit_fee"`
}

// SpotPrice
type SpotPrice struct {
	QuoteAssetDenom string `json:"quote_asset_denom"`
	BaseAssetDenom  string `json:"base_asset_denom"`
}

type SpotPriceQueryMsg struct {
	SpotPrice SpotPrice `json:"spot_price"`
}

type SpotPriceQueryMsgResponse struct {
	SpotPrice string `json:"spot_price"`
}

// GetTotalPoolLiquidityQueryMsg
type GetTotalPoolLiquidityQueryMsg struct {
	GetTotalPoolLiquidity EmptyStruct `json:"get_total_pool_liquidity"`
}

type GetTotalPoolLiquidityQueryMsgResponse struct {
	TotalPoolLiquidity sdk.Coins `json:"total_pool_liquidity"`
}
