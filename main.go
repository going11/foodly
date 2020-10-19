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
  log.Println("contacts /register")
  w.Header().Add("content-type", "application/json")
  var t RegisterRequest
  ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
  defer cancel()
  client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
  defer client.Disconnect(ctx)
  err = client.Ping(ctx, readpref.Primary())
  if err != nil {
    data := map[string]string{"error": "Unable to connect to db. Please try again later!"}
    res, _ := json.Marshal(data)
    http.Error(w, string(res) ,http.StatusInternalServerError)
  }
  err = json.NewDecoder(req.Body).Decode(&t)
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
  client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
  defer client.Disconnect(ctx)
  err = client.Ping(ctx, readpref.Primary())
  if err != nil {
    data := map[string]string{"error": "Unable to connect to db. Please try again later!"}
    res, _ := json.Marshal(data)
    http.Error(w, string(res),http.StatusInternalServerError)
  }
  err = json.NewDecoder(req.Body).Decode(&t)
  collection := client.Database("foodly-go").Collection("users")
  err = collection.FindOne(ctx, bson.M{"username": t.username, "password": Hash(t.password)}).Decode(&result)
  if err != nil {
    data := map[string]string{"error": "User not found. maybe your password is incorrect"}
    res, _ := json.Marshal(data)
    http.Error(w, string(res), http.StatusBadRequest)
  }
  json.NewEncoder(w).Encode(result)
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/login", login).Methods("POST")
  r.HandleFunc("/register", register).Methods("POST")
  r.HandleFunc("/ok", ok)
  http.ListenAndServe(":5000", r)
}
