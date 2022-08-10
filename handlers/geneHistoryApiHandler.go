package handlers

import (
	"context"
	"encoding/json"
	"os"
	"rarity-backend/constants"
	"rarity-backend/db"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

// GetPolymorphHistory endpoints accepts id of a single polymorph and returns all history snapshot from the database.
//
// History snapshots represent the changes made by scrambling or morphing this polymorph.
func GetPolymorphHistory(c *fiber.Ctx) {
	godotenv.Load()

	polymorphDBName := os.Getenv("POLYMORPH_DB")
	historyColelctionName := os.Getenv("HISTORY_COLLECTION")

	collection, err := db.GetMongoDbCollection(polymorphDBName, historyColelctionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var filter bson.M = bson.M{}
	if c.Params("id") != "" {
		id := c.Params("id")
		filter = bson.M{constants.MorphFieldNames.TokenId: id}
	}

	var results []bson.M
	curr, err := collection.Find(context.Background(), filter)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	defer curr.Close(context.Background())

	curr.All(context.Background(), &results)
	//TODO: Convert tokenid to be integer in database

	if results == nil {
		c.Send(bson.M{})
		return
	}

	json, _ := json.Marshal(results)
	c.Send(json)

}
