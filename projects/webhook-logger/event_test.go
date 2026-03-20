package main

import (
	"testing"
)

func TestNewEvent(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		eventType string
		amount    int
	}{
		{"valid payment event", "evt_123", "payment_intent.succeeded", 4999},
		{"zero amount", "evt_456", "charge.refunded", 0},
		{"negative amount", "evt_456", "charge.refunded", -5000},
		{"empty string", "", "", -5000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := tt.id
			amount := tt.amount
			eventType := tt.eventType
			event := newEvent(id, eventType, amount)
			if event.ID != id {
				t.Errorf("Received: %s, Expected: %s", event.ID, id)
			}
			if event.Amount != amount {
				t.Errorf("Received: %d, Expected: %d", event.Amount, amount)
			}
			if event.Type != eventType {
				t.Errorf("Received: %s, Expected: %s", event.Type, eventType)
			}
			if event.CreatedAt == "" {
				t.Errorf("Received: %s, Expected: non empty", event.CreatedAt)
			}
		})
	}
}

func TestSaveAndLoadEvents(t *testing.T) {
	dir := t.TempDir()
	path := dir + "/events.json"
	events := make([]Event, 0)
	rawEvents := []struct {
		name      string
		id        string
		eventType string
		amount    int
	}{
		{"valid payment event", "evt_123", "payment_intent.succeeded", 4999},
		{"zero amount", "evt_456", "charge.refunded", 0},
		{"negative amount", "evt_456", "charge.refunded", -5000},
		{"empty string", "", "", -5000},
	}
	for _, rawEvent := range rawEvents {
		events = append(events, newEvent(rawEvent.id, rawEvent.eventType, rawEvent.amount))
	}
	saveEvents(path, events)

	loadedEvents, err := loadEvents(path)
	if err != nil {
		t.Fatalf("Error %v when loading event, expected nil", err)
	}
	if len(loadedEvents) != len(rawEvents) {
		t.Errorf("Loaded %d events, expected %d events", len(loadedEvents), len(rawEvents))
	}
	for i := range rawEvents {
		if loadedEvents[i].ID != rawEvents[i].id {
			t.Errorf("ID mismatch, actual: %s, expected %s", loadedEvents[i].ID, rawEvents[i].id)
		}
		if loadedEvents[i].Type != rawEvents[i].eventType {
			t.Errorf("Type mismatch, actual: %s, expected %s", loadedEvents[i].Type, rawEvents[i].eventType)
		}
		if loadedEvents[i].Amount != rawEvents[i].amount {
			t.Errorf("Amount mismatch, actual: %d, expected %d", loadedEvents[i].Amount, rawEvents[i].amount)
		}
	}
}

func TestLoadEventsFileNotFound(t *testing.T) {
	dir := t.TempDir()
	path := dir + "/events.json"

	loadedEvents, err := loadEvents(path)
	// loadedEvents, err := loadEvents(path)
	if err != nil {
		t.Errorf("expected nil error but got %v", err)
	}

	if len(loadedEvents) != 0 {
		t.Errorf("expected loadedEvents to be empty but got %v", loadedEvents)
	}
}
