package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var once sync.Once
var instance *mongo.Client

// GetDbConnection returns an instance of a mongo db client. This is a singleton pattern in order to have only one alive connection to the database.
//
// If no connection exists, it will connect to database.
//
// If connection exists, it will return the instance of the database
func GetDbConnection() (*mongo.Client, error) {
	if instance != nil {
		log.Println("Fetching existing MongoDB client")
		err := instance.Ping(context.Background(), nil)
		if err != nil {
			return nil, err
		} else {
			log.Println("Successfully fetched existing MongoDB client.")
			return instance, nil
		}
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbUrl := os.Getenv("DB_URL")

	if username == "" {
		log.Fatal("Missing username in .env")
	}
	if password == "" {
		log.Fatal("Missing password in .env")
	}
	if dbUrl == "" {
		log.Fatal("Missing db url in .env")
	}

	connectionStr := "mongodb+srv://" + username + ":" + password + "@" + dbUrl + "?retryWrites=true&w=majority"
	var err error
	instance, err = mongo.Connect(context.Background(), options.Client().ApplyURI(connectionStr))
	if err != nil {
		return nil, err
	}

	// check the connection
	err = instance.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	} else {
		log.Println("Connected to mongo client")
	}
	return instance, nil
}

const (
	// Timeout operations after N seconds
	connectTimeout = 5
	queryTimeout   = 30

	// Which instances to read from
	readPreference           = "secondaryPreferred"
	connectionStringTemplate = "mongodb://%s:%s@%s/test?replicaSet=rs0&readpreference=%s&connect=direct&sslInsecure=true&retryWrites=false"
)

// connectToDb retrieves db config from .env and tries to conenct to the database.
func connectToDb() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbUrl := os.Getenv("DB_URL")

	if username == "" {
		log.Fatalln("Missing username in .env")
	}
	if password == "" {
		log.Fatalln("Missing password in .env")
	}
	if dbUrl == "" {
		log.Fatalln("Missing db url in .env")
	}

	// connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, dbUrl, readPreference)

	connectionStr := "mongodb+srv://" + username + ":" + password + "@" + dbUrl + "?retryWrites=true&w=majority"

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionStr))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping cluster: %v", err)
	}

	fmt.Println("Connected to DocumentDB!")

	return client
}

// checkConnectionAndRestore ping the client and it throws and error, it tries to reconnect.
func checkConnectionAndRestore(client *mongo.Client) {
	err := client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal(err)
		newClient := connectToDb()
		client = newClient
	}
}

// GetMongoDbCollection accepts dbName and collectionname and returns an instance of the specified collection.
func GetMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {
	client, err := GetDbConnection()
	if err != nil {
		return nil, err
	}

	collection := client.Database(DbName).Collection(CollectionName)
	return collection, nil
}

func DisconnectDB() {
	if instance == nil {
		return
	}

	err := instance.Disconnect(context.TODO())
	if err != nil {
		fmt.Println("FAILED TO CLOSE Mongo Connection")
		fmt.Println(err)
	} else {
		fmt.Println("Connection to MongoDB closed.")
	}
}
