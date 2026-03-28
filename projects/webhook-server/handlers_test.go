package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleListEvents(t *testing.T) {
	path := t.TempDir() + "/events.json"
	s := &server{
		eventsPath: path,
	}

	events := []Event{
		{ID: "evt_1", Type: "payment_created", Amount: 100, CreatedAt: "2026-03-23T20:08:21+08:00"},
		{ID: "evt_2", Type: "payment_created", Amount: 200, CreatedAt: "2026-03-22T20:08:21+08:00"},
	}
	saveEvents(path, events)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/events", nil)

	s.handleListEvents(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Received %d status code, expected %d", rec.Code, http.StatusOK)
	}

	expectedResponseBody := `[{"id":"evt_1","type":"payment_created","amount":100,"created_at":"2026-03-23T20:08:21+08:00"},{"id":"evt_2","type":"payment_created","amount":200,"created_at":"2026-03-22T20:08:21+08:00"}]`
	actualResponseBody := strings.TrimSpace(rec.Body.String())
	if actualResponseBody != expectedResponseBody {
		t.Errorf("Received response %s, expected %s", actualResponseBody, expectedResponseBody)

	}
}

func TestHandleCreateEvent(t *testing.T) {
	path := t.TempDir() + "/events.json"
	s := &server{
		eventsPath: path,
	}

	rec := httptest.NewRecorder()
	body := strings.NewReader(`{"id":"evt_1","type":"payment_created","amount":500}`)
	req := httptest.NewRequest(http.MethodPost, "/events", body)

	s.handleCreateEvent(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("Received %d status code, expected %d", rec.Code, http.StatusCreated)
	}

	var event Event
	json.NewDecoder(rec.Body).Decode(&event)
	if event.CreatedAt == "" {
		t.Errorf("Received empty created_at")
	}

	events, err := loadEvents(path)
	if err != nil {
		t.Fatalf("failed to load events %v", err)
	}

	if len(events) != 1 {
		t.Errorf("Received %d events, expected 1", len(events))
	}
	if events[0].ID != "evt_1" {
		t.Errorf("Received %s id, expected %s", events[0].ID, "evt_1")
	}
}
