package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s := &server{eventsPath: "events.json"}
	http.HandleFunc("POST /events", logging(s.handleCreateEvent))

	http.HandleFunc("GET /events", logging(s.handleListEvents))

	http.HandleFunc("DELETE /events/{id}", logging(s.handleDeleteEvent))

	http.HandleFunc("PUT /events/{id}", logging(s.handleUpdateEvent))

	http.HandleFunc("GET /health", logging(s.handleHealth))

	http.HandleFunc("GET /delay/{seconds}", logging(s.handleDelay))

	server := &http.Server{Addr: ":8080"}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("shutdown error", "err", err)
	}

	logger.Info("server stopped")

}
