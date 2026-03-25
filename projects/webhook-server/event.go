package main

import (
	"encoding/json"
	"os"
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
