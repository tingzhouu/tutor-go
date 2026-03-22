package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Event struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Amount    int    `json:"amount"`
	CreatedAt string `json:"created_at"`
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status": "ok"}`))
	})

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var event Event
		fmt.Printf("event, before decoding: %v\n", event)
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&event); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		fmt.Printf("event, after decoding: %+v\n", event)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		encoder := json.NewEncoder(w)
		encoder.Encode(event)
	})

	http.ListenAndServe(":8080", nil)

}
