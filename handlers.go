package main

import (
	"crypto/sha512"
	"encoding/base64"
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

func HashEncode(v string) string {
	h := sha512.New()
	h.Write([]byte(v))
	var empty []byte
	result := h.Sum(empty)
	return base64.StdEncoding.EncodeToString(result)
}
