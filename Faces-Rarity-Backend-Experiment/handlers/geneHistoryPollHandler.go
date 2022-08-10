package handlers

import (
	"context"
	"encoding/json"
	"log"
	"rarity-backend/db"
	"rarity-backend/models"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

// SavePolymorphHistory persists the polymorph history snapshot to the database.
func SavePolymorphHistory(entity models.PolymorphHistory, polymorphDBName string, historyCollectionName string) error {
	collection, err := db.GetMongoDbCollection(polymorphDBName, historyCollectionName)

	if err != nil {
		return err
	}

	var bdoc interface{}
	json, err := json.Marshal(entity)
	if err != nil {
		return err
	}
	err = bson.UnmarshalExtJSON(json, false, &bdoc)
	if err != nil {
		return err
	}

	//fmt.Println(bdoc)

	_, err = collection.InsertOne(context.Background(), bdoc)
	if err != nil {
		log.Println("Error inserting a document in History collection. ", err)
	}

	log.Println("Inserted history snapshot for polymorph #" + strconv.Itoa(entity.TokenId))
	return nil
}
