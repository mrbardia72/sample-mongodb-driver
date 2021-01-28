package main

import (
	"github.com/gorilla/mux"
	"github.com/mrbardia72/sample-mongodb-driver/routers"
)

func main() {
	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path
	//Routes map
	routers.RoutersMap(s)
}