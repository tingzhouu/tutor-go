package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type server struct {
	eventsPath string
}

func (s *server) handleCreateEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&event); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	event.CreatedAt = time.Now().Format(time.RFC3339)
	events, err := loadEvents(s.eventsPath)
	if err != nil {
		http.Error(w, "error loading json", http.StatusInternalServerError)
		return
	}
	events = append(events, event)
	err = saveEvents(s.eventsPath, events)
	if err != nil {
		http.Error(w, "error saving events", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	encoder := json.NewEncoder(w)
	encoder.Encode(event)
}

func (s *server) handleListEvents(w http.ResponseWriter, r *http.Request) {
	eventType := r.URL.Query().Get("type")
	events, err := loadEvents(s.eventsPath)
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
