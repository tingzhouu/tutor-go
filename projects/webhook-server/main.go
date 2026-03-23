package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Event struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Amount    int    `json:"amount"`
	CreatedAt string `json:"created_at"`
}

func saveEvents(path string, events []Event) error {
	data, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func loadEvents(path string) ([]Event, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Event{}, nil
		}
		return nil, err
	}

	var events []Event
	err = json.Unmarshal(data, &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func handleCreateEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	fmt.Printf("event, before decoding: %v\n", event)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&event); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	event.CreatedAt = time.Now().Format(time.RFC3339)
	events, err := loadEvents("events.json")
	if err != nil {
		http.Error(w, "error loading json", http.StatusInternalServerError)
		return
	}
	events = append(events, event)
	err = saveEvents("events.json", events)
	if err != nil {
		http.Error(w, "error saving events", http.StatusInternalServerError)
		return
	}

	fmt.Printf("event, after decoding: %+v\n", event)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	encoder := json.NewEncoder(w)
	encoder.Encode(event)
}

func handleListEvents(w http.ResponseWriter, r *http.Request) {
	eventType := r.URL.Query().Get("type")
	events, err := loadEvents("events.json")
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	eventsResult := make([]Event, 0)
	for _, event := range events {
		if eventType == "" || event.Type == eventType {
			eventsResult = append(eventsResult, event)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(eventsResult)
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status": "ok"}`))
	})

	http.HandleFunc("POST /events", handleCreateEvent)

	http.HandleFunc("GET /events", handleListEvents)

	http.ListenAndServe(":8080", nil)

}
