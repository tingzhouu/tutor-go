package main

import (
	"os"
	"strconv"
	"time"
)

type config struct {
	Port            string
	EventsPath      string
	ShutdownTimeout time.Duration
}

func loadConfig() config {
	port := os.Getenv("PORT")
	eventsPath := os.Getenv("EVENTS_PATH")
	shutdownTimeoutStr := os.Getenv("SHUTDOWN_TIMEOUT_SECONDS")
	var shutdownTimeoutDuration time.Duration
	if port == "" {
		port = "8080"
	}
	if eventsPath == "" {
		eventsPath = "events.json"
	}

	res, err := strconv.Atoi(shutdownTimeoutStr)
	if err != nil || res <= 0 {
		shutdownTimeoutDuration = time.Duration(10) * time.Second
	} else {
		shutdownTimeoutDuration = time.Duration(res) * time.Second
	}

	return config{
		Port:            port,
		EventsPath:      eventsPath,
		ShutdownTimeout: shutdownTimeoutDuration,
	}
}
