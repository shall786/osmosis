package cli

import (
	flag "github.com/spf13/pflag"

	"github.com/spf13/cobra"

	"github.com/osmosis-labs/osmosis/osmoutils/osmocli"
	"github.com/osmosis-labs/osmosis/v15/x/poolmanager/types"
)

func NewTxCmd() *cobra.Command {
	txCmd := osmocli.TxIndexCmd(types.ModuleName)

	osmocli.AddTxCmd(txCmd, NewGeneratedCmd)

	return txCmd
}

func NewGeneratedCmd() (*osmocli.TxCliDesc, *types.MsgSwapExactAmountIn) {
	return &osmocli.TxCliDesc{
		Use:   "generated-cmd [token-in] [token-out-min-amount]",
		Short: "TODO: code-generated command, change this",
		Flags: osmocli.FlagDesc{RequiredFlags: []*flag.FlagSet{FlagCodeGenerated()}},
	}, &types.MsgSwapExactAmountIn{}
}
