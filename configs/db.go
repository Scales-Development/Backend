package configs

import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo() *mongo.Client {
	opts := options.Client().ApplyURI(EnvMongoURI())
	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatal(err)
	}

	

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel();

	defer func() { // If the client unexpectedly disconnects, it will panic and error it out
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err) // Basically it's establishing a ping and if it's not able to ping or establish a proper connection it will error out
	}

	fmt.Println("Successfully established a connection to MongoDB")
	return client
}

// Global variable to import to other files to establish a connection to mongo
var DB *mongo.Client = ConnectToMongo()

// A method to fetch a specific collection in the database on mongo
func FetchCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("Scales").Collection(collectionName)
	return collection
}