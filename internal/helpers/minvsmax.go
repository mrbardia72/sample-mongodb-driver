package helpers

import (
	"github.com/mrbardia72/sample-mongodb-driver/config"

	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)
var userCollection = config.DbConfig().Database(config.MONGO_DATABASE).Collection(config.MONGO_COLLECTION) // get collection "users" from db() which returns *mongo.Client

func MinVSMax(w http.ResponseWriter, opertion int,a int64) {

	w.Header().Set("Content-Type", "application/json")

	var results []primitive.M

	mg := options.Find()
	mg.SetSort(bson.D{{"age", opertion}})
	mg.SetLimit(a)

	ctx := context.Background()
	cur, err := userCollection.Find(ctx, bson.D{}, mg)

	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(ctx) {
		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}
	cur.Close(ctx)
	json.NewEncoder(w).Encode(results)
}