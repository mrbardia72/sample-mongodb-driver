package config


import (
	"context"
	"fmt" 
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CheckErr(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func DbConfig() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	CheckErr(err)

	// Check the connection
	err = client.Ping(context.Background(), nil)
	CheckErr(err)

	fmt.Println("Connected to MongoDB!")
	return client
}
