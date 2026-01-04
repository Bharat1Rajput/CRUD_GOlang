package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Bharat1Rajput/crud_golang/config"
	"github.com/Bharat1Rajput/crud_golang/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user.ID = primitive.NewObjectID()

	_, err = config.UserCollection.InsertOne(ctx, user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// Additional handler functions (GetUser, GetUsers, UpdateUser, DeleteUser) can be added here

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting all users
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := config.UserCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var users []models.User

	if err = cursor.All(ctx, &users); err != nil {
		http.Error(w, "error decoding users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single user by ID
	w.Header().Set("Content-Type", "application/json")

	// Get the user ID from the URL path
	params := mux.Vars(r)
	idParams := params["id"]

	objectId, err := primitive.ObjectIDFromHex(idParams)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err = config.UserCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a user by ID
	w.Header().Set("Content-Type", "application/json")

	parmas := mux.Vars(r)
	idParams := parmas["id"]

	objectId, err := primitive.ObjectIDFromHex(idParams)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "error decoding user", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": user,
	}

	_, err = config.UserCollection.UpdateOne(ctx, bson.M{"_id": objectId}, update)
	if err != nil {
		http.Error(w, "error updating user", http.StatusInternalServerError)
		return
	}

	user.ID = objectId

	json.NewEncoder(w).Encode(user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a user by ID
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idParams := params["id"]
	objectId, err := primitive.ObjectIDFromHex(idParams)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = config.UserCollection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		http.Error(w, "error deleting user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "user deleted successfully"})
}
