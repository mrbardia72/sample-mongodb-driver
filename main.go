package main

import (
	"./controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path

	//Routes

	s.HandleFunc("/createProfile", controllers.CreateProfile).Methods("POST")
	s.HandleFunc("/getAllUsers", controllers.GetAllUsers).Methods("GET")
	s.HandleFunc("/getUserProfile", controllers.GetUserProfile).Methods("POST")
	s.HandleFunc("/updateProfile/{id}", controllers.UpdateProfile).Methods("PUT")

	s.HandleFunc("/deleteProfile/{id}", controllers.DeleteProfile).Methods("DELETE")

	s.HandleFunc("/GetBettwenName", controllers.GetBettwenName).Methods("GET")
	s.HandleFunc("/GetMax", controllers.GetMax).Methods("GET")
	s.HandleFunc("/GetMin", controllers.GetMin).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server
}
