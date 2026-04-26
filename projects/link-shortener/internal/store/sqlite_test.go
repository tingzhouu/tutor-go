package store

import (
	"fmt"
	"testing"
	"tutor-go/projects/link-shortener/internal/model"
)

func TestCreate(t *testing.T) {
	store, err := New(t.TempDir() + "/test.db")

	if err != nil {
		t.Fatalf("could not create store %v", err)
	}

	l := model.Link{
		Url:       "mockurl",
		ShortCode: "mockshortcode",
	}

	l, err = store.Create(l)
	if err != nil {
		t.Fatalf("could not create link record %v", err)
	}

	if l.Id <= 0 {
		t.Errorf("unexpected id %d", l.Id)
	}
}

func TestGetOne(t *testing.T) {
	store, err := New(t.TempDir() + "/test.db")

	if err != nil {
		t.Fatalf("could not create store %v", err)
	}

	l := model.Link{
		Url:       "mockurl",
		ShortCode: "mockshortcode",
	}

	l, err = store.Create(l)

	if err != nil {
		t.Fatalf("could not create link record %v", err)
	}
	l, err = store.GetOne("mockshortcode")

	if err != nil {
		t.Fatalf("could not get link record %v", err)
	}

	if l.Url != "mockurl" {
		t.Fatalf("unexpected URL %s", l.Url)
	}
	if l.ShortCode != "mockshortcode" {
		t.Fatalf("unexpected shortcode %s", l.ShortCode)
	}
}

func TestGetAll(t *testing.T) {
	store, err := New(t.TempDir() + "/test.db")

	if err != nil {
		t.Fatalf("could not create store %v", err)
	}

	for i := range 3 {
		l := model.Link{
			Url:       fmt.Sprintf("mockurl%d", i),
			ShortCode: fmt.Sprintf("mockshortcode%d", i),
		}
		l, err = store.Create(l)
		if err != nil {
			t.Fatalf("could not create link %v", err)
		}
	}

	links, err := store.GetAll()
	if err != nil {
		t.Fatalf("could not get links")
	}

	if len(links) != 3 {
		t.Fatalf("unexpected number of links %d", len(links))
	}
}

func TestDeleteOne(t *testing.T) {
	store, err := New(t.TempDir() + "/test.db")

	if err != nil {
		t.Fatalf("could not create store %v", err)
	}

	for i := range 3 {
		l := model.Link{
			Url:       fmt.Sprintf("mockurl%d", i),
			ShortCode: fmt.Sprintf("mockshortcode%d", i),
		}
		l, err = store.Create(l)
	}

	err = store.Delete(2)

	if err != nil {
		t.Fatalf("could not delete link")
	}

	links, err := store.GetAll()
	if len(links) != 2 {
		t.Fatalf("unexpected number of links %d", len(links))
	}
}
