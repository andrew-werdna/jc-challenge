package main

import (
	"crypto/sha512"
	"encoding/base64"
	"net/http"
	"strconv"
)

var numHashes int

type DataSet struct {
	NumPosts int
	HashSet  map[int]string
}

func HashCreationHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parseFormErr := r.ParseForm()
	if parseFormErr != nil {
		http.Error(w, "unable to parse body as form", http.StatusInternalServerError)
		return
	}

	password, ok := r.Form["password"]
	if !ok {
		http.Error(w, "password field not found", http.StatusBadRequest)
		return
	}

	if len(password) == 1 {
		p := password[0]
		logger.Print(p)
		/**
		* TODO: write tests for sending the password into the other goroutine
		 */
	}

	numHashes++
	val := strconv.Itoa(numHashes)
	w.Write([]byte(val))

}

func HashRetriever(w http.ResponseWriter, r *http.Request) {

}

func HashEncode(v string) string {
	h := sha512.New()
	h.Write([]byte(v))
	var empty []byte
	result := h.Sum(empty)
	return base64.StdEncoding.EncodeToString(result)
}
