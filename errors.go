/*
Copyright 2021 Adevinta
*/

package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error represents an application error. It does contain:
// - an array of previous errors,
// - a textual representation of the current error,
// - an error type that can be compared with standard errors,
// - an http status code.
type Error struct {
	ID             int
	Message        string
	Kind           error
	HTTPStatusCode int
}

type ErrorStack struct {
	Errors []Error
}

// StatusCode returns an HTTP status code
// Satisfies GoKit StatusCoder interface.
func (e *Error) StatusCode() int {
	return e.HTTPStatusCode
}

// StatusCode returns an HTTP status code
// Satisfies GoKit StatusCoder interface.
func (e *ErrorStack) StatusCode() int {
	if len(e.Errors) == 0 {
		// https://github.com/go-kit/kit/blob/master/transport/http/server.go#L205
		return http.StatusInternalServerError
	}
	lastIndex := len(e.Errors) - 1

	return e.Errors[lastIndex].HTTPStatusCode
}

// Error returns a textual representation of the error.
// Satisfies the standard error interface.
func (e Error) Error() string {
	return e.Message
}

// Error returns a textual representation of the error.
// Satisfies the standard error interface.
func (e ErrorStack) Error() string {
	if len(e.Errors) > 0 {
		lastIndex := len(e.Errors) - 1
		return e.Errors[lastIndex].Error()
	}

	return ""
}

// MarshalJSON returns an array of bytes which contains a marshaled json version
// of the error.
// Satisfies the Marshaler interface.
func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID   int    `json:"id"`
		Code int    `json:"code"`
		Err  string `json:"error"`
		Type string `json:"type"`
	}{
		ID:   e.ID,
		Code: e.HTTPStatusCode,
		Err:  e.Message,
		Type: e.Kind.Error(),
	})
}

type errorStackPayload struct {
	Code  int     `json:"code"`
	Err   string  `json:"error"`
	Type  string  `json:"type"`
	Stack []Error `json:"parent_errors,omitempty"`
}

// MarshalJSON returns an array of bytes which contains a marshaled json version
// of the error.
// Satisfies the Marshaler interface.
func (e *ErrorStack) MarshalJSON() ([]byte, error) {
	lastIndex := len(e.Errors) - 1
	if lastIndex >= 0 {
		lastError := e.Errors[lastIndex]
		return json.Marshal(errorStackPayload{
			Code:  lastError.HTTPStatusCode,
			Err:   lastError.Message,
			Type:  lastError.Kind.Error(),
			Stack: e.Errors[:lastIndex],
		})
	}
	return []byte{}, nil
}

// UnmarshalJSON allows to unmarshal an ErrorStack from his json representation.
func (e *ErrorStack) UnmarshalJSON(data []byte) error {
	// By now this will work we should think seriously about refactoring this
	// package to make easier to marshal and unmarshal from json.
	payload := new(errorStackPayload)
	if err := json.Unmarshal(data, payload); err != nil {
		return Default(string(data))
	}
	var err *ErrorStack
	switch payload.Code {
	case http.StatusForbidden:
		err = Forbidden("")
	case http.StatusUnauthorized:
		err = Unauthorized("")
	case http.StatusNotFound:
		err = NotFound("")
	case http.StatusUnprocessableEntity:
		err = Validation("")
	case http.StatusConflict:
		err = Duplicated("")
	case http.StatusBadRequest:
		err = Assertion("")
	case http.StatusMethodNotAllowed:
		err = MethodNotAllowed("")
	default:
		err = Default("")
	}
	e.Errors = err.Errors
	e.Errors[0].Message = payload.Err
	return nil
}

// IsKind checks if an error is an error stack, and if it is, checks if the last
// error of the stack is of a given kind. Notice that the kind is also an error.
// The kind comparision is done not by the string representation of the two
// error kinds but direct instance comparision.
func IsKind(err error, kind error) bool {
	if e, ok := err.(*ErrorStack); ok {
		if len(e.Errors) > 0 {
			return e.Errors[len(e.Errors)-1].Kind == kind
		}
	}
	return false
}

// IsRootOfKind checks if an error is an error stack, and if it is, checks if
// the root error of the stack is of a given kind. Notice that the kind is also
// an error. The kind comparision is done not by the string representation of
// the two error kinds but direct instance comparision.
func IsRootOfKind(err error, kind error) bool {
	if e, ok := err.(*ErrorStack); ok {
		if len(e.Errors) > 0 {
			return e.Errors[0].Kind == kind
		}
	}
	return false
}

// interfaceToStr receives an object which the type is unknown and tries do
// detect it's textual representation, parting from the assumption that it is
// either an object that satifies the standard error interface or an string.
// If it's not an error or a string, then an empty string is returned.
func interfaceToStr(err interface{}, resources ...string) string {
	errMsg := ""
	if e, ok := err.(error); ok {
		errMsg = e.Error()
	}
	if e, ok := err.(string); ok {
		errMsg = e
	}
	if len(resources) > 0 {
		preffix := ""
		for _, r := range resources {
			preffix = fmt.Sprintf("%s[%s]", preffix, r)
		}
		errMsg = fmt.Sprintf("%s %s", preffix, errMsg)
	}
	return errMsg
}

func newErrorStack(obj interface{}, newError Error) *ErrorStack {
	var result = &ErrorStack{}

	// Check if obj type is already an `ErrorStack`.
	errorStack, ok := obj.(*ErrorStack)
	if ok {
		newError.ID = errorStack.Errors[len(errorStack.Errors)-1].ID + 1
		errorStack.Errors = append(errorStack.Errors, newError)
		result = errorStack
	} else {
		newError.ID = 0
		result.Errors = append(result.Errors, newError)
	}

	return result
}
