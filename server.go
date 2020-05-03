package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sourcegraph/server/middleware"
	"sourcegraph/server/request"
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
