package integration
//
//
//import (
//	"context"
//	"fmt"
//	"log"
//
//	"github.com/mrbardia72/sample-mongodb-driver/config"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//)
//
//func CheckErr(err error)  {
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//func DbTestConfig() *mongo.Client {
//
//	ctx := context.Background()
//
//	url := options.Client().ApplyURI(config.MONGO_URL)
//
//	// Connect to MongoDB
//	connect, err := mongo.Connect(ctx, url)
//	CheckErr(err)
//
//	// Check the connection
//	err = connect.Ping(ctx, nil)
//	CheckErr(err)
//
//	fmt.Println("Connected to MongoDB!")
//	return connect
//}

