package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"rarity-backend/db"
	"rarity-backend/models"

	"go.mongodb.org/mongo-driver/bson"
)

// GetTransactionsMapping fetches all records from the transactions collections. Returns a mapping of the records.
//
// The application has to know which morph events have already been processed in order to prevent duplicate false information stored in database
func GetTransactionsMapping(polymorphDBName string, transactionsColl string) (map[string]map[uint]bool, error) {
	collection, err := db.GetMongoDbCollection(polymorphDBName, transactionsColl)
	if err != nil {
		return nil, err
	}

	var transactions []models.Transaction
	results, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error fetching transactions document. ", err)
	}
	results.All(context.Background(), &transactions)

	txMap := make(map[string]map[uint]bool)

	for _, tx := range transactions {
		nestedMap, ok := txMap[tx.TxHash]
		if !ok {
			nestedMap = make(map[uint]bool)
			txMap[tx.TxHash] = nestedMap
		}
		txMap[tx.TxHash][tx.LogIndex] = true
	}

	return txMap, nil
}

// SaveTransaction persists the processed transaction in the database
//
// If the application stops it will be able to load the processed event in memory from the database
func SaveTransaction(polymorphDBName string, transactionsColl string, transaction models.Transaction) {
	collection, err := db.GetMongoDbCollection(polymorphDBName, transactionsColl)
	if err != nil {
		log.Fatalln(err)
	}

	var bdoc interface{}
	json, _ := json.Marshal(transaction)
	bson.UnmarshalExtJSON(json, false, &bdoc)

	_, err = collection.InsertOne(context.Background(), bdoc)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("\nInserted new transaction in DB:\ntxHash: %v\nLogIndex: %v\n", transaction.TxHash, transaction.LogIndex)
}

// PersistMintEvents persists all the processed mints in the database in one go.
// Bulk writing to database saves a lot of time
func SaveTransactions(bsonDocs []interface{}, polymorphDBName string, transactionsCollectionName string) error {
	collection, err := db.GetMongoDbCollection(polymorphDBName, transactionsCollectionName)
	if err != nil {
		return err
	}

	res, err := collection.InsertMany(context.Background(), bsonDocs)
	if err != nil {
		log.Println("Error inserting many Transaction documents into DB. ", err)
		return err
	}
	log.Println(fmt.Sprintf("Inserted %v transactions in DB", len(res.InsertedIDs)))
	return nil
}
