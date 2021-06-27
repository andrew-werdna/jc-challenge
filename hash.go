package main

import (
	"crypto/sha512"
	"encoding/base64"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type DataSet struct {
	NumPosts    int
	HashSet     map[int]string
	TrackedTime int // time spent processing /hash POST requests in microseconds
	m           *sync.RWMutex
}

type NumHash struct {
	Key   int
	Value string
}

func (d DataSet) New() DataSet {
	return DataSet{
		NumPosts:    0,
		HashSet:     make(map[int]string),
		TrackedTime: 0,
		m:           &sync.RWMutex{},
	}
}

func HashCreationHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		logger.Println("ERROR: not a post request")
		return
	}

	parseFormErr := r.ParseForm()
	if parseFormErr != nil {
		http.Error(w, "unable to parse body as form", http.StatusInternalServerError)
		logger.Println("ERROR: unable to parse body as form")
		return
	}

	password, ok := r.Form["password"]
	if !ok {
		http.Error(w, "password field not found", http.StatusBadRequest)
		logger.Println("ERROR: password field not found")
		return
	}

	defer func(t time.Time) {
		duration := time.Since(t)
		PostTracker.m.Lock()
		PostTracker.TrackedTime += int(duration.Microseconds())
		PostTracker.m.Unlock()
	}(start)

	if len(password) == 1 {
		p := password[0]
		logger.Print(p)
		/**
		* TODO: write tests for sending the password into the other goroutine
		 */
	}

	PostTracker.NumPosts++
	val := strconv.Itoa(PostTracker.NumPosts)
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
