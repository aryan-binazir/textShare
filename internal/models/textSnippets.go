package models

import (
	"database/sql"
	"errors"
	"time"
)

type TextSnippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type TextSnippetModel struct {
	DB *sql.DB
}

func (m *TextSnippetModel) Insert(title string, content string, expires int) (int, error) {
	sqlStatement := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(sqlStatement, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *TextSnippetModel) Get(id int) (*TextSnippet, error) {
	s := &TextSnippet{}

	err := m.DB.QueryRow("SELECT ...", id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *TextSnippetModel) Latest() ([]*TextSnippet, error) {
	return nil, nil
}
