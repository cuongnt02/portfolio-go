package models

import (
    "database/sql"
    "time"
    "errors"
)

type Thread struct {
    ID int
    Name string
    Created time.Time 
    Updated time.Time
}

type ThreadModel struct {
    DB *sql.DB
}

func (m *ThreadModel) Insert(name string) (int, error) {
    stmt := `INSERT INTO threads (name, created_time, updated_time) 
    VALUES(?, UTC_TIMESTAMP(), UTC_TIMESTAMP())`

    result, err := m.DB.Exec(stmt, name)
    if err != nil {
        return 0, err
    }

    id, err := result.LastInsertId()

    if err != nil {
        return 0, err
    }

    return int(id), nil
}

func (m *ThreadModel) Get(id int) (*Thread, error) {
    stmt := `SELECT id, name, created_time, updated_time FROM threads
    WHERE id=?`


    t := &Thread{}

    err := m.DB.QueryRow(stmt, id).Scan(&t.ID, &t.Name, &t.Created, &t.Updated)

    if err != nil {
        
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrNoRecord
        } else {
            return nil, err
        }
    }
    
    return t, nil
}

func (m *ThreadModel) GetAll() ([]*Thread, error) {
    stmt := `SELECT id, name, created_time, updated_time FROM threads`

    rows, err := m.DB.Query(stmt)
    if err != nil {
        return nil, err
    }

    defer rows.Close()


    threads := []*Thread{}

    for rows.Next() {
        t := &Thread{}

        err = rows.Scan(&t.ID, &t.Name, &t.Created, &t.Updated)

        if err != nil {
            return nil, err
        }

        threads = append(threads, t)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }


    return threads, nil
}

