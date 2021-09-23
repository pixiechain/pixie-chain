package sysc

import (
	"math/big"

	"github.com/ethereum/go-ethereum/consensus/congress/sysc/sidechainproposal"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

type IUpgradeAction interface {
	GetName() string
	Upgrade(config *params.ChainConfig, height *big.Int, state *state.StateDB) error
	Init(state *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) error
	Execute(state *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) error
}

var (
	sysContracts []IUpgradeAction
)

func init() {
	sysContracts = []IUpgradeAction{
		&sidechainproposal.SidechainProposal{},
	}
}

func UpgradeSystemContracts(state *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (err error) {
	if config == nil || header == nil || state == nil {
		return
	}
	height := header.Number

	for _, contract := range sysContracts {
		log.Info("Upgrade system contract", "name", contract.GetName(), "height", height, "chainId", config.ChainID.String())
		if err = contract.Upgrade(config, height, state); err != nil {
			log.Error("Upgrade system contract error", "name", contract.GetName(), "err", err)
			return
		}

		log.Info("Init system contract", "name", contract.GetName(), "height", header.Number, "chainId", config.ChainID.String())
		if err = contract.Init(state, header, chainContext, config); err != nil {
			log.Error("Init system contract error", "name", contract.GetName(), "err", err)
			return
		}

		log.Info("System contract upgrade execution", "name", contract.GetName(), "height", header.Number, "chainId", config.ChainID.String())
		if err = contract.Execute(state, header, chainContext, config); err != nil {
			log.Error("System contract execution error", "name", contract.GetName(), "err", err)
			return
		}
	}

	// Update the state with pending changes
	state.Finalise(true)

	return
}
