package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"tutor-go/projects/link-shortener/internal/model"
)

type fakeStore struct {
	links []model.Link
}

func (f *fakeStore) Create(l model.Link) (model.Link, error) {
	l.Id = len(f.links) + 1
	f.links = append(f.links, l)
	return l, nil
}

func (f *fakeStore) GetOne(shortCode string) (model.Link, error) {
	for _, link := range f.links {
		if link.ShortCode == shortCode {
			return link, nil
		}
	}
	return model.Link{}, errors.New("Not Found")
}
func (f *fakeStore) Delete(id int) error {
	res := []model.Link{}

	for _, link := range f.links {
		if link.Id != id {
			res = append(res, link)
		}
	}

	f.links = res
	return nil
}

func (f *fakeStore) GetAll() ([]model.Link, error) {
	return []model.Link{
		{Id: 1, ShortCode: "shortcode1", Url: "url1"},
		{Id: 2, ShortCode: "shortcode2", Url: "url2"},
		{Id: 3, ShortCode: "shortcode3", Url: "url3"},
	}, nil
}

func TestHandlerCreate(t *testing.T) {
	f := &fakeStore{}
	h := &Handler{Store: f}

	body := strings.NewReader(`{"url":"https://myverylongurl.com"}`)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/shorten", body)

	h.Create(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("Unexpected status code %d", rec.Code)
	}

	var link model.Link
	json.NewDecoder(rec.Body).Decode(&link)

	if link.Url != "https://myverylongurl.com" {
		t.Errorf("unexpected url: %s", link.Url)
	}
	if len(link.ShortCode) != 6 {
		t.Errorf("unexpected short code length: %d", len(link.ShortCode))
	}
	if link.Id != 1 {
		t.Errorf("unexpected id: %d", link.Id)
	}
	if len(f.links) != 1 {
		t.Errorf("Unexpected length of bookmarks %d", len(f.links))
	}
}

func TestHandlerRedirect(t *testing.T) {
	f := &fakeStore{}
	h := &Handler{Store: f}

	f.Create(model.Link{Url: "https://my-special-url.com", ShortCode: "abc123"})

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/abc123", nil)
	req.SetPathValue("code", "abc123")

	h.Redirect(rec, req)

	if rec.Code != http.StatusFound {
		t.Errorf("Unexpected status code %d", rec.Code)
	}

	if rec.Header().Get("Location") != "https://my-special-url.com" {
		t.Errorf("Unexpcted location header %s", rec.Header().Get("Location"))
	}
}

func TestHandlerDelete(t *testing.T) {
	f := &fakeStore{}
	h := &Handler{Store: f}

	f.Create(model.Link{Url: "https://my-special-url.com"})
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/1", nil)
	req.SetPathValue("id", "1")

	h.Delete(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Unexpected status code %d", rec.Code)
	}
}
