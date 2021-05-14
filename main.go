package main

import (
	"net/http"

	"github.com/gorilla/mux"

	router "poe-debugserver/src/router"
)

var port = ":9999"

func main() {
	r := mux.NewRouter()
	router.SetRoute(r)

	http.ListenAndServe(port, r)
}
