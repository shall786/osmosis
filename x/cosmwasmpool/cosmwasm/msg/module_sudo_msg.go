// This file defines requests and responses for issuing
// sudo messages to the cosmwasm pool contract from the cosmwasm pool module.
package msg

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type EmptyStruct struct{}

// SwapExactAmountIn
type SwapExactAmountIn struct {
	Sender        string   `json:"sender"`
	TokenIn       sdk.Coin `json:"token_in"`
	TokenOutDenom string   `json:"token_out_denom"`
	TokenOutMin   sdk.Int  `json:"token_out_min_amount"`
	SwapFee       sdk.Dec  `json:"swap_fee"`
}

type SwapExactAmountInSudoMsg struct {
	SwapExactAmountIn SwapExactAmountIn `json:"swap_exact_amount_in"`
}

func NewSwapExactAmountInSudoMsg(sender string, tokenIn sdk.Coin, tokenOutDenom string, tokenOutMin sdk.Int, swapFee sdk.Dec) SwapExactAmountInSudoMsg {
	return SwapExactAmountInSudoMsg{
		SwapExactAmountIn: SwapExactAmountIn{
			Sender:        sender,
			TokenIn:       tokenIn,
			TokenOutDenom: tokenOutDenom,
			TokenOutMin:   tokenOutMin,
			SwapFee:       swapFee,
		},
	}
}

type SwapExactAmountInSudoMsgResponse struct {
	TokenOutAmount sdk.Int `json:"token_out_amount"`
}

// SwapExactAmountOut
type SwapExactAmountOutSudoMsg struct {
	SwapExactAmountOut SwapExactAmountOut `json:"swap_exact_amount_out"`
}

type SwapExactAmountOut struct {
	Sender           string   `json:"sender"`
	TokenInDenom     string   `json:"token_in_denom"`
	TokenOut         sdk.Coin `json:"token_out"`
	TokenInMaxAmount sdk.Int  `json:"token_in_max_amount"`
	SwapFee          sdk.Dec  `json:"swap_fee"`
}

func NewSwapExactAmountOutSudoMsg(sender string, tokenInDenom string, tokenOut sdk.Coin, tokenInMaxAmount sdk.Int, swapFee sdk.Dec) SwapExactAmountOutSudoMsg {
	return SwapExactAmountOutSudoMsg{
		SwapExactAmountOut: SwapExactAmountOut{
			Sender:           sender,
			TokenInDenom:     tokenInDenom,
			TokenOut:         tokenOut,
			TokenInMaxAmount: tokenInMaxAmount,
			SwapFee:          swapFee,
		},
	}
}

type SwapExactAmountOutSudoMsgResponse struct {
	TokenInAmount sdk.Int `json:"token_in_amount"`
}
