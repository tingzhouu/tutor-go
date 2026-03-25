package main

import (
	"fmt"
	"net/http"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (sr *statusRecorder) WriteHeader(statusCode int) {
	sr.statusCode = statusCode
	sr.ResponseWriter.WriteHeader(statusCode)
}

func logging(next http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		recorder := &statusRecorder{ResponseWriter: w, statusCode: 200}

		next(recorder, r)
		elapsed := time.Since(now)
		fmt.Printf("%s %s %v %s\n", r.Method, r.URL.Path, recorder.statusCode, elapsed)
	}
}
