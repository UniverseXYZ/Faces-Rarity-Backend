package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"rarity-backend/config"
	"rarity-backend/dlt"
	"rarity-backend/handlers"
	"rarity-backend/services"
	"rarity-backend/store"
	"rarity-backend/structs"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

func connectToEthereum() *dlt.EthereumClient {

	nodeURL := os.Getenv("NODE_URL")

	client, err := dlt.NewEthereumClient(nodeURL)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to ethereum client")

	return client
}

// initResources is a wrapper function which tries to initialize all .env variables, contract abi, new contract instance.
//
// It connects to the ethereum client and returns all information which will be needed at some point from the application
func initResources() (*dlt.EthereumClient, abi.ABI, *store.Store, string, *structs.ConfigService, structs.DBInfo) {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	// Inital step: Recover to be up to date
	ethClient := connectToEthereum()
	FacesDBName := os.Getenv("FACES_DB")
	rarityCollectionName := os.Getenv("RARITY_COLLECTION")
	blocksCollectionName := os.Getenv("BLOCKS_COLLECTION")
	contractAddress := os.Getenv("CONTRACT_ADDRESS")
	transactionsCollectionName := os.Getenv("TRANSACTIONS_COLLECTION")
	historyCollectionName := os.Getenv("HISTORY_COLLECTION")
	morphCostCollectionName := os.Getenv("MORPH_COST_COLLECTION")

	if contractAddress == "" {
		log.Fatal("Missing contract address in .env")
	}
	if FacesDBName == "" {
		log.Fatal("Missing faces db name in .env")
	}
	if rarityCollectionName == "" {
		log.Fatal("Missing rarity collection name in .env")
	}
	if blocksCollectionName == "" {
		log.Fatal("Missing block collection name in .env")
	}
	if transactionsCollectionName == "" {
		log.Fatal("Missing transactions collection name in .env")
	}
	if historyCollectionName == "" {
		log.Fatal("Missing morph history collection name in .env")
	}
	if morphCostCollectionName == "" {
		log.Fatal("Missing morph cost collection name in .env")
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(store.PolymorphicFacesRootMetaData.ABI)))
	if err != nil {
		log.Fatal(err)
	}

	instance, err := store.NewStore(common.HexToAddress(contractAddress), ethClient.Client)
	if err != nil {
		log.Fatalln(err)
	}

	configService := config.NewConfigService("./config.json")
	dbInfo := structs.DBInfo{
		PolymorphDBName:            FacesDBName,
		RarityCollectionName:       rarityCollectionName,
		TransactionsCollectionName: transactionsCollectionName,
		BlocksCollectionName:       blocksCollectionName,
		HistoryCollectionName:      historyCollectionName,
		MorphCostCollectionName:    morphCostCollectionName,
	}
	return ethClient, contractAbi, instance, contractAddress, configService, dbInfo
}

// main is the entry point of the application.
// It fetches all configurations and starts 2 concurrent processes:
//
// 1. API which handles GET requests
//
// 2. Polling process which processes mint and morph events and stores their metadata in the database
func main() {
	ethClient,
		contractAbi,
		instance,
		contractAddress,
		configService,
		dbInfo := initResources()

	go recoverAndPoll(
		ethClient,
		contractAbi,
		instance,
		contractAddress,
		configService,
		dbInfo)

	startAPI()
}

// startAPI registers the endpoints for API and listens for requests
// API has moved to a cloud function due to bad response times
func startAPI() {
	// Routine two: API -> Should start after deploy?
	app := fiber.New()
	// app.Get("/morphs/", handlers.GetPolymorphs)
	// app.Get("/morphs/:id", handlers.GetPolymorphById)
	// app.Get("/morphs/history/:id", handlers.GetPolymorphHistory)
	apiPORT := os.Getenv("API_PORT")
	log.Fatal(app.Listen(apiPORT))
}

// recoverAndPoll loads transactions and morph cost state in memory from the database and initiates polling mechanism.
//
// Recovery function and polling function is the same.
// Currently the polling timer doesn't wait for the previous one to finish before starting the new countdown
func recoverAndPoll(ethClient *dlt.EthereumClient, contractAbi abi.ABI, store *store.Store, contractAddress string, configService *structs.ConfigService, dbInfo structs.DBInfo) {
	// Build transactions scramble transaction mapping from db
	txMap, err := handlers.GetTransactionsMapping(dbInfo.PolymorphDBName, dbInfo.TransactionsCollectionName)
	if err != nil {
		log.Println("Error getting transactions mapping initially.")
	}

	// Build polymorph cost mapping from db
	morphCostMap, err := handlers.GetMorphPriceMapping(dbInfo.PolymorphDBName, dbInfo.HistoryCollectionName)
	if err != nil {
		log.Println("Error getting morph prices mapping initially.")
	}
	// Recover immediately
	// services.RecoverProcess(ethClient, contractAbi, store, contractAddress, configService, dbInfo, txMap, morphCostMap)
	// Routine one: Start polling after recovery

	for {
		err = services.RecoverProcess(ethClient, contractAbi, store, contractAddress, configService, dbInfo, txMap, morphCostMap)
		if err != nil {
			fmt.Println("Recovering from error... ", err)
			time.Sleep(15 * time.Second)
			continue
		} else {
			time.Sleep(15 * time.Second)
		}
	}
}
