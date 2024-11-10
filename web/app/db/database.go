package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = Client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func Disconnect() {
	if err := Client.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Disconnected from MongoDB!")
}
