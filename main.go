package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/books/{title}", bookTitleHandler).Methods("POST")

	http.ListenAndServe(":9200", route)
}

// func bookTitleHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	title := vars["title"]

// 	log.Println("Router:: Book-Title: " + title)
// }

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}
