package config

import (
	"context"
	"fmt" 
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CheckErr(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func DbConfig() *mongo.Client {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	url := options.Client().ApplyURI(MONGO_URL)

	// Connect to MongoDB
	connect, err := mongo.Connect(ctx, url)
	CheckErr(err)

	// Check the connection
	err = connect.Ping(ctx, nil)
	CheckErr(err)

	fmt.Println("Connected to MongoDB!")
	return connect
}
