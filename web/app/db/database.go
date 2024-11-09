package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// Connect initializes the MongoDB client and connects to the database.
func Connect() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Replace with your MongoDB URI

	// Set a timeout for the connection
	Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ensure the connection is established
	err = Client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

// Disconnect closes the MongoDB connection.
func Disconnect() {
	if err := Client.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Disconnected from MongoDB!")
}
