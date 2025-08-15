package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectMongo () {
	
	uri := os.Getenv("MONGOURI")

	if uri == "" {
	log.Fatal("Mongo uri is required")
	}


	ClientOption := options.Client().ApplyURI(uri)

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancle()

	client, err := mongo.Connect(ctx, ClientOption)

	if err != nil {
		log.Fatal(err)
	}

	Client = client

	log.Printf("mongo connected")
}