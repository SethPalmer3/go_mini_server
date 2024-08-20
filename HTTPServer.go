package main

import (
	"log"
	"net/http"
	"os"
)

type SimpleServer struct {
	htmlDir     string
	cssDir      string
	templateDir string
	multiplexer *http.ServeMux
}

func NewSimpleServer(htmlDir string, cssDir, templateDir string, multiplxr *http.ServeMux) *SimpleServer {
	s := new(SimpleServer)
	err := os.MkdirAll(htmlDir, os.ModePerm)
	if err != nil {
		panic("Could not create " + htmlDir)
	}
	s.htmlDir = htmlDir
	err = os.MkdirAll(cssDir, os.ModePerm)
	if err != nil {
		panic("Could not create " + cssDir)
	}
	s.cssDir = cssDir
	err = os.MkdirAll(templateDir, os.ModePerm)
	if err != nil {
		panic("Could not create " + templateDir)
	}
	s.templateDir = templateDir
	if multiplxr == nil {
		s.multiplexer = http.NewServeMux()
	} else {
		s.multiplexer = multiplxr
	}
	return s
}

func (s *SimpleServer) Start(addr string) {
	log.Fatal(http.ListenAndServe(addr, s.multiplexer))
}
