package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mc *mongo.Client

func getMongoClient() *mongo.Client {
	if mc != nil {
		err := mc.Ping(context.Background(), readpref.Primary())
		if err == nil {

			return mc
		}
	}
	// 1. Build Connection URI
	// Format: mongodb://admin:secretpassword@mongo:27017
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		os.Getenv("MONGODB_USERNAME"),
		os.Getenv("MONGODB_PASSWORD"),
		os.Getenv("MONGODB_HOST"),
		os.Getenv("MONGODB_PORT"),
	)

	// 2. Set Client Options
	clientOptions := options.Client().ApplyURI(uri)

	// 3. Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Connection failed:", err)
	}

	// 4. Ping the database to verify credentials/connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Ping failed (Check your credentials!):", err)
	}

	fmt.Println("Successfully connected to MongoDB!")
	mc = client
	return client
}
