package main

import (
	"encoding/json"
	"net/http"
)

// The Stat type is used to return the statistics for the number of POST requests
// and the average time spent processing those requests NOT including the long running
// goroutine that creates the hash.
type Stat struct {
	Total   int `json:"total"`
	Average int `json:"average"`
}

// StatsHandler is an http.Handler bound to the /stats endpoint and returns
// a JSON object marshalled from an instance of the Stat structure above.
func StatsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	s := Stat{}
	RequestInfo.m.RLock()
	s.Total = RequestInfo.Posts
	trackedTime := RequestInfo.TrackedTime
	RequestInfo.m.RUnlock()

	if s.Total == 0 {
		http.Error(w, "no posts or hashes yet", http.StatusForbidden)
		return
	}

	s.Average = trackedTime / s.Total
	data, err := json.Marshal(s)

	if err != nil {
		http.Error(w, "something went wrong marshalling json resonse", http.StatusInternalServerError)
		logger.Println("something went wrong marshalling json resonse")
		return
	}
	w.Write(data)

}
