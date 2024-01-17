package ierr

import "fmt"

type Error struct {
	Code    string
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Error 500
var (
	ErrInternal   = Error{Code: "500-000", Message: "something error on system"}
	ErrInternalDB = Error{Code: "500-001", Message: "something error on our database"}
)

// Error 404
var (
	ErrNotFound = Error{Code: "404-000", Message: "resource not found"}
)

// Error 400
var (
	ErrorBadRequest = Error{Code: "400-000", Message: "bad request"}
)

// Error 403
var (
	ErrForbidden = Error{Code: "403-000", Message: "forbidden action"}
)

// Error 401
var (
	ErrUnauthorized = Error{Code: "401-000", Message: "unauthorized user"}
)
