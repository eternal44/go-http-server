package main

import (
	"net/http"
  "sourcegraph/server/request"
	"log"
  "github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", handlers.HeartBeat)
	r.HandleFunc("/view/{topic}", handlers.ViewHandler).Name("view")
  http.Handle("/", r)

  // http.HandleFunc("/", handlers.MultipleMiddleware(
  //   handlers.HeartBeat))
	http.HandleFunc("/edit/", handlers.MakeHandler(handlers.EditHandler))
	http.HandleFunc("/save/", handlers.MakeHandler(handlers.SaveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
