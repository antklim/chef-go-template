package http

import (
	"chef-go-template/app/echo"
	"fmt"
	"net/http"
)

const echoRoute = "/echo"

func init() {
	router.Handle(echoRoute, echoHandler())
}

func echoHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, echo.Reply())
	})
}
