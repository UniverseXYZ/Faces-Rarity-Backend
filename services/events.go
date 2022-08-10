package services

import (
	"context"
	"log"
	"math/big"
	"rarity-backend/constants"
	"rarity-backend/dlt"
	"rarity-backend/handlers"
	"rarity-backend/store"
	"rarity-backend/structs"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// collectEvents sends request using the ethereum client for events emitted from the Polymorphs contract. It iterates over the events and filters for mint and morph events.
// Events iteration is implemented concurrently.
// Returns last processed block so it can be persisted in the database after the events have been fully processed.
//
// If events in the block range is > 10,000 the range is split in two and the function is called recursively until the blocks range can be processed.(10,000 limit: https://infura.io/docs/ethereum/json-rpc/eth_getLogs)
func collectEvents(ethClient *dlt.EthereumClient, contractAbi abi.ABI, instance *store.Store, address string, configService *structs.ConfigService, polymorphDBName string, rarityCollectionName string, blocksCollectionName string, startBlock int64, endBlock int64, wg *sync.WaitGroup, elm *structs.EventLogsMutex) (uint64, error) {
	var lastProcessedBlockNumber, lastChainBlockNumberInt64 int64

	if startBlock != 0 {
		lastProcessedBlockNumber = startBlock
	} else {
		lastProcessedBlockNumber, _ = handlers.GetLastProcessedBlockNumber(polymorphDBName, blocksCollectionName)
	}

	if endBlock != 0 {
		lastChainBlockNumberInt64 = endBlock
	} else {
		latestBlock, err := ethClient.Client.BlockNumber(context.Background())
		if err != nil {
			log.Println("Error fetching latest block number, ", err)
			lastChainBlockNumberInt64 = int64(lastProcessedBlockNumber)
		} else {
			lastChainBlockNumberInt64 = int64(latestBlock)
		}
	}

	// If by any chance, the network returns a block that is less than the last processed, return
	if lastProcessedBlockNumber > lastChainBlockNumberInt64 {
		log.Printf("Last process block number [%d] exceeds last chain block [%d]", lastProcessedBlockNumber, lastChainBlockNumberInt64)
		return uint64(lastProcessedBlockNumber), nil
	}

	// If the blocks that have to be processed are more than 1000, process only 1000 blocks.
	// Given that the function is in an endless loop, you will process them slowly, but all.
	if lastChainBlockNumberInt64-lastProcessedBlockNumber > 1000 {
		log.Println("Splitting blocks into chunks of 1000")
		lastChainBlockNumberInt64 = lastProcessedBlockNumber + 1000
	}

	ethLogs, err := ethClient.Client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(lastProcessedBlockNumber),
		ToBlock:   big.NewInt(lastChainBlockNumberInt64),
		Addresses: []common.Address{common.HexToAddress(address)},
	})
	if err != nil {
		log.Println("Error filtering logs, ", err)
		return uint64(lastProcessedBlockNumber), err
	} else {
		log.Printf("Processing blocks [%v] - [%v] for polymorph events", lastProcessedBlockNumber, lastChainBlockNumberInt64)
		wg.Add(1)
		go saveToEventLogMutex(ethLogs, elm, wg)
	}
	wg.Wait()
	return uint64(lastChainBlockNumberInt64) + 1, nil
}

// saveToEventLogMutex concurrently saves mint and morph events an array which will be processed after all events have been filtered for these events.
//
// Uses Mutex and WaitGroup to prevent race conditions
func saveToEventLogMutex(ethLogs []types.Log, elm *structs.EventLogsMutex, wg *sync.WaitGroup) {
	defer wg.Done()
	elm.Mutex.Lock()
	for _, ethLog := range ethLogs {
		eventSig := ethLog.Topics[0].String()
		switch eventSig {
		case constants.MintEvent.Signature, constants.MorphEvent.Signature:
			elm.EventLogs = append(elm.EventLogs, ethLog)
		}
	}
	elm.Mutex.Unlock()
}
