package helpers

import (
	"../config"

	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)
var userCollection = config.DbConfig().Database("goTest").Collection("users") // get collection "users" from db() which returns *mongo.Client

func MinVSMax(w http.ResponseWriter, opertion int) {
	w.Header().Set("Content-Type", "application/json")

	var results []primitive.M

	options := options.Find()
	options.SetSort(bson.D{{"age", opertion}})
	options.SetLimit(1)

	ctx := context.Background()
	cur, err := userCollection.Find(ctx, bson.D{}, options)

	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(context.TODO()) {
		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}
	cur.Close(context.TODO())
	json.NewEncoder(w).Encode(results)
}