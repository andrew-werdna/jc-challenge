package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var (
	logger      *log.Logger
	addr        string
	PostTracker DataSet
)

func main() {
	PostTracker = DataSet{}.New()
	flag.StringVar(&addr, "p", ":8880", "port to listen on")
	flag.Parse()
	logger = log.New(os.Stdout, "http: ", log.LstdFlags)

	RegisterHandlers()
	logger.Println("starting server...")
	log.Fatal(http.ListenAndServe(addr, nil))
}

func RegisterHandlers() {
	http.HandleFunc("/hash", HashCreationHandler)
	http.HandleFunc("/hash/", HashRetriever)
	http.HandleFunc("/stats", StatsHandler)
	http.HandleFunc("/shutdown", Shutdown)
}
