package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hash", HashHandler)
	http.HandleFunc("/stats", StatsHandler)
	log.Fatal(http.ListenAndServe(":8880", nil))
}
