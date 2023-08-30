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
	sqlStatement := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(sqlStatement, id)

	s := &TextSnippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
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
	sqlStatement := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	textSnippets := []*TextSnippet{}

	for rows.Next() {
		t := &TextSnippet{}

		err = rows.Scan(&t.ID, &t.Title, &t.Content, &t.Created, &t.Expires)
		if err != nil {
			return nil, err
		}

		textSnippets = append(textSnippets, t)

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return textSnippets, nil
}
