package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Text   string `json:"quote"`
}

var quotes []Quote
var curID int

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/quotes", addQuote).Methods("POST")
	router.HandleFunc("/quotes", getAllQuotes).Methods("GET")
	router.HandleFunc("/quotes/random", getRandomQuote).Methods("GET")
	router.HandleFunc("/quotes", getQuotesByAuthor).Methods("GET")
	router.HandleFunc("/quotes/{id}", deleteQuoteByID).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}

func addQuote(w http.ResponseWriter, r *http.Request) {
	var q Quote
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	curID += 1
	q.ID = curID
	quotes = append(quotes, q)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(q)

}

func getAllQuotes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(quotes)
}

func getRandomQuote(w http.ResponseWriter, r *http.Request) {
	if len(quotes) == 0 {
		http.Error(w, "No quotes found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(quotes[rand.Intn(len(quotes))])
}

func getQuotesByAuthor(w http.ResponseWriter, r *http.Request) {
	if len(quotes) == 0 {
		http.Error(w, "No quotes found", http.StatusNotFound)
		return
	}

	author := r.URL.Query().Get("author")

	if author != "" {
		var quotesByAuthor []Quote
		for _, q := range quotes {
			if q.Author == author {
				quotesByAuthor = append(quotesByAuthor, q)
			}
		}
		json.NewEncoder(w).Encode(quotesByAuthor)
		return
	}

	json.NewEncoder(w).Encode(quotes)
}

func deleteQuoteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, q := range quotes {
		if q.ID == id {
			quotes[i] = quotes[len(quotes)-1]
			quotes = quotes[:len(quotes)-1]
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "Quote not found", http.StatusNotFound)
}
