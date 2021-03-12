/*
Copyright 2021 Adevinta
*/

package errors

import (
	"net/http"
)

// Default receives an object (error or string) and creates an ErrorStack.
// If object is already an ErrorStack then it appends the current error to the
// existing stack.
func Default(e interface{}) *ErrorStack {
	// Create an Error.
	newError := Error{
		Kind:           ErrInternal,
		Message:        interfaceToStr(e),
		HTTPStatusCode: http.StatusInternalServerError,
	}

	return newErrorStack(e, newError)
}

func Database(e interface{}) *ErrorStack {
	newError := Error{
		Kind:           ErrDatabase,
		Message:        interfaceToStr(e),
		HTTPStatusCode: http.StatusInternalServerError,
	}
	return newErrorStack(e, newError)
}

func Forbidden(e interface{}) *ErrorStack {
	newError := Error{
		Kind:           ErrForbidden,
		Message:        interfaceToStr(e),
		HTTPStatusCode: http.StatusForbidden,
	}
	return newErrorStack(e, newError)
}

func Unauthorized(e interface{}) *ErrorStack {
	newError := Error{
		Kind:           ErrUnauthorized,
		Message:        interfaceToStr(e),
		HTTPStatusCode: http.StatusUnauthorized,
	}
	return newErrorStack(e, newError)
}

func NotFound(e interface{}) *ErrorStack {
	newError := Error{
		Kind:           ErrNotFound,
		Message:        interfaceToStr(e),
		HTTPStatusCode: http.StatusNotFound,
	}
	return newErrorStack(e, newError)
}

func Create(e interface{}, extras ...string) *ErrorStack {
	newError := Error{
		Kind:           ErrCreate,
		Message:        interfaceToStr(e, extras...),
		HTTPStatusCode: http.StatusInternalServerError,
	}
	return newErrorStack(e, newError)
}

func Update(e interface{}) *ErrorStack {
	newError := Error{
		Kind:           ErrUpdate,
		Message:        interfaceToStr(e),
		HTTPStatusCode: http.StatusInternalServerError,
	}
	return newErrorStack(e, newError)
}

func Delete(e interface{}) *ErrorStack {
	newError := Error{
		Kind:           ErrDelete,
		Message:        interfaceToStr(e),
		HTTPStatusCode: http.StatusInternalServerError,
	}
	return newErrorStack(e, newError)
}

func Validation(e interface{}, extras ...string) *ErrorStack {
	newError := Error{
		Kind:           ErrValidation,
		Message:        interfaceToStr(e, extras...),
		HTTPStatusCode: http.StatusUnprocessableEntity,
	}
	return newErrorStack(e, newError)
}

func Duplicated(e interface{}) *ErrorStack {
	newError := Error{
		Kind:           ErrDuplicated,
		Message:        interfaceToStr(e),
		HTTPStatusCode: http.StatusConflict,
	}
	return newErrorStack(e, newError)
}

func Assertion(e interface{}) *ErrorStack {
	newError := Error{
		Kind:           ErrAssertion,
		Message:        interfaceToStr(e),
		HTTPStatusCode: http.StatusBadRequest,
	}
	return newErrorStack(e, newError)
}

func MethodNotAllowed(e interface{}) *ErrorStack {
	newError := Error{
		Kind:           ErrAssertion,
		Message:        interfaceToStr(e),
		HTTPStatusCode: http.StatusMethodNotAllowed,
	}
	return newErrorStack(e, newError)
}
