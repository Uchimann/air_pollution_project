package repository

import "errors"

var ErrNotFound = errors.New("data not found")
var ErrDatabaseError = errors.New("database error")