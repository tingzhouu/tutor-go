package main

import (
	"net/http"
)

func main() {
	s := &server{eventsPath: "events.json"}

	http.HandleFunc("/health", logging(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status": "ok"}`))
	}))

	http.HandleFunc("POST /events", logging(s.handleCreateEvent))

	http.HandleFunc("GET /events", logging(s.handleListEvents))

	http.HandleFunc("DELETE /events/{id}", logging(s.handleDeleteEvent))

	http.HandleFunc("PUT /events/{id}", logging(s.handleUpdateEvent))

	http.ListenAndServe(":8080", nil)

}
