package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Email string `json:"email" validate:"required,email"`
}

func main() {
	v := validator.New()
	a := User{
		Email: "a",
	}
	err := v.Struct(a)

	//for _, e := range err.(validator.ValidationErrors) {
		fmt.Println(err)
	//}
}
