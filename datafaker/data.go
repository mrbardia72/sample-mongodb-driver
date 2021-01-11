package main

import (
	"../config"
	"../models"
	"context"
	"fmt"
	"github.com/icrowley/fake"
	"log"
	"time"
)
var userCollection = config.DbConfig().Database("goTest").Collection("users")


func main()  {
	fmt.Println("four data faker for users",time.Now())
	for i := 0; i < 4; i++ {

		userfake := models.User{
			Email:    fake.EmailAddress(),
			Password: "xv5tdj4%F%&%%f",
			Phone:    fake.Phone(),
			Name:     fake.FemaleFirstName(),
			Age:      33,
			City:     fake.City(),
		}
		insertResult, err := userCollection.InsertOne(context.TODO(), userfake)
		// userCollection.Ins
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Inserted a single document: ", insertResult)
	}
}
