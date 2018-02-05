package eth

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/log"
)

// BlueMetadata collects info during geth sync
type BlueMetadata struct {
	state string
	eth   *Ethereum
}

var instance *BlueMetadata

// SetBlueEth notifies us of the Ethereum object instance being used so we can log from it here
func SetBlueEth(eth *Ethereum) {
	if instance == nil {
		instance = &BlueMetadata{}
	}
	return
	headBlockHash := core.GetHeadBlockHash(eth.chainDb)
	headBlockNum := core.GetBlockNumber(eth.chainDb, headBlockHash)

	genesisBlock := eth.blockchain.GetBlockByNumber(0)
	log.Info("BLUE: Got Eth instance", "Head block hash", headBlockHash.String(), "Genesis", genesisBlock.Hash().String())
	expectedGenesisBlockHash := "0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3"
	if genesisBlock.Hash().String() != expectedGenesisBlockHash {
		log.Warn("You do not appear to be on the main Ethereum chain")
	}
	(*instance).eth = eth
	(*instance).state = "LISTENING"

	blueDeployedTx := common.HexToHash("0x451e35d6a5639960ff3df90d7e9214649855571f149d0d0be4545d8b1924e245")
	receipt, hash, _, _ := core.GetReceipt(eth.chainDb, blueDeployedTx)
	log.Info("BLUE: Got BLUE token contract info", "Deployed address", receipt.ContractAddress.String(), "Hash", hash)

	// for i := 0; i < headBlockNum; i++ {
	// eth.ApiBackend.
	// api := NewPublicBlockchainAPI(eth.ApiBackend)
	// ethapi := NewPublicEthereumAPI(eth)

	expiry := time.Now().Add(10)
	ctx, _ := context.WithDeadline(context.Background(), expiry)
	evm, errFunc, err := eth.ApiBackend.GetEVM(ctx, nil, nil, nil, (*&vm.Config{}))
	if err != nil {
		log.Warn("BLUE Error", "error", err.Error())
	}
	if errFunc != nil {
		log.Warn("BLUE ErrorFunc")
		errFunc().Error()
	}
	log.Info("BLUE Got EVM", "EVM", evm)
	// ethapi.

	log.Info("BLUE Data", "head block num", headBlockNum)

	// }
}

// TODO: Emulate contract interactions in fuzz testing
// Example: github.com/bluecrypto/go-ethereum/internal/ethapi/api.go
// Line 695

func (*BlueMetadata) ProcessBlockReceipts(receipts types.Receipts) {
	for i := 0; i < len(receipts); i++ {
		receipt := receipts[0]
		contractAddress := receipt.ContractAddress
		log.Info("BLUE: Receipt info found", "Address", contractAddress.String())

		log.Info(receipt.String())
	}
}
