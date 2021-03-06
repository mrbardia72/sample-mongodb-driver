package routers

import (
	"github.com/gorilla/mux"
	ctr "github.com/mrbardia72/sample-mongodb-driver/controllers"
	"log"
	"net/http"
)

func RoutersMap(s *mux.Router) {
	s.HandleFunc("/createProfile", ctr.CreateProfile).Methods("POST")
	s.HandleFunc("/getAllUsers", ctr.GetAllUsers).Methods("GET")
	s.HandleFunc("/getUserProfile", ctr.GetUserProfile).Methods("POST")
	s.HandleFunc("/updateProfile/{id}", ctr.UpdateProfile).Methods("PUT")
	s.HandleFunc("/deleteProfile/{id}", ctr.DeleteProfile).Methods("DELETE")
	s.HandleFunc("/GetBettwenName", ctr.GetBettwenName).Methods("GET")
	s.HandleFunc("/GetMax/{count}", ctr.GetMax).Methods("GET")
	s.HandleFunc("/GetMin/{count}", ctr.GetMin).Methods("GET")
	s.HandleFunc("/countposts",ctr.CountPost).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server
}