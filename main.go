package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"sync"
)

var (
	logger      *log.Logger
	addr        string
	RequestInfo RequestData
	server      *http.Server
	WG          sync.WaitGroup
)

func init() {
	RequestInfo = RequestData{}.New()
	logger = log.New(os.Stdout, "http: ", log.LstdFlags)
}

func main() {
	flag.StringVar(&addr, "p", ":8880", "port to listen on")
	flag.Parse()

	router := http.NewServeMux()
	RegisterHandlers(router)

	server = &http.Server{
		Addr:     addr,
		Handler:  router,
		ErrorLog: logger,
	}

	logger.Println("starting server...")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", addr, err)
	}

	WG.Wait()
	logger.Println("server stopped")
}

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/hash", HashCreationHandler)
	mux.HandleFunc("/hash/", HashRetriever)
	mux.HandleFunc("/stats", StatsHandler)
	mux.HandleFunc("/shutdown", Shutdown)
}
