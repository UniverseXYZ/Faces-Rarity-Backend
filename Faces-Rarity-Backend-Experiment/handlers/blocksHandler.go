package handlers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"rarity-backend/constants"
	"rarity-backend/db"
	"rarity-backend/models"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetLastProcessedBlockNumber @GetLastProcessedBlockNumber fetches the last processed block number from the block mongo collection.
// At any point of the application there should be only one record in the collection
// If no collection or record exists - returns the block number of the contract deployment transaction.
func GetLastProcessedBlockNumber(facesDBName string, ethereum bool) (int64, error) {
	deploymentBlockNumberEth := os.Getenv("DEPLOYMENT_BLOCK_NUMBER_ETHEREUM")
	deploymentBlockNumberPolygon := os.Getenv("DEPLOYMENT_BLOCK_NUMBER_POLYGON")

	blocksCollectionNameEth := os.Getenv("BLOCKS_COLLECTION")
	blocksCollectionNamePolygon := os.Getenv("BLOCKS_COLLECTION_POLYGON")

	ethBlockNumber, err := strconv.Atoi(deploymentBlockNumberEth)

	if err != nil {
		log.Printf("%d is not a valid block number integer ", deploymentBlockNumberEth)
	}

	polygonBlockNumber, err := strconv.Atoi(deploymentBlockNumberPolygon)

	if err != nil {
		log.Printf("%d is not a valid block number integer ", deploymentBlockNumberPolygon)
	}
	var collection *mongo.Collection
	if ethereum {
		collection, err = db.GetMongoDbCollection(facesDBName, blocksCollectionNameEth)
	} else {
		collection, err = db.GetMongoDbCollection(facesDBName, blocksCollectionNamePolygon)
	}

	if err != nil && ethereum {
		return int64(ethBlockNumber), err
	} else if err != nil && !ethereum {
		return int64(polygonBlockNumber), err
	}

	lastBlock := collection.FindOne(context.Background(), bson.M{})

	if lastBlock.Err() != nil {
		return 0, lastBlock.Err()
	}

	var result bson.M
	lastBlock.Decode(&result)

	if result == nil {
		return 0, err
	}

	lastProcessedBlockNumber := result[constants.BlockFieldNames.Number]
	block := lastProcessedBlockNumber.(int64)
	if ethereum && block < int64(ethBlockNumber) {
		return int64(ethBlockNumber), nil
	} else if !ethereum && block < int64(polygonBlockNumber) {
		return int64(polygonBlockNumber), nil
	}
	return block, nil
}

// CreateOrUpdateLastProcessedBlock persists the passed block number in the parameters to the block collection. At any point of the application there should be only one record in the collection
//
// If no collection or records exists - it will create a new one.
func CreateOrUpdateLastProcessedBlock(number uint64, dbName string, blocksCollectionName string) (string, error) {
	collection, err := db.GetMongoDbCollection(dbName, blocksCollectionName)
	if err != nil {
		return "", err
	}

	entity := models.ProcessedBlockEntity{Number: number}

	update := bson.M{
		"$set": entity,
	}

	// This option will create new entity if no matching is found
	opts := options.Update().SetUpsert(true)

	objID, _ := primitive.ObjectIDFromHex(strconv.FormatInt(0, 16))
	filter := bson.M{constants.BlockFieldNames.ObjId: objID}

	_, err = collection.UpdateOne(context.Background(), filter, update, opts)

	if err != nil {
		return "", err
	}

	return "Successfully persisted new last processed block number: " + strconv.FormatUint(number, 10), nil
}
