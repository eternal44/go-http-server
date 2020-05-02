package main

import (
  "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
  "http/server/request/handlers"
)

func main() {
  fmt.Println("hello")
	// http.HandleFunc("/view/", makeHandler(viewHandler))
	// http.HandleFunc("/edit/", makeHandler(editHandler))
	// http.HandleFunc("/save/", makeHandler(saveHandler))

	// log.Fatal(http.ListenAndServe(":8080", nil))
}
