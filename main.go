package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"rarity-backend/config"
	"rarity-backend/dlt"
	"rarity-backend/handlers"
	"rarity-backend/services"
	"rarity-backend/store"
	"rarity-backend/structs"
	"strings"
	"time"
)

func connectToEthereum() *dlt.EthereumClient {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	nodeURL := os.Getenv("NODE_URL_ETHEREUM")

	client, err := dlt.NewEthereumClient(nodeURL)

	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Successfully connected to ethereum client")

	return client
}

func connectToPolygon() (*dlt.EthereumClient, error) {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	nodeURL := os.Getenv("NODE_URL_POLYGON")

	client, err := dlt.NewEthereumClient(nodeURL)

	if err != nil {
		return nil, err
	}

	log.Infof("Successfully connected to polygon client")

	return client, nil
}

// initResources is a wrapper function which tries to initialize all .env variables, contract abi, new contract instance.
//
// It connects to the ethereum client and returns all information which will be needed at some point from the application
func initResources() (*dlt.EthereumClient, *dlt.EthereumClient, abi.ABI, *store.Store, *store.Store, string, string, *structs.ConfigService, structs.DBInfo) {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	// Initial step: Recover to be up-to-date
	ethClient := connectToEthereum()
	polygonClient, _ := connectToPolygon()
	FacesDBName := os.Getenv("FACES_DB")
	rarityCollectionName := os.Getenv("RARITY_COLLECTION")
	blocksCollectionName := os.Getenv("BLOCKS_COLLECTION")
	blocksCollectionNamePolygon := os.Getenv("BLOCKS_COLLECTION_POLYGON")
	contractAddress := os.Getenv("CONTRACT_ADDRESS")
	contractAddressPolygon := os.Getenv("CONTRACT_ADDRESS_POLYGON")
	rootTunnelAddress := os.Getenv("ROOT_TUNNEL_ADDRESS")
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
	if blocksCollectionNamePolygon == "" {
		log.Fatal("Missing block collection name for Polygon in .env")
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

	contractAbi, err := abi.JSON(strings.NewReader(store.PolymorphicFacesRootMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}

	instance, err := store.NewStore(common.HexToAddress(contractAddress), ethClient.Client)
	if err != nil {
		log.Fatalln(err)
	}

	instancePolygon, err := store.NewStore(common.HexToAddress(contractAddressPolygon), polygonClient.Client)
	if err != nil {
		log.Fatalln(err)
	}

	configService := config.NewConfigService("./config.json")
	dbInfo := structs.DBInfo{
		FacesDBName:                 FacesDBName,
		RarityCollectionName:        rarityCollectionName,
		TransactionsCollectionName:  transactionsCollectionName,
		BlocksCollectionName:        blocksCollectionName,
		BlocksCollectionNamePolygon: blocksCollectionNamePolygon,
		HistoryCollectionName:       historyCollectionName,
		MorphCostCollectionName:     morphCostCollectionName,
	}
	return ethClient, polygonClient, contractAbi, instance, instancePolygon, contractAddress, rootTunnelAddress, configService, dbInfo
}

// main is the entry point of the application.
// It fetches all configurations and starts 1 process:
//- Polling process which processes mint and morph events and stores their metadata in the database
func main() {
	ethClient,
		polygonClient,
		contractAbi,
		instance,
		instancePolygon,
		contractAddress,
		rootTunnelAddress,
		configService,
		dbInfo := initResources()

	recoverAndPoll(
		ethClient,
		polygonClient,
		contractAbi,
		instance,
		instancePolygon,
		contractAddress,
		rootTunnelAddress,
		configService,
		dbInfo)
}

// recoverAndPoll loads transactions and morph cost state in memory from the database and initiates polling mechanism.
//
// Recovery function and polling function is the same.
func recoverAndPoll(ethClient *dlt.EthereumClient, polygonClient *dlt.EthereumClient, contractAbi abi.ABI, store *store.Store, storePolygon *store.Store, contractAddress string, rootTunnelAddress string, configService *structs.ConfigService, dbInfo structs.DBInfo) {
	// Build transactions scramble transaction mapping from db
	txMap, err := handlers.GetTransactionsMapping(dbInfo.FacesDBName, dbInfo.TransactionsCollectionName)
	if err != nil {
		log.Println("Error getting transactions mapping initially.")
	}

	// Build polymorph cost mapping from db
	morphCostMap, err := handlers.GetMorphPriceMapping(dbInfo.FacesDBName, dbInfo.HistoryCollectionName)
	if err != nil {
		log.Println("Error getting morph prices mapping initially.")
	}
	// Recover immediately
	// services.RecoverProcess(ethClient, contractAbi, store, contractAddress, configService, dbInfo, txMap, morphCostMap)
	// Routine one: Start polling after recovery
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	for {
		err = services.RecoverProcess(ethClient, polygonClient, contractAbi, store, storePolygon, contractAddress, rootTunnelAddress, configService, dbInfo, txMap, morphCostMap)
		if err != nil {
			log.WithFields(log.Fields{"error: ": err}).Error("Recovering from error...")
			time.Sleep(15 * time.Second)
			continue
		} else {
			time.Sleep(15 * time.Second)
		}
	}
}
