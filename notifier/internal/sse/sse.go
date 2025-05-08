package sse

import (
	"fmt"
	"net/http"
	"time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Send a simple message every second
	for {
		// Write the event data to the response
		fmt.Fprintf(w, "data: %s\n\n", "Hello from SSE!")
		w.(http.Flusher).Flush() // Flush the response to the client

		time.Sleep(1 * time.Second) // Wait for 1 second before sending the next message
	}
}