package handlers

import (
	"context"
	"log"
	"rarity-backend/constants"
	"rarity-backend/db"
	"rarity-backend/models"
	"rarity-backend/structs"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpdateAllRanking fetches all polymorph from the database and calculates the ranks.
//
// After the ranking is done, the changes to ranks are persisted in the database.
func UpdateAllRanking(polymorphDBName string, rarityCollectionName string) error {
	ranking := structs.RankMutex{}
	collection, err := db.GetMongoDbCollection(polymorphDBName, rarityCollectionName)
	if err != nil {
		return err
	}

	var entities []models.PolymorphEntity

	var findOptions options.FindOptions
	findOptions.SetLimit(10000)
	findOptions.SetSort(bson.D{{constants.MorphFieldNames.RarityScore, -1}, {constants.MorphFieldNames.TokenId, 1}})
	results, err := collection.Find(context.Background(), bson.D{}, &findOptions)
	if err != nil {
		log.Println("Error finding document from rarity collection. ", err)
	}

	results.All(context.Background(), &entities)

	var wg sync.WaitGroup
	for i, entity := range entities {
		wg.Add(1)
		go setRank(entity, i+1, &ranking, &wg)
	}
	wg.Wait()
	if len(ranking.Operations) > 0 {
		err = PersistMultiplePolymorphs(ranking.Operations, polymorphDBName, rarityCollectionName)
		if err != nil {
			log.Println("Error persisting multiple polymorphs when updating ranking. ", err)
			return err
		}
	}
	return nil
}

// setRank computes if there should be a change in the current polymorph's rank.
// If there's a change we concurrently prepare an update operation and append it to the update operations array
//
// Mutes and WaitGroup are used to prevent race conditions
func setRank(entity models.PolymorphEntity, newRank int, ranking *structs.RankMutex, wg *sync.WaitGroup) {

	if entity.Rank != newRank {
		operation := mongo.NewUpdateOneModel()
		operation.SetFilter(bson.M{constants.MorphFieldNames.TokenId: entity.TokenId})
		operation.SetUpdate(bson.M{"$set": bson.M{constants.MorphFieldNames.Rank: newRank}})
		ranking.Mutex.Lock()
		ranking.Operations = append(ranking.Operations, operation)
		ranking.Mutex.Unlock()
	}
	wg.Done()
}
