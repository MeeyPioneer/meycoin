/**
 *  @file
 *  @copyright defined in meycoin/LICENSE.txt
 */

package cmd

import (
	"encoding/json"
	"errors"
	"github.com/meeypioneer/meycoin/cmd/meycoincli/util"
	"github.com/meeypioneer/meycoin/types"
	"github.com/spf13/cobra"
)

var stakeCmd = &cobra.Command{
	Use:    "stake",
	Short:  "Stake balance to meycoin system",
	RunE:   execStake,
	PreRun: connectMeyCoin,
}

func execStake(cmd *cobra.Command, args []string) error {
	return sendStake(cmd, true)
}

var unstakeCmd = &cobra.Command{
	Use:    "unstake",
	Short:  "Unstake balance from meycoin system",
	RunE:   execUnstake,
	PreRun: connectMeyCoin,
}

func execUnstake(cmd *cobra.Command, args []string) error {
	return sendStake(cmd, false)
}

func sendStake(cmd *cobra.Command, s bool) error {
	account, err := types.DecodeAddress(address)
	if err != nil {
		return errors.New("Failed to parse --address flag (" + address + ")\n" + err.Error())
	}
	var ci types.CallInfo
	if s {
		ci.Name = types.Opstake.Cmd()
	} else {
		ci.Name = types.Opunstake.Cmd()
	}
	amountBigInt, err := util.ParseUnit(amount)
	if err != nil {
		return errors.New("Failed to parse --amount flag\n" + err.Error())
	}
	payload, err := json.Marshal(ci)
	if err != nil {
		cmd.Printf("Failed: %s\n", err.Error())
		return nil
	}
	tx := &types.Tx{
		Body: &types.TxBody{
			Account:   account,
			Recipient: []byte(types.MeyCoinSystem),
			Amount:    amountBigInt.Bytes(),
			Payload:   payload,
			GasLimit:  0,
			Type:      types.TxType_GOVERNANCE,
		},
	}

	cmd.Println(sendTX(cmd, tx, account))
	return nil
}
