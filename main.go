package main

import (
  "fmt"
  "net/http"
  "log"
  "encoding/json"
)

func ok(w http.ResponseWriter, req *http.Request){
  log.Printf("%s contacts /ok", req.RemoteAddr)
  data := map[string]string{"status": "ok"}
  res, _ := json.Marshal(data)
  fmt.Fprintf(w, string(res))
}

func main() {
  http.HandleFunc("/ok", ok)
  http.ListenAndServe(":5000", nil)
}
