package models

import "errors"

var ErrNoRecord = errors.New("models: no matching record found")
var ErrEntityNotChanged = errors.New("models: entity not changed")
