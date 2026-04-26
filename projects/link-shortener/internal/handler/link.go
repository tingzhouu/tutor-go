package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tutor-go/projects/link-shortener/internal/model"
)

type Handler struct {
	Store LinkStore
}

type LinkStore interface {
	Create(l model.Link) (model.Link, error)
	GetOne(shortCode string) (model.Link, error)
	GetAll() ([]model.Link, error)
	Delete(id int) error
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var l model.Link
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&l); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	l.ShortCode = randomString(6)

	l, err := h.Store.Create(l)

	if err != nil {
		http.Error(w, "error saving link", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	encoder := json.NewEncoder(w)
	encoder.Encode(l)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortCode := r.PathValue("code")
	l, err := h.Store.GetOne(shortCode)
	if err != nil {
		http.Error(w, "error fetching link", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, l.Url, http.StatusFound)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	links, err := h.Store.GetAll()

	if err != nil {
		http.Error(w, "error getting links", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(links)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	idInt, err := strconv.Atoi(idStr)

	if err != nil || idInt <= 0 {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.Store.Delete(idInt)
	if err != nil {
		http.Error(w, "error deleting link", http.StatusInternalServerError)
		return
	}
}
