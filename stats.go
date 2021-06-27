package main

import (
	"encoding/json"
	"net/http"
)

type Stat struct {
	NumPosts    int `json:"total"`
	AverageTime int `json:"average"`
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	s := Stat{}
	PostTracker.m.RLock()
	s.NumPosts = PostTracker.NumPosts
	s.AverageTime = PostTracker.TrackedTime / s.NumPosts
	PostTracker.m.RUnlock()
	data, err := json.Marshal(s)

	if err != nil {
		w.Write([]byte("something went wrong marshalling json resonse"))
		w.WriteHeader(http.StatusInternalServerError)
		logger.Println("something went wrong marshalling json resonse")
		return
	}
	w.Write(data)

}
