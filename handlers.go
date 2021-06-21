package main

import (
	"net/http"
)

func RegisterHandlers() {
	http.HandleFunc("/hash", HashCreationHandler)
	http.HandleFunc("/hash/", HashRetriever)
	http.HandleFunc("/stats", StatsHandler)
	http.HandleFunc("/shutdown", Shutdown)
}

func Shutdown(w http.ResponseWriter, r *http.Request) {

}
