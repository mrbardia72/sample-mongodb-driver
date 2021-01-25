package controllers

import (
	"../config"
	"../helpers"
	"../models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

var userCollection = config.DbConfig().Database("goTest").Collection("users") // get collection "users" from db() which returns *mongo.Client

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json") // for adding Content-type

	var person models.User

	err := json.NewDecoder(r.Body).Decode(&person) // storing in person variable of type user
	if err != nil {
		fmt.Print(err)
	}

	insertResult, err := userCollection.InsertOne(ctx, person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult)
	json.NewEncoder(w).Encode(insertResult.InsertedID) // return the mongodb ID of generated document

}
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")

	var body models.User
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	err := userCollection.FindOne(ctx, bson.D{{"name", body.Name}}).Decode(&result)
	if err != nil {

		//fmt.Println(err)
		json.NewEncoder(w).Encode("not existis record")
	}
	json.NewEncoder(w).Encode(result) // returns a Map containing document

}
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}

	var body models.UpdateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {
		fmt.Print(e)
	}

	after := options.After // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"city", body.City}, {"name", body.Name}}}}
	updateResult := userCollection.FindOneAndUpdate(ctx, bson.D{{"_id", _id}}, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)

	json.NewEncoder(w).Encode(result)
}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}
	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase
	res, err := userCollection.DeleteOne(ctx, bson.D{{"_id", _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M                                   //slice for multiple documents

	cur, err := userCollection.Find(ctx, bson.D{{}}) //returns a *mongo.Cursor
	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(ctx) { //Next() gets the next document for corresponding cursor

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem) // appending document pointed by Next()
	}
	cur.Close(context.Background()) // close the cursor once stream of documents has exhausted
	fmt.Println("get all users information")
	json.NewEncoder(w).Encode(results)
}
func GetBettwenName(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M
	cur, err := userCollection.Find(ctx,bson.D{{"name", bson.D{{"$in", bson.A{"erfan", "bardia"}}}}})
	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(ctx) { //Next() gets the next document for corresponding cursor
		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem) // appending document pointed by Next()
	}
	cur.Close(ctx) // close the cursor once stream of documents has exhausted
	fmt.Println("get all users information")
	json.NewEncoder(w).Encode(results)
}

func GetMax(w http.ResponseWriter, r *http.Request) {
	opertion := -1
	helpers.MinVSMax(w, opertion)
}

func GetMin(w http.ResponseWriter, r *http.Request)  {
	opertion := 1
	helpers.MinVSMax(w, opertion)
}

func CountPost(w http.ResponseWriter, r *http.Request)  {
	ctx := context.Background()
	opts := options.Count().SetMaxTime(2 * time.Second)
	count, err := userCollection.CountDocuments(ctx, bson.D{{}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name Bob appears in %v documents", count)
	json.NewEncoder(w).Encode(count)
}