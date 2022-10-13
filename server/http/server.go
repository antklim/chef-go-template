package http

import (
	handler "chef-go-template/handler/http"
	"log"
	"net/http"
)

const defaultAddress = ":8080"

func Start() {
	s := &http.Server{
		Addr:    defaultAddress,
		Handler: handler.Mux(),
	}

	log.Printf("HTTP server is listening at %s", defaultAddress)
	log.Fatalf("HTTP server stopped: %v", s.ListenAndServe())
}
