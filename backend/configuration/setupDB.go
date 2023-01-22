package configuration

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://nguyen:nguyen123@mongo:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = client.Connect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to DB !")
	return client
}

var DB *mongo.Database = GetDatabase(ConnectDB(), "realtime_chat")

func GetDatabase(client *mongo.Client, databaseName string) *mongo.Database {
	database := client.Database(databaseName)
	return database
}

func GetCollection(database *mongo.Database, collectionName string) *mongo.Collection {
	collection := database.Collection(collectionName)
	return collection
}
