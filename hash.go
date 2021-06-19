package main

import "net/http"

func HashCreator(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func HashRetriever(w http.ResponseWriter, r *http.Request) {

}
