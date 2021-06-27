package main

import (
	"encoding/json"
	"net/http"
)

type Stat struct {
	Total   int `json:"total"`
	Average int `json:"average"`
}

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
