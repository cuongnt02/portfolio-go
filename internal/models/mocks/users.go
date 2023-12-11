package mocks

import "notetaker.ntc02.net/internal/models"

type UserModel struct{}

func (m *UserModel) Insert(username, email, password string) error {
    switch email {
        case "oops@duped.com":
            return models.ErrDuplicateEmail
        default:
            return nil
    }
}


func (m *UserModel) Authenticate(email, password string) (int, error) {
    if email == "befruit@gmail.com" && password == "P@ssw0rd" {
        return 1, nil
    }

    return 0, models.ErrInvalidCredentials
}


func (m *UserModel) Exists(id int) (bool, error) {
    switch id {
    case 1:
        return true, nil
    default:
        return false, nil
    }
}
