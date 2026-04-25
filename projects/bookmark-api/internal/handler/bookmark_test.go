package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"tutor-go/projects/bookmark-api/internal/model"
)

type fakeStore struct {
	bookmarks []model.Bookmark
}

func (f *fakeStore) Create(b model.Bookmark) (model.Bookmark, error) {
	b.Id = len(f.bookmarks) + 1
	f.bookmarks = append(f.bookmarks, b)
	return b, nil
}

func (f *fakeStore) GetAll() ([]model.Bookmark, error) {
	return f.bookmarks, nil
}

func (f *fakeStore) GetOne(id int) (model.Bookmark, error) {
	for _, b := range f.bookmarks {
		if b.Id == id {
			return b, nil
		}
	}
	return model.Bookmark{}, errors.New("not found")
}

func (f *fakeStore) Delete(id int) error {
	return nil
}

func TestHandlerCreate(t *testing.T) {
	f := &fakeStore{}
	h := &Handler{Store: f}

	body := strings.NewReader(`{"url":"https://go.dev","title":"Go docs","tags":"golang"}`)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/bookmarks", body)

	h.Create(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("Unexpected status code %d", rec.Code)
	}

	if strings.TrimSpace(rec.Body.String()) != `{"id":1,"url":"https://go.dev","title":"Go docs","tags":"golang"}` {
		t.Errorf("Unexpected response %s", rec.Body.String())
	}

	if len(f.bookmarks) != 1 {
		t.Errorf("Unexpected length of bookmarks %d", len(f.bookmarks))
	}
}
