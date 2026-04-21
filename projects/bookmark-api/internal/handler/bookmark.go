package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tutor-go/projects/bookmark-api/internal/model"
	"tutor-go/projects/bookmark-api/internal/store"
)

type Handler struct {
	Store *store.Store
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var b model.Bookmark
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	b, err := h.Store.Create(b)

	if err != nil {
		http.Error(w, "error saving bookmark", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	encoder := json.NewEncoder(w)
	encoder.Encode(b)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	bookmarks, err := h.Store.GetAll()
	if err != nil {
		http.Error(w, "error fetching all bookmarks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(bookmarks)
}

func (h *Handler) GetOne(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	idInt, err := strconv.Atoi(idString)
	if err != nil || idInt <= 0 {
		http.Error(w, "invalid integer", http.StatusBadRequest)
		return
	}

	b, err := h.Store.GetOne(idInt)
	if err != nil {
		http.Error(w, "error fetching bookmark", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(b)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "invalid integer", http.StatusInternalServerError)
		return
	}
	err = h.Store.Delete(idInt)
	if err != nil {
		http.Error(w, "unable to delete record", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
