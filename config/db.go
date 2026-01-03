package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client         *mongo.Client
	DB             *mongo.Database
	UserCollection *mongo.Collection
)

func ConnectDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	var err error
	Client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	err = Client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	DB = Client.Database("myappGO")
	UserCollection = DB.Collection("users")

	log.Println("Connected to MongoDB!")

	return nil

}

func DisconnectDB() error {
	if Client == nil {
		return nil // nothing to disconnect
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := Client.Disconnect(ctx); err != nil {
		return err
	}

	log.Println("Disconnected from MongoDB!")
	return nil
}
