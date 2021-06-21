package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var (
	logger *log.Logger
	port   string
)

func main() {
	flag.StringVar(&port, "p", ":8880", "port to listen on")
	flag.Parse()
	logger = log.New(os.Stdout, "http: ", log.LstdFlags)

	RegisterHandlers()
	log.Fatal(http.ListenAndServe(port, nil))
}
