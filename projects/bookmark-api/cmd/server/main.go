package main

import (
	"fmt"
	"net/http"
	"time"
	"tutor-go/projects/bookmark-api/internal/handler"
	"tutor-go/projects/bookmark-api/internal/store"
)

func delay(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 3)
}

func main() {
	s, err := store.New("bookmarks.db")
	if err != nil {
		fmt.Printf("error loading store %v\n", err)
		return
	}
	h := handler.Handler{Store: s}

	http.HandleFunc("GET /delay", delay)
	http.HandleFunc("POST /bookmarks", h.Create)
	http.HandleFunc("GET /bookmarks", h.GetAll)
	http.HandleFunc("GET /bookmarks/{id}", h.GetOne)
	http.HandleFunc("DELETE /bookmarks/{id}", h.Delete)

	server := http.Server{Addr: ":3000"}

	server.ListenAndServe()
}
