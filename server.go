package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// check it is okay
func ok(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("content-type", "application/json")
	log.Println("contacts /ok")
	data := map[string]string{"status": "ok"}
	res, _ := json.Marshal(data)
	fmt.Fprintf(w, string(res))
}

// register
func register(w http.ResponseWriter, req *http.Request) {
	log.Println("contacts /register")
	w.Header().Add("content-type", "application/json")
	var t RegisterRequest
	var result bson.M
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		data := map[string]string{"error": "Unable to connect to db. Please try again later!"}
		res, _ := json.Marshal(data)
		http.Error(w, string(res), http.StatusInternalServerError)
		return
	}
	err = json.NewDecoder(req.Body).Decode(&t)
	collection := client.Database("foodly-go").Collection("users")
	token := GenerateToken(t.password)
	_, err = collection.InsertOne(ctx, bson.M{"username": t.username, "password": Hash(t.password), "city": t.city, "email": t.email, "role": t.role, "token": token})
	json.NewEncoder(w).Encode(result)
}

// login
func login(w http.ResponseWriter, req *http.Request) {
	log.Println("contacts /login")
	var t LoginRequest
	var result bson.M
	w.Header().Add("content-type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		data := map[string]string{"error": "Unable to connect to db. Please try again later!"}
		res, _ := json.Marshal(data)
		http.Error(w, string(res), http.StatusInternalServerError)
		return
	}
	err = json.NewDecoder(req.Body).Decode(&t)
	collection := client.Database("foodly-go").Collection("users")
	err = collection.FindOne(ctx, bson.M{"username": t.username, "password": Hash(t.password)}).Decode(&result)
	if err != nil {
		fmt.Println(err)
		data := map[string]string{"error": "User not found. maybe your password is incorrect"}
		res, _ := json.Marshal(data)
		http.Error(w, string(res), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(result)
}
