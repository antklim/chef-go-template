package http

import (
	"fmt"
	"net/http"
)

const healthRoute = "/health"

func init() {
	router.Handle(healthRoute, healthHandler())
}

func healthHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})
}
