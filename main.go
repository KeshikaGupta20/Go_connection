package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	c "github.com/KeshikaGupta20/Go_connection/fiber/crud"

)

func main() {
	//app := fiber.New()
	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path
	//Routes
	s.HandleFunc("/createProduct", c.createProduct).Methods("POST")
	s.HandleFunc("/getAllUsers", c.getAllUsers).Methods("GET")
	s.HandleFunc("/getproductProfile", c.getProductProfile).Methods("POST")
	s.HandleFunc("/updateProduct", c.updateProduct).Methods("PUT")
	s.HandleFunc("/deleteProduct/{id}", c.deleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server
	//app.Listen(":3001")
}
