package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"sync"
	"time"
)

// RequestData is a type intended for keeping track of the number of POST requests for hashes,
// the association between POST number as a key and the actual hash as value,
// and the tracked time for the /stats endpoint.
type RequestData struct {
	Posts       int
	Hashes      map[int]string
	TrackedTime int
	m           *sync.RWMutex
}

// New() is a convenience function for creating a new struct to track what we care about.
// See RequestData struct type for the information we care to track.
func (d RequestData) New() RequestData {
	return RequestData{
		Posts:       0,
		Hashes:      make(map[int]string),
		TrackedTime: 0,
		m:           &sync.RWMutex{},
	}
}

// HashArgs is a type meant to make a long argument list to the HashProcess func
// shorter and easier to work with.
type HashArgs struct {
	key       int
	password  string
	waitUntil time.Duration
	wg        *sync.WaitGroup
}

// HashCreationHandler is bound to the /hash endpoint and returns an incrementing
// integer number immediately, while kicking off a background process that hashes
// the passed in password. To simulate a long running process the goroutine sleeps
// for five seconds before doing its work.
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
		waitUntil: waitTime,
		wg:        &WG,
	}

	WG.Add(1)
	go HashProcess(a)
	val := strconv.Itoa(RequestInfo.Posts)
	fmt.Fprintln(w, val)

}

// HashRetriever is an http.Handler bound to the /hash/{/[1-9]+/s} (regex between curly braces and delimited with slashes) endpoint.
// This handler receives the number in the url returned from the /hash endpoint and returns the hashed, base64 encoded password that
// was also given to the server on the /hash endpoint.
func HashRetriever(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		logger.Println("ERROR: not a post request")
		return
	}

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

// HashProcess is a helper func ready to be run in another goroutine. It will update our storage object
// with the integer key returned to the user, and the hashed, base64 encoded password the user passed in.
func HashProcess(ha HashArgs) {
	defer ha.wg.Done()
	time.Sleep(ha.waitUntil)
	hashed := HashEncode(ha.password)
	RequestInfo.m.Lock()
	RequestInfo.Hashes[ha.key] = hashed
	RequestInfo.m.Unlock()
}

// HashEncode is a helper func that takes a password string, hashes it with sha512 encoding scheme,
// and then returns the base64 string encoding of the hashed result.
func HashEncode(v string) string {
	h := sha512.New()
	h.Write([]byte(v))
	var empty []byte
	result := h.Sum(empty)
	return base64.StdEncoding.EncodeToString(result)
}
