package main

import (
	"context"
	"fmt"
	"github.com/icrowley/fake"
	"github.com/mrbardia72/sample-mongodb-driver/config"
	"github.com/mrbardia72/sample-mongodb-driver/models"
	"log"
	"time"
)
var userCollection = config.DbConfig().Database(config.MONGO_DATABASE).Collection(config.MONGO_COLLECTION)

func main()  {
	ctx := context.Background()
	fmt.Println("four data faker for users",time.Now())
	for i := 0; i < 4; i++ {

		userfake := models.User{
			Email:    fake.EmailAddress(),
			Password: "password",
			Phone:    fake.Phone(),
			Name:     fake.FemaleFirstName(),
			Age:      33,
			City:     fake.City(),
		}
		insertResult, err := userCollection.InsertOne(ctx, userfake)
		// userCollection.Ins
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a single document: ", insertResult)
	}
}
