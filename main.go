package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var (
	logger      *log.Logger
	port        string
	PostTracker DataSet
)

func main() {
	PostTracker = DataSet{}.New()
	flag.StringVar(&port, "p", ":8880", "port to listen on")
	flag.Parse()
	logger = log.New(os.Stdout, "http: ", log.LstdFlags)

	RegisterHandlers()
	log.Fatal(http.ListenAndServe(port, nil))
}

func RegisterHandlers() {
	http.HandleFunc("/hash", HashCreationHandler)
	http.HandleFunc("/hash/", HashRetriever)
	http.HandleFunc("/stats", StatsHandler)
	http.HandleFunc("/shutdown", Shutdown)
}
