package main

import (
	"context"
	"net/http"
	"time"
)

func Shutdown(w http.ResponseWriter, r *http.Request) {
	logger.Println("Server is shutting down...")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Shutdown Started..."))

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
