package models

import (
	"gopkg.in/go-playground/validator.v9"
	"strings"

	"fmt"
)

type User struct {

	Email 		string `json:"email" validate:"required"`
	Password 	string `json:password`
	Phone 		string `json:phone`
	Name 		string `json:name`
	Age  		int    `json:age`
	City 		string `json:city`
}

type Post struct {

	Title 		string `json:"title"`
	Body 		string `json:"body"`
	User 		[]*User `json:"user"`
	Category 	[]*Category `json:"category"`
	Comment 	[]*Comment `json:"comment"`
}

type Comment struct {

	Body 		string `json:"body"`
}

type Category struct {

	name 		string `json:"name"`
}

type UpdateBody struct {

	Name 		string `json:"name"` //value that has to be matched
	City 		string `json:"city"` // value that has to be modified
}