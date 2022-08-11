package handlers

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
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

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	var bdoc interface{}
	json, err := json.Marshal(entity)
	if err != nil {
		return err
	}
	err = bson.UnmarshalExtJSON(json, false, &bdoc)
	if err != nil {
		return err
	}

	_, err = collection.InsertOne(context.Background(), bdoc)
	if err != nil {
		log.WithFields(log.Fields{"original error: ": err}).Error("error inserting a document in History collection")
	}

	log.Infof("Inserted history snapshot for polymorph #" + strconv.Itoa(entity.TokenId))
	return nil
}
