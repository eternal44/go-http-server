package main

import (
	"net/http"
  "sourcegraph/server/request/handlers"
	"log"
)

func main() {
	http.HandleFunc("/view/", handlers.MakeHandler(handlers.ViewHandler))
	http.HandleFunc("/edit/", handlers.MakeHandler(handlers.EditHandler))
	http.HandleFunc("/save/", handlers.MakeHandler(handlers.SaveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
