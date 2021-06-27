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

type RequestData struct {
	Posts       int
	Hashes      map[int]string
	TrackedTime int // time spent processing /hash POST requests in microseconds NOT including the hash function with 5 second delay
	m           *sync.RWMutex
}

type HashArgs struct {
	key       int
	password  string
	waitUntil time.Duration
	wg        *sync.WaitGroup
}

func (d RequestData) New() RequestData {
	return RequestData{
		Posts:       0,
		Hashes:      make(map[int]string),
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
		RequestInfo.m.Lock()
		RequestInfo.TrackedTime += int(duration.Microseconds())
		RequestInfo.m.Unlock()
	}(start)

	RequestInfo.m.Lock()
	RequestInfo.Posts++
	key = RequestInfo.Posts
	RequestInfo.m.Unlock()

	a := HashArgs{
		key:       key,
		password:  p,
		waitUntil: 5 * time.Second,
		wg:        &WG,
	}

	WG.Add(1)
	go HashProcess(a)
	val := strconv.Itoa(RequestInfo.Posts)
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
	RequestInfo.m.RLock()
	hash, ok := RequestInfo.Hashes[key]
	RequestInfo.m.RUnlock()
	if !ok {
		http.Error(w, "unable to retrieve hash", http.StatusNotFound)
		logger.Println("unable to retrieve hash")
		return
	}
	w.Write([]byte(hash))
}

func HashProcess(ha HashArgs) {
	defer ha.wg.Done()
	time.Sleep(ha.waitUntil)
	hashed := HashEncode(ha.password)
	RequestInfo.m.Lock()
	RequestInfo.Hashes[ha.key] = hashed
	RequestInfo.m.Unlock()
}

func HashEncode(v string) string {
	h := sha512.New()
	h.Write([]byte(v))
	var empty []byte
	result := h.Sum(empty)
	return base64.StdEncoding.EncodeToString(result)
}
