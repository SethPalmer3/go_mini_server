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
	s := NewSimpleServer("./html", "./css/", "./temp", nil)

	fmt.Println("Starting server on localhost on port " + PORT)
	log.Fatal(http.ListenAndServe(ADDR+":"+PORT, s.multiplexer))
}
