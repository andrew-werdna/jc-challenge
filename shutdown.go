package main

import (
	"context"
	"net/http"
	"time"
)

func Shutdown(w http.ResponseWriter, r *http.Request) {
	logger.Println("Server is shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	w.WriteHeader(http.StatusAccepted)
	ctx.Done()
}
