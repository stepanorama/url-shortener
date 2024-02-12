package server

import (
	"github.com/stepanorama/url-shortener/internal/app/handlers"
	"net/http"
)

func RunServer() {
	// http.NewServeMux() creates a new server instance instead of the default server.
	mux := http.NewServeMux()
	// HandleFunc receives MainHandler function.
	// MainHandler's signature matches the one for ServeHTTP() method of http.Handler
	// so MainHandler can be type-casted down the road with http.HandlerFunc(), which
	// implements http.Handler interface.
	mux.HandleFunc("/", handlers.MainHandler)
	
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
