package handlers

import (
	"context"
	"rarity-backend/constants"
	"rarity-backend/db"
	"rarity-backend/models"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetLastProcessedBlockNumber fetches the last processed block number from the block collection. At any point of the application there should be only one record in the collection
//
// If no collection or record exists - returns 0. This means the application will start processing from the beggining.
func GetLastProcessedBlockNumber(polymorphDBName string, blocksCollectionName string) (int64, error) {
	collection, err := db.GetMongoDbCollection(polymorphDBName, blocksCollectionName)
	if err != nil {
		return 0, err
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
	return block, nil
}

// CreateOrUpdateLastProcessedBlock persists the passed block number in the parameters to the block collection. At any point of the application there should be only one record in the collection
//
// If no collection or records exists - it will create a new one.
func CreateOrUpdateLastProcessedBlock(number uint64, polymorphDBName string, blocksCollectionName string) (string, error) {
	collection, err := db.GetMongoDbCollection(polymorphDBName, blocksCollectionName)
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
