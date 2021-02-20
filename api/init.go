package api

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ctx context.Context
var stonksDatabase *mongo.Database
var stocksCollection *mongo.Collection
var stocksListCollection *mongo.Collection
var err error

func init() {
	mongoURI := os.Getenv("MONGO_URI")
	client, err = mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		fmt.Println("Error Connecting to MongoDB Atlas")
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	// defer client.Disconnect(ctx)
	if err != nil {
		fmt.Println("Error in connection to mongo cluster")
		log.Fatal(err)
	}

	stonksDatabase = client.Database("stonks")
	stocksCollection = stonksDatabase.Collection("stocksData")
	stocksListCollection = stonksDatabase.Collection("stocksList")
}
