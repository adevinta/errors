# Package Errors

#### Goals:

- Simple
- Predictable
- All methods will return non `nil` objects that satifies the standard error interface
- Errors have types that can be comparable
- No `reflect`


#### Usage

```
err := readToken(user)
if err != nil {
    return errors.Unauthorized(err)
}
```

`errors.Unauthorized` will return an error object configured as follows:

- `type: errors.ErrUnauthorized`
- `code: http.StatusUnauthorized`

output
```
{
  "code": 401,
  "error": "Cannot read token",
  "type": "Unauthorized"
}
```

#### Error types

```
var ErrDatabase = errors.New("Database")
var ErrInternal = errors.New("Internal")
var ErrForbiden = errors.New("Forbiden")
var ErrUnauthorized = errors.New("Unauthorized")
var ErrNotFound = errors.New("Record not found")
var ErrDuplicated = errors.New("Duplicated record")
var ErrCreate = errors.New("Cannot create record")
var ErrUpdate = errors.New("Cannot update record")
var ErrDelete = errors.New("Cannot delete record")
var ErrValidation = errors.New("Validation")
var ErrAssertion = errors.New("Assertion")
```

#### Comparing Errors

```
errors.IsKind(err, errors.ErrUnauthorized)
```

#### Internals

This package works with the concept of error stacks.

```
type ErrorStack struct {
	Errors []Error
}
```

An `ErrorStack` contains an array of errors. You can chain multiple errors, one after another in order to get a stack trace from the original error until the final error.

```
// Error represents an application error. It does contain:
// - a textual representation of the current error Also
// - an error type that can be compared with standard errors
// - an http status code
type Error struct {
	ID      int
	Message string
	Kind    error
	HTTPStatusCode int
}
```
