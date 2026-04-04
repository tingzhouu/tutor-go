package main

import (
	"net/http"
)

func main() {
	s := &server{eventsPath: "events.json"}
	http.HandleFunc("POST /events", logging(s.handleCreateEvent))

	http.HandleFunc("GET /events", logging(s.handleListEvents))

	http.HandleFunc("DELETE /events/{id}", logging(s.handleDeleteEvent))

	http.HandleFunc("PUT /events/{id}", logging(s.handleUpdateEvent))

	http.HandleFunc("GET /health", logging(s.handleHealth))

	http.ListenAndServe(":8080", nil)

}
