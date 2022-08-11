package services

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"math/big"
	"os"
	"rarity-backend/constants"
	"rarity-backend/dlt"
	"rarity-backend/handlers"
	"rarity-backend/structs"
	"sync"
)

// collectEvents sends request using the ethereum client for events emitted from the Faces contract. It iterates over the events and filters for mint and morph events.
// Events iteration is implemented concurrently.
// Returns last processed block, so it can be persisted in the database after the events have been fully processed.
func collectEvents(ethClient *dlt.EthereumClient, address string, polymorphDBName string, elm *structs.EventLogsMutex, wg *sync.WaitGroup) (uint64, error) {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	var lastProcessedBlockNumber, lastChainBlockNumberInt64 int64

	lastProcessedBlockNumber, _ = handlers.GetLastProcessedBlockNumber(polymorphDBName, true)

	latestBlock, err := ethClient.Client.BlockNumber(context.Background())
	if err != nil {
		log.WithFields(log.Fields{"network": "Ethereum", "original error: ": err}).Error("error fetching latest block number. Last chain block number will be the last processed block number")
		lastChainBlockNumberInt64 = lastProcessedBlockNumber
	} else {
		lastChainBlockNumberInt64 = int64(latestBlock)
	}

	// If by any chance, the network returns a block that is less than the last processed, return
	if lastProcessedBlockNumber > lastChainBlockNumberInt64 {
		log.WithFields(log.Fields{"network": "Ethereum"}).Warnf("last process block number [%d] exceeds last chain block [%d]", lastProcessedBlockNumber, lastChainBlockNumberInt64)
		return uint64(lastProcessedBlockNumber), nil
	}

	// If the blocks that have to be processed are more than 1000, process only 1000 blocks.
	// Given that the function is in an endless loop, you will process them slowly, but all.
	if lastChainBlockNumberInt64-lastProcessedBlockNumber > 1000 {
		log.WithFields(log.Fields{"network": "Ethereum"}).Warn("splitting blocks into chunks of 1000")
		lastChainBlockNumberInt64 = lastProcessedBlockNumber + 1000
	}

	ethLogs, err := ethClient.Client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(lastProcessedBlockNumber),
		ToBlock:   big.NewInt(lastChainBlockNumberInt64),
		Addresses: []common.Address{common.HexToAddress(address)},
	})
	if err != nil {
		log.WithFields(log.Fields{"network": "Ethereum", "original error: ": err}).Error("error filtering logs")
		return uint64(lastProcessedBlockNumber), err
	}
	log.Infof("Processing blocks [%v] - [%v] for faces events", lastProcessedBlockNumber, lastChainBlockNumberInt64)
	wg.Add(1)
	go saveToEventLogMutex(ethLogs, elm, wg)
	wg.Wait()
	return uint64(lastChainBlockNumberInt64) + 1, nil
}

func collectAndRetrieveEventsPolygon(polygonClient *dlt.EthereumClient, facesDBName string) (uint64, []types.Log, error) {
	var lastProcessedBlockNumber, lastChainBlockNumberInt64 int64

	lastProcessedBlockNumber, _ = handlers.GetLastProcessedBlockNumber(facesDBName, false)
	latestBlock, err := polygonClient.Client.BlockNumber(context.Background())
	if err != nil {
		log.Println("Error fetching latest block number, ", err)
		lastChainBlockNumberInt64 = lastProcessedBlockNumber
	} else {
		lastChainBlockNumberInt64 = int64(latestBlock)
	}

	// If by any chance, the network returns a block that is less than the last processed, return
	if lastProcessedBlockNumber > lastChainBlockNumberInt64 {
		log.Printf("Last process block number [%d] exceeds last chain block [%d]", lastProcessedBlockNumber, lastChainBlockNumberInt64)
		return uint64(lastProcessedBlockNumber), nil, nil
	}

	// If the blocks that have to be processed are more than 1000, process only 1000 blocks.
	// Given that the function is in an endless loop, you will process them slowly, but all.
	if lastChainBlockNumberInt64-lastProcessedBlockNumber > 1000 {
		log.Println("Splitting blocks into chunks of 1000")
		lastChainBlockNumberInt64 = lastProcessedBlockNumber + 1000
	}

	facesPolygonAddress := os.Getenv("CONTRACT_ADDRESS_POLYGON")

	log.Printf("Processing blocks [%v] - [%v] for faces events", lastProcessedBlockNumber, lastChainBlockNumberInt64)

	currentBatchLogs, err := polygonClient.Client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(lastProcessedBlockNumber),
		ToBlock:   big.NewInt(lastChainBlockNumberInt64),
		Addresses: []common.Address{common.HexToAddress(facesPolygonAddress)},
		Topics:    [][]common.Hash{{common.HexToHash(constants.MintEvent.Signature), common.HexToHash(constants.TransferEvent.Signature), common.HexToHash(constants.MorphEvent.Signature)}},
	})
	if err != nil {
		log.Println("Error filtering logs for this current batch, ", err)
		return uint64(lastProcessedBlockNumber), nil, err
	}
	return uint64(lastChainBlockNumberInt64) + 1, currentBatchLogs, nil
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
		case constants.MintEvent.Signature, constants.MorphEvent.Signature, constants.TransferEvent.Signature:
			elm.EventLogs = append(elm.EventLogs, ethLog)
		}
	}
	elm.Mutex.Unlock()
}
