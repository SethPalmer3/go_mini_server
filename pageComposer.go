package main

import (
	"fmt"
	"html/template"
	"io"
)

type ComposeData struct {
	CSSfile      string
	mainPage     string
	templateData string
}

func compose(w io.Writer, page string, templateData any, errorFunc func()) {
	if templateData == nil {
		errorFunc()
		return
	}
	out, err := template.ParseFiles(page)
	if err != nil {
		fmt.Println("Failed ot parse files")
		fmt.Println(err)
		return
	}
	err = out.Execute(w, templateData)
	if err != nil {
		fmt.Println("Failed to execute template")
		fmt.Println(err)
	}
}
