package hello

import (
  "net/http"
  "log"
)

func Get(w http.ResponseWriter, r *http.Request) {
  log.Println("Responding to /hello request")
}
