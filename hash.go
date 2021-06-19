package main

import "net/http"

type HashData struct {
	Posts  int
	Hashes map[int]string
}

func HashCreator(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("1"))

}

func HashRetriever(w http.ResponseWriter, r *http.Request) {

}
