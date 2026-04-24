package store

import (
	"testing"
	"tutor-go/projects/bookmark-api/internal/model"
)

func TestCreate(t *testing.T) {
	store, err := New(t.TempDir() + "/test.db")
	if err != nil {
		t.Fatalf("could not create store %v", err)
	}

	b := model.Bookmark{
		Url:   "mockurl",
		Title: "mocktitle",
		Tags:  "mocktags",
	}

	b, err = store.Create(b)
	if err != nil {
		t.Fatalf("could not create bookmark record %v", err)
	}

	if b.Id <= 0 {
		t.Errorf("Unexpected id %d", b.Id)
	}
}

func TestGetAll(t *testing.T) {
	store, err := New(t.TempDir() + "/test.db")
	if err != nil {
		t.Fatalf("could not create store %v", err)
	}

	b := model.Bookmark{
		Url:   "mockurl",
		Title: "mocktitle",
		Tags:  "mocktags",
	}

	b, err = store.Create(b)
	if err != nil {
		t.Fatalf("could not create bookmark record %v", err)
	}

	b, err = store.Create(b)
	if err != nil {
		t.Fatalf("could not create bookmark record %v", err)
	}

	bookmarks, err := store.GetAll()

	if err != nil {
		t.Fatalf("could not get all bookmark records %v", err)
	}

	if len(bookmarks) != 2 {
		t.Errorf("Unexpected len %d", len(bookmarks))
	}
}

func TestGetOne(t *testing.T) {
	store, err := New(t.TempDir() + "/test.db")
	if err != nil {
		t.Fatalf("could not create store %v", err)
	}

	b := model.Bookmark{
		Url:   "mockurl",
		Title: "mocktitle",
		Tags:  "mocktags",
	}

	b, err = store.Create(b)
	if err != nil {
		t.Fatalf("could not create bookmark record %v", err)
	}

	bookmark, err := store.GetOne(b.Id)

	if err != nil {
		t.Fatalf("could not get all bookmark records %v", err)
	}

	if bookmark.Id != b.Id || bookmark.Tags != b.Tags || bookmark.Title != b.Title || bookmark.Url != b.Url {
		t.Errorf("bookmark records don't match bookmark: %v b: %v", bookmark, b)
	}
}

func TestDelete(t *testing.T) {
	store, err := New(t.TempDir() + "/test.db")
	if err != nil {
		t.Fatalf("could not create store %v", err)
	}

	b := model.Bookmark{
		Url:   "mockurl",
		Title: "mocktitle",
		Tags:  "mocktags",
	}

	b, err = store.Create(b)
	if err != nil {
		t.Fatalf("could not create bookmark record %v", err)
	}

	err = store.Delete(b.Id)

	if err != nil {
		t.Fatalf("could not delete %v", err)
	}

	bookmarks, err := store.GetAll()

	if err != nil {
		t.Fatalf("could not get all bookmark records %v", err)
	}

	if len(bookmarks) != 0 {
		t.Errorf("unexpected bookmarks found %d", len(bookmarks))
	}
}
