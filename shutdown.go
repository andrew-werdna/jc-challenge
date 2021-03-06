package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Shutdown is the http.Handler responsible for shutting down the server and
// initiating it in a graceful fashion.
func Shutdown(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		logger.Println("ERROR: not a post request")
		return
	}

	logger.Println("Server is shutting down...")
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, "Shutdown Started...")

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		time.Sleep(200 * time.Millisecond)
		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		ctx.Done()
	}()

}
