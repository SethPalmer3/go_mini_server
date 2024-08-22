package main

import (
	"fmt"
)

const CSSFILE = "home.css"

const PORT = "8000"
const ADDR = "localhost"

func main() {
	s := NewSimpleServer("html", "css", "temp", nil)
	s.ServeHtml("home.html")

	fmt.Println("Starting server on localhost on port " + PORT)
	s.Start(ADDR + ":" + PORT)
}
