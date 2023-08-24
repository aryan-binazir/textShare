package models

import (
	"database/sql"
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
	return 0, nil
}

func (m *TextSnippetModel) Get(id int) (*TextSnippet, error) {
	return nil, nil
}

func (m *TextSnippetModel) Latest() ([]*TextSnippet, error) {
	return nil, nil
}
