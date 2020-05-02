package main

import (
	"net/http"
	"log"
  "github.com/gorilla/mux"
  "sourcegraph/server/request"
  "sourcegraph/server/middleware"
)

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/view/{topic}", middleware.Chain(
    handlers.View,
    middleware.Log,
  ))
  r.HandleFunc("/edit/{topic}", middleware.Chain(
    handlers.Edit,
    middleware.Log,
  ))
  r.HandleFunc("/save/{topic}", middleware.Chain(
    handlers.Save,
    middleware.Log,
  ))
  http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
