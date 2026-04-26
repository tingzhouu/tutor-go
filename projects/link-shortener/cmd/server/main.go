package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"tutor-go/projects/link-shortener/internal/handler"
	"tutor-go/projects/link-shortener/internal/store"
)

func main() {
	s, err := store.New("links.db")
	if err != nil {
		fmt.Printf("error loading store %v\n", err)
		return
	}

	h := handler.Handler{Store: s}

	http.HandleFunc("POST /shorten", logging(h.Create))
	http.HandleFunc("GET /links", logging(h.GetAll))
	http.HandleFunc("GET /{code}", logging(h.Redirect))
	http.HandleFunc("DELETE /links/{id}", logging(h.Delete))

	server := http.Server{Addr: ":3000"}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("shutdown error", "err", err)
	}

	logger.Info("server stopped")
}
