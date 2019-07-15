package errors

import (
	"errors"
)

var ErrDatabase = errors.New("Database")
var ErrInternal = errors.New("Internal")
var ErrForbidden = errors.New("Forbidden")
var ErrUnauthorized = errors.New("Unauthorized")
var ErrNotFound = errors.New("Record not found")
var ErrDuplicated = errors.New("Duplicated record")
var ErrCreate = errors.New("Cannot create record")
var ErrUpdate = errors.New("Cannot update record")
var ErrDelete = errors.New("Cannot delete record")
var ErrValidation = errors.New("Validation")
var ErrAssertion = errors.New("Assertion")
