package endpoint

import (
  "net/http"

  "github.com/CWright777/go-api/endpoint/hello"
  "github.com/gorilla/mux"
)

func CreateRouter() http.Handler {
  r := mux.NewRouter()

  r.HandleFunc("/", hello.Get).Methods("GET")

  return r
}
