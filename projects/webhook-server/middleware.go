package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

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
		reqLogger := logger.With("method", r.Method, "path", r.URL.Path, "statusCode", recorder.statusCode, "elapsed", elapsed.Milliseconds())
		if recorder.statusCode >= 500 {
			reqLogger.Error("request handled")
		} else if recorder.statusCode >= 400 {
			reqLogger.Warn("request handled")
		} else {
			reqLogger.Info("request handled")

		}
	}
}
