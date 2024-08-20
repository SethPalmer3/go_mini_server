package main

import (
	// "html/template"
	// "io"
	// "os"

	"fmt"
	"log"
	"net/http"
)

const CSSFILE = "home.css"

const PORT = "8000"
const ADDR = "localhost"

func main() {
	s := NewHTTPServer(nil)
	s.addEndPoint("/{page}", struct {
		Para string
		CSS  string
	}{
		Para: "Hello",
		CSS:  CSSFILE,
	})
	s.addHandler("/css/", http.FileServer(http.Dir("./css/")))

	fmt.Println("Starting server on localhost on port " + PORT)
	log.Fatal(http.ListenAndServe(ADDR+":"+PORT, s.multiplexer))
}
