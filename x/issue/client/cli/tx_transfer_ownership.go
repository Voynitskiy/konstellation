package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

var _ = strconv.Itoa(0)

func CmdTransferOwnership() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-ownership [to_address] [denom]",
		Short: "Broadcast message TransferOwnership",
		Long:  "Transfers token from one owner to another",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			toAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			denom := args[1]

			msg := types.NewMsgTransferOwnership(clientCtx.GetFromAddress(), toAddr, denom)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return validateErr
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
