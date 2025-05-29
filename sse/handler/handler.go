package handler

import (
	"log"
	"net/http"
	"time"
)

func SSEHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	notify := r.Context().Done()

	for {
		select {
		case <-notify:
			log.Println("Client disconnected")
			return
		default:
			event := "data: " + time.Now().Format(time.RFC3339) + "\n\n"
			if _, err := w.Write([]byte(event)); err != nil {
				log.Println("Write error:", err)
				return
			}
			flusher.Flush()
			time.Sleep(1 * time.Second)
		}
	}
}
