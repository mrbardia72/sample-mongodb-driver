package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	ctr "github.com/mrbardia72/sample-mongodb-driver/controllers"
	"github.com/mrbardia72/sample-mongodb-driver/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateProfilePass(t *testing.T) {

	data := []struct {
		inputJSON  	string
		statusCode 	int
		email      	string
		password   	string
		phone		string
		name		string
		age 		int
		city 		string
		errMessage 	string
	}{
		{
			inputJSON:  `{"email":"wwwwwwww@sama.sm", "password": "11q2we3r45t","phone":"09878465399","name":"Roomba","age":374,"city":"japan"}`,
			statusCode: 201,
			email:"wwwwwwww@sama.sm",
			password:"11q2we3r45t",
			phone:"09878465399",
			name: "Roomba",
			age: 374,
			city: "japan",
			errMessage: "",
		},
	}
	for _, v := range data {

		s := mux.NewRouter()
		s.HandleFunc("/createProfile", ctr.CreateProfile).Methods("POST")

		req, err := http.NewRequest(http.MethodPost, "/createProfile", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}

		rr := httptest.NewRecorder()
		s.ServeHTTP(rr, req)

		var pp models.User
		err = json.Unmarshal(rr.Body.Bytes(), &pp)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}
		fmt.Println("this is the response data: ", pp)
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 201 {
			//casting the interface to map:
			assert.Equal(t, pp.Email, v.email)
			//assert.Equal(t, responseMap["password"], v.password)
			//assert.Equal(t, responseMap["phone"], v.phone)
			//assert.Equal(t, responseMap["name"], v.name)
			//assert.Equal(t, responseMap["age"], v.age)
			//assert.Equal(t, responseMap["city"], v.city)
		}
		if v.statusCode == 400 || v.statusCode == 422 || v.statusCode == 500 && v.errMessage != "" {
			assert.Equal(t, pp.Email, v.errMessage)
		}
	}
}