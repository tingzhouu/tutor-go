package main

import (
	"encoding/json"
	"os"
	"time"
)

// Event represents a webhook event.
// Struct tags control JSON serialization — like a built-in class-transformer.
type Event struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Amount    int    `json:"amount"`
	CreatedAt string `json:"created_at"`
}

const eventsFile = "events.json"

// loadEvents reads all events from the JSON file.
// Returns an empty slice if the file doesn't exist yet.
func loadEvents() ([]Event, error) {
	data, err := os.ReadFile(eventsFile)
	if err != nil {
		// File doesn't exist yet — that's fine, return empty slice
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

// saveEvents writes all events to the JSON file.
// 0644 = owner can read/write, others can read (standard file permission).
func saveEvents(events []Event) error {
	data, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(eventsFile, data, 0644)
}

// newEvent creates an Event with the current timestamp.
func newEvent(id, eventType string, amount int) Event {
	return Event{
		ID:        id,
		Type:      eventType,
		Amount:    amount,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}
