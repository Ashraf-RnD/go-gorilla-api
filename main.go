package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	bookRouter := route.PathPrefix("/books").Subrouter()

	bookRouter.HandleFunc("/{title}", bookTitleHandler).Methods("POST")
	bookRouter.HandleFunc("/title/{title}/author/{author}", bookInfoHandler).Methods("GET")

	http.Handle("/", route)
	http.ListenAndServe(":9200", route)
}
