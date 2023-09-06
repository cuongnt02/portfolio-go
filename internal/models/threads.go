package models

import (
    "database/sql"
    "time"
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
    return nil, nil
}

func (m *ThreadModel) Recent() ([]*Thread, error) {
    return nil, nil
}
