package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	c "github.com/KeshikaGupta20/Go_connection/fiber/crud/controller"
)

func main() {
	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path
	//Routes
	s.HandleFunc("/createProfile", c.createProfile).Methods("POST")
	s.HandleFunc("/getAllUsers", c.getAllUsers).Methods("GET")
	s.HandleFunc("/getUserProfile", c.getUserProfile).Methods("POST")
	s.HandleFunc("/updateProfile", c.updateProfile).Methods("PUT")
	s.HandleFunc("/deleteProfile/{id}", c.deleteProfile).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server
}
