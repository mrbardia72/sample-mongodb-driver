package main

import (
	"./routers"
	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path
	//Routes map
	routers.RoutersMap(s)
}