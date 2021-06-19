package main

import (
	"log"
	"net/http"
)

func main() {
	RegisterHandlers()
	log.Fatal(http.ListenAndServe(":8880", nil))
}
