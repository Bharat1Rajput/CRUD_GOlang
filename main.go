package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var items []Item
var NextID int = 1

// get all items
func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// logic to get all items
	// respond with items
	json.NewEncoder(w).Encode(items)

}

// create a new item
func CreateItem(w http.ResponseWriter, r *http.Request) {
	// logic to create a new item
	var newItem Item
	json.NewDecoder(r.Body).Decode(&newItem)
	newItem.ID = NextID
	NextID++
	items = append(items, newItem)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newItem)
}

// get a specific item by id
func GetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, item := range items {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "Item not Found"})
}

// update an existing item by id
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var updatedItem Item
	json.NewDecoder(r.Body).Decode(&updatedItem)

	for i, item := range items {
		if item.ID == id {
			items[i].Name = updatedItem.Name
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(items[i])
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"msg": "Item not found"})
}

// delete an item by id
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "item deleted successfully"})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"msg": "Item not found"})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/items", GetItems).Methods("GET")
	r.HandleFunc("/items", CreateItem).Methods("POST")
	r.HandleFunc("/items/{id}", GetItem).Methods("GET")
	r.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")
	fmt.Printf("your server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
