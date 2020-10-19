package main

import (
  "fmt"
  "net/http"
  "log"
  "encoding/json"
  "context"
  "time"
  "github.com/gorilla/mux"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/mongo/readpref"
  "go.mongodb.org/mongo-driver/bson"
)

var client *mongo.Client

// check it is okay
func ok(w http.ResponseWriter, req *http.Request){
  w.Header().Add("content-type", "application/json")
  log.Printf("contacts /ok")
  data := map[string]string{"status": "ok"}
  res, _ := json.Marshal(data)
  fmt.Fprintf(w, string(res))
}

// register
func register(w http.ResponseWriter, req *http.Request) {
  ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
  defer cancel()
  log.Println("contacts /register")
  w.Header().Add("content-type", "application/json")
  var t RegisterRequest
  err := json.NewDecoder(req.Body).Decode(&t)
  if err != nil {
    http.Error(w, "unable to parse json", http.StatusBadRequest)
  }
  collection := client.Database("foodly-go").Collection("users")
  token := GenerateToken(t.password)
  _, err = collection.InsertOne(ctx, bson.M{"username": t.username, "password": Hash(t.password), "city":t.city, "email":t.email, "role":t.role, "token": token})
  data := map[string]string{"token": token}
  res, _ := json.Marshal(data)
  fmt.Fprintf(w, string(res))
}

// login 
func login(w http.ResponseWriter, req *http.Request){
  log.Println("contacts /login")
  var t LoginRequest
  var result bson.M
  w.Header().Add("content-type", "application/json")
  ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
  defer cancel()
  err := json.NewDecoder(req.Body).Decode(&t)
  if err != nil {
    http.Error(w, "Wrong Request", http.StatusBadRequest)
  }
  collection := client.Database("foodly-go").Collection("users")
  err = collection.FindOne(ctx, bson.M{"username": t.username, "password": Hash(t.password)}).Decode(&result)
  if err != nil {
    http.Error(w, "User not found. maybe your password is incorrect", http.StatusBadRequest)
  }
  json.NewEncoder(w).Encode(result)
}

func main() {
  ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
  defer cancel()
  client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
  defer client.Disconnect(ctx)
  err = client.Ping(ctx, readpref.Primary())
  if err != nil {
    panic("database error")
  }
  r := mux.NewRouter()
  r.HandleFunc("/login", login).Methods("POST")
  r.HandleFunc("/register", register).Methods("POST")
  r.HandleFunc("/ok", ok)
  http.ListenAndServe(":5000", r)
}
