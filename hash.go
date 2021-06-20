package main

import (
	"net/http"
	"strconv"
)

var numHashes int

func HashCreationHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	numHashes++
	val := strconv.Itoa(numHashes)
	w.Write([]byte(val))

}

func HashRetriever(w http.ResponseWriter, r *http.Request) {

}
