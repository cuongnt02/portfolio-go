package models

import "errors"

var (
    ErrNoRecord = errors.New("models: no matching record found")
    ErrEntityNotChanged = errors.New("models: entity not changed")
    ErrInvalidCredentials = errors.New("models: invalid credentials")
    ErrDuplicateEmail = errors.New("models: duplicate email")
)
