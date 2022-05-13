package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func bookTitleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]

	log.Println("Router:: bookTitleHandler:: Book-Title: " + title)
}

func bookInfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Router:: bookInfoHandler:: Book-Info:: URL:  ", r.URL.Path)

	vars := mux.Vars(r)

	book := Book{vars["title"], vars["author"]}

	resp, err := json.Marshal(book)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Router:: bookInfoHandler:: Book-Info:: response: ", string(resp))

	prepareResponse(w, resp)
}

func prepareResponse(w http.ResponseWriter, resp []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
