package models

import (
	"database/sql"
	"errors"
	"time"
)

type Note struct {
    ID          int
    Title       string
    Content     string
    CreatedDate time.Time
    UpdatedDate time.Time
}

type NoteModel struct {
    DB *sql.DB
}

func (m *NoteModel) Insert(title string, content string) (int, error) {
    stmt := `INSERT INTO notes (title, content, created_date, updated_date) 
    VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id`

    lastInsertId := 0
    err := m.DB.QueryRow(stmt, title, content).Scan(&lastInsertId)
    if err != nil {
        return 0, err
    }

    return int(lastInsertId), nil
}

func (m *NoteModel) Get(id int) (*Note, error) {
    stmt := `SELECT id, title, content, created_date, updated_date 
            FROM notes 
            WHERE id=$1`

    row := m.DB.QueryRow(stmt, id)

    s := &Note{}

    err := row.Scan(&s.ID, &s.Title, &s.Content, &s.CreatedDate, &s.UpdatedDate)
    if err != nil {

        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrNoRecord
        } else {
            return nil, err
        }


    }

    return s, nil
}

func (m *NoteModel) GetAll() ([]*Note, error) {
    stmt := `SELECT id, title, content, created_date, updated_date 
            FROM notes 
            ORDER BY id DESC 
            LIMIT 10`

    rows, err := m.DB.Query(stmt)
    if err != nil {
        return nil, err
    }


    defer rows.Close()

    notes := []*Note{}

    for rows.Next() {

        n := &Note{}

        err = rows.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedDate, &n.UpdatedDate)
        if err != nil {
            return nil, err
        }

        notes = append(notes, n)

    }

    if err = rows.Err(); err != nil {
        return nil, err
    }
    
    return notes, nil
}
