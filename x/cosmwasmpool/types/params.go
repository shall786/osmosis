package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys.
var (
	KeyParamField = []byte("TODO: CHANGE ME")
)

// ParamTable for gamm module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(poolCreationFee sdk.Coins) Params {
	return Params{}
}

// DefaultParams are the default cosmwasmpool module parameters.
func DefaultParams() Params {
	return Params{}
}

// Vlidate validates params.
func (p Params) Validate() error {
	return nil
}

// Implements params.ParamSet.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		// paramtypes.NewParamSetPair(KeyParamField, &p.Field, validateFn),
	}
}
