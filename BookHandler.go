package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func bookTitleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]

	log.Println("Router:: Book-Title: " + title)
}
