package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

const dbName = "go_mongo"
const mongoURI = "mongodb://localhost:27017/" + dbName

func Connect() error{
	// Database Config
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()


	err = client.Connect(ctx)
	//Cancel context to avoid memory leak

	// Ping our db connection
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	// Connect to the database
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}


