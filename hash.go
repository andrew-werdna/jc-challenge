package main

import (
	"crypto/sha512"
	"encoding/base64"
	"net/http"
	"path"
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

	if len(password) != 1 {
		http.Error(w, "can only handle 1 password per POST request", http.StatusInternalServerError)
		logger.Println("can only handle 1 password per POST request")
		return
	}

	p := password[0]
	var key int

	defer func(t time.Time) {
		duration := time.Since(t)
		PostTracker.m.Lock()
		PostTracker.TrackedTime += int(duration.Microseconds())
		PostTracker.m.Unlock()
	}(start)

	PostTracker.m.Lock()
	PostTracker.NumPosts++
	key = PostTracker.NumPosts
	PostTracker.m.Unlock()

	go HashProcess(key, p, 5*time.Second)
	val := strconv.Itoa(PostTracker.NumPosts)
	w.Write([]byte(val))

}

func HashRetriever(w http.ResponseWriter, r *http.Request) {
	b := path.Base(r.URL.Path)
	key, err := strconv.Atoi(b)
	if err != nil {
		http.Error(w, "unable to convert path segment to number", http.StatusInternalServerError)
		logger.Println("unable to convert path segment to number")
		return
	}
	PostTracker.m.RLock()
	hash, ok := PostTracker.HashSet[key]
	PostTracker.m.RUnlock()
	if !ok {
		http.Error(w, "unable to retrieve hash", http.StatusNotFound)
		logger.Println("unable to retrieve hash")
		return
	}
	w.Write([]byte(hash))
}

func HashProcess(key int, password string, waitUntil time.Duration) {
	time.Sleep(waitUntil)
	hashed := HashEncode(password)
	PostTracker.m.Lock()
	PostTracker.HashSet[key] = hashed
	PostTracker.m.Unlock()
}

func HashEncode(v string) string {
	h := sha512.New()
	h.Write([]byte(v))
	var empty []byte
	result := h.Sum(empty)
	return base64.StdEncoding.EncodeToString(result)
}
