package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type server struct {
	eventsPath string
}

func (s *server) handleDelay(w http.ResponseWriter, r *http.Request) {
	seconds := r.PathValue("seconds")
	n, err := strconv.Atoi(seconds)
	if err != nil || n <= 0 {
		http.Error(w, "seconds is not a valid integer", http.StatusBadRequest)
		return
	}
	time.Sleep(time.Second * time.Duration(n))
	fmt.Fprintf(w, "slept for %d seconds\n", n)
}

func (s *server) handleUpdateEvent(w http.ResponseWriter, r *http.Request) {
	var eventUpdate Event
	eventId := r.PathValue("id")
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&eventUpdate); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	events, err := loadEvents(s.eventsPath)
	if err != nil {
		http.Error(w, "failed to load events", http.StatusInternalServerError)
		return
	}

	var event *Event
	for i, existingEvent := range events {
		if existingEvent.ID == eventId {
			event = &events[i]
			break
		}
	}

	if event == nil {
		http.Error(w, "event not found", http.StatusNotFound)
		return
	}

	event.Type = eventUpdate.Type
	event.Amount = eventUpdate.Amount

	err = saveEvents(s.eventsPath, events)
	if err != nil {
		http.Error(w, "failed to save events", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(event)
}

func (s *server) handleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	events, err := loadEvents(s.eventsPath)
	if err != nil {
		http.Error(w, "failed to load events", http.StatusInternalServerError)
		return
	}
	eventResult := make([]Event, 0)
	for _, e := range events {
		if e.ID != id {
			eventResult = append(eventResult, e)
		}
	}

	if len(events) == len(eventResult) {
		http.Error(w, "no event found", http.StatusNotFound)
		return
	}

	err = saveEvents(s.eventsPath, eventResult)
	if err != nil {
		http.Error(w, "failed to save events", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
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

func checkService(ctx context.Context, serviceName string, delay time.Duration) healthResult {
	select {
	case <-ctx.Done():
		return healthResult{serviceName, "timeout"}
	case <-time.After(delay):
		return healthResult{serviceName, "healthy"}
	}
}

type healthResult struct {
	serviceName string
	status      string
}

func (s *server) handleHealth(w http.ResponseWriter, r *http.Request) {
	services := []struct {
		serviceName string
		delayMs     time.Duration
	}{
		{"database", 50 * time.Millisecond},
		{"cache", 500 * time.Millisecond},
		{"queue", 1500 * time.Millisecond},
	}

	ch := make(chan healthResult)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	var wg sync.WaitGroup

	for _, service := range services {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- checkService(ctx, service.serviceName, service.delayMs)
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	finalStatus := "healthy"
	checks := map[string]string{}

	for res := range ch {
		if res.status != "healthy" {
			finalStatus = "degraded"
		}
		checks[res.serviceName] = res.status
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	finalResult := struct {
		Checks map[string]string `json:"checks"`
		Status string            `json:"status"`
	}{
		Checks: checks,
		Status: finalStatus,
	}
	encoder.Encode(finalResult)
}
