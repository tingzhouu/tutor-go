package store

import (
	"database/sql"
	"tutor-go/projects/bookmark-api/internal/model"

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
		CREATE TABLE IF NOT EXISTS bookmarks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			url TEXT NOT NULL,
			title TEXT NOT NULL,
			tags TEXT NOT NULL
		)
	`)
	if err != nil {
		return nil, err
	}

	return &Store{
		db,
	}, nil
}

func (s *Store) Create(b model.Bookmark) (model.Bookmark, error) {
	result, err := s.db.Exec(`
		INSERT INTO bookmarks (url, title, tags) VALUES (
			?, ?, ?
		)
	`, b.Url, b.Title, b.Tags)

	if err != nil {
		return model.Bookmark{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return model.Bookmark{}, err
	}

	b.Id = int(id)
	return b, nil
}

func (s *Store) GetAll() ([]model.Bookmark, error) {
	bookmarks := []model.Bookmark{}
	rows, err := s.db.Query(`SELECT * FROM bookmarks`)
	if err != nil {
		return []model.Bookmark{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var b model.Bookmark
		err := rows.Scan(&b.Id, &b.Url, &b.Title, &b.Tags)
		if err != nil {
			return bookmarks, err
		}
		bookmarks = append(bookmarks, b)
	}
	return bookmarks, nil
}

func (s *Store) GetOne(id int) (model.Bookmark, error) {
	b := model.Bookmark{}
	row := s.db.QueryRow(`SELECT * FROM bookmarks WHERE id = ?`, id)

	err := row.Scan(&b.Id, &b.Url, &b.Title, &b.Tags)
	if err != nil {
		return b, err
	}
	return b, nil
}

func (s *Store) Delete(id int) error {
	_, err := s.db.Exec(`DELETE FROM bookmarks WHERE id = ?`, id)
	if err != nil {
		return err
	}
	return nil
}
