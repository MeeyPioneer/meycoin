/**
 *  @file
 *  @copyright defined in meycoin/LICENSE.txt
 */

package chain

import (
	"math/big"

	"github.com/meeypioneer/meycoin/consensus"

	"github.com/meeypioneer/meycoin/contract"
	"github.com/meeypioneer/meycoin/contract/enterprise"
	"github.com/meeypioneer/meycoin/contract/name"
	"github.com/meeypioneer/meycoin/contract/system"
	"github.com/meeypioneer/meycoin/state"
	"github.com/meeypioneer/meycoin/types"
)

func executeGovernanceTx(ccc consensus.ChainConsensusCluster, bs *state.BlockState, txBody *types.TxBody, sender, receiver *state.V,
	blockInfo *types.BlockHeaderInfo) ([]*types.Event, error) {

	if len(txBody.Payload) <= 0 {
		return nil, types.ErrTxFormatInvalid
	}

	governance := string(txBody.Recipient)

	scs, err := bs.StateDB.OpenContractState(receiver.AccountID(), receiver.State())
	if err != nil {
		return nil, err
	}
	blockNo := blockInfo.No
	var events []*types.Event
	switch governance {
	case types.MeyCoinSystem:
		events, err = system.ExecuteSystemTx(scs, txBody, sender, receiver, blockInfo)
	case types.MeyCoinName:
		events, err = name.ExecuteNameTx(bs, scs, txBody, sender, receiver, blockInfo)
	case types.MeyCoinEnterprise:
		events, err = enterprise.ExecuteEnterpriseTx(bs, ccc, scs, txBody, sender, receiver, blockNo)
		if err != nil {
			err = contract.NewGovEntErr(err)
		}
	default:
		logger.Warn().Str("governance", governance).Msg("receive unknown recipient")
		err = types.ErrTxInvalidRecipient
	}
	if err == nil {
		err = bs.StateDB.StageContractState(scs)
	}

	return events, err
}

// InitGenesisBPs opens system contract and put initial voting result
// it also set *State in Genesis to use statedb
func InitGenesisBPs(states *state.StateDB, genesis *types.Genesis) error {
	aid := types.ToAccountID([]byte(types.MeyCoinSystem))
	scs, err := states.OpenContractStateAccount(aid)
	if err != nil {
		return err
	}

	voteResult := make(map[string]*big.Int)
	for _, v := range genesis.BPs {
		voteResult[v] = new(big.Int).SetUint64(0)
	}
	if err = system.InitVoteResult(scs, voteResult); err != nil {
		return err
	}

	// Set genesis.BPs to the votes-ordered BPs. This will be used later for
	// bootstrapping.
	genesis.BPs = system.BuildOrderedCandidates(voteResult)
	if err = states.StageContractState(scs); err != nil {
		return err
	}
	if err = states.Update(); err != nil {
		return err
	}
	if err = states.Commit(); err != nil {
		return err
	}

	return nil
}
