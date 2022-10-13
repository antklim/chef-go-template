package http

import "net/http"

var router = http.NewServeMux()

func Mux() *http.ServeMux {
	return router
}
