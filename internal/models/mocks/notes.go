package mocks

import (
	"time"

	"notetaker.ntc02.net/internal/models"
)

var mockNote = &models.Note{
    ID: 1,
    Title: "HANDMADE 1",
    Content: "NOTHING HERE YET",
    CreatedDate: time.Now(),
}

type NoteModel struct {}

func (m *NoteModel) Insert(title string, content string) (int, error) {
    return 2, nil
}

func (m *NoteModel) Get(id int) (*models.Note, error){
    switch id {
    case 1:
        return mockNote, nil
    default:
        return nil, models.ErrNoRecord
    }
}

func (m *NoteModel) Update(title string, content string, id int)  (int, error) {
    return 1, nil
}

func (m *NoteModel) GetAll() ([]*models.Note, error) {
    return nil, nil
}
