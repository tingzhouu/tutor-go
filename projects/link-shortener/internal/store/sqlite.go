package store

import (
	"database/sql"
	"tutor-go/projects/link-shortener/internal/model"

	_ "modernc.org/sqlite"
)

type Store struct {
	db *sql.DB
}

func New(dbPath string) (*Store, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS links (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			url TEXT NOT NULL,
			shortCode TEXT NOT NULL
		)
	`)
	if err != nil {
		return nil, err
	}

	return &Store{
		db,
	}, nil
}

func (s *Store) Create(l model.Link) (model.Link, error) {
	result, err := s.db.Exec(`
		INSERT INTO links (url, shortCode) VALUES (
			?, ?
		)
	`, l.Url, l.ShortCode)

	if err != nil {
		return model.Link{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return model.Link{}, err
	}

	l.Id = int(id)
	return l, nil
}

func (s *Store) GetOne(shortCode string) (model.Link, error) {
	l := model.Link{}
	row := s.db.QueryRow("SELECT * FROM links WHERE shortCode = ?", shortCode)
	err := row.Scan(&l.Id, &l.Url, &l.ShortCode)
	if err != nil {
		return l, err
	}
	return l, nil
}

func (s *Store) GetAll() ([]model.Link, error) {
	links := []model.Link{}
	rows, err := s.db.Query("SELECT * FROM links")
	if err != nil {
		return []model.Link{}, err
	}

	defer rows.Close()
	for rows.Next() {
		var l model.Link
		err := rows.Scan(&l.Id, &l.Url, &l.ShortCode)
		if err != nil {
			return links, err
		}
		links = append(links, l)
	}
	return links, nil
}

func (s *Store) Delete(id int) error {
	_, err := s.db.Exec(`DELETE FROM links WHERE id = ?`, id)
	if err != nil {
		return err
	}
	return nil
}
