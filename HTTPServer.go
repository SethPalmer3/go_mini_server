package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
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
	s.multiplexer.HandleFunc("/"+s.cssDir+"/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Serving CSS: " + r.URL.Path)
		w.Header().Set("Content-type", "text/css")
		http.ServeFile(w, r, "./"+r.URL.Path[1:])
	})
	return s
}

func (s *SimpleServer) Start(addr string) {
	log.Fatal(http.ListenAndServe(addr, s.multiplexer))
}

func (s *SimpleServer) ServeHtml(fileName string) {
	s.multiplexer.HandleFunc("/"+s.htmlDir+"/"+fileName, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Serving HTML: " + r.URL.Path)
		mainTemplatePath := filepath.Join(s.htmlDir, fileName)
		mainTemplate, err := template.ParseFiles(mainTemplatePath)
		if err != nil {
			fmt.Println("Cannot parse main template")
			fmt.Println(err)
			return
		}
		embeddedTemplatePath := filepath.Join(s.templateDir, "temp_"+fileName)
		embbededTemplate, err := template.Must(mainTemplate.ParseFiles(embeddedTemplatePath)).ParseFiles(embeddedTemplatePath)
		if err != nil {
			fmt.Println("Cannot parse embedded template")
			fmt.Println(err)
			return
		}

		w.Header().Set("Content-type", "text/html")
		err = embbededTemplate.Execute(w, nil)

		if err != nil {
			fmt.Println("Error executing files")
			fmt.Println(err)
			return
		}
	})
}
