package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

type EndPointAdditions struct {
	cssPaths     []string
	jsPaths      []string
	templateData any
}

type EndPointData struct {
	requestPath string
	// hostPath       string
	additionalData any
}

type HTTPServer struct {
	multiplexer *http.ServeMux
	endPoints   []EndPointData
}

func NewHTTPServer(sm *http.ServeMux) *HTTPServer {
	s := new(HTTPServer)
	if sm != nil {
		s.multiplexer = sm
	} else {
		s.multiplexer = http.NewServeMux()
	}
	return s
}

func (s *HTTPServer) addHandler(requstPath string, handler http.Handler) {
	s.endPoints = append(s.endPoints, EndPointData{
		requestPath:    requstPath,
		additionalData: handler,
	})
	s.multiplexer.Handle(requstPath, handler)
}

func (s *HTTPServer) addEndPoint(requestPath string, additionalData any) {
	s.endPoints = append(s.endPoints, EndPointData{
		requestPath:    requestPath,
		additionalData: additionalData,
		// hostPath:       hostPath,
	})

	s.multiplexer.HandleFunc(requestPath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Println("Serving: " + r.URL.Path)
		filePath := "." + r.URL.Path
		ext := filepath.Ext(filePath)

		w.Header().Set("Content-type", "text/"+ext[1:])

		compose(w, filePath, additionalData, func() {
			http.ServeFile(w, r, filePath)
		})
	})
}
