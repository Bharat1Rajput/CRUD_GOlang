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
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the User Management API"))
	})
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/user/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", handlers.DeleteUser).Methods("DELETE")
	fmt.Printf("your server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
