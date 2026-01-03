package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Bharat1Rajput/crud_golang/config"
	"github.com/Bharat1Rajput/crud_golang/handlers"

	"github.com/gorilla/mux"
)

func main() {

	if err := config.ConnectDB(); err != nil {
		log.Fatal("Could not connect to the database", err)
	}

	defer config.DisconnectDB()

	r := mux.NewRouter()

	// r.HandleFunc("/items", GetItems).Methods("GET")
	r.HandleFunc("/user", handlers.CreateItem).Methods("POST")
	// r.HandleFunc("/items/{id}", GetItem).Methods("GET")
	// r.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")
	// r.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")
	fmt.Printf("your server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
