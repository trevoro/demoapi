package errors

import (
	"errors"
	"net/http"
)

type ApplicationError struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	Status  int    `json:"status"`
}

func (a ApplicationError) Error() string {
	return a.Message
}

var ErrPanic = ApplicationError{
	Message: "something went wrong",
	Name:    "internal_server_error",
	Status:  http.StatusInternalServerError,
}

var ErrBadRequest = ApplicationError{
	Message: "bad request",
	Name:    "bad_request",
	Status:  http.StatusBadRequest,
}

var ErrNotFound = ApplicationError{
	Message: "not found",
	Name:    "not_found",
	Status:  http.StatusNotFound,
}

var ErrNoAssociation = ApplicationError{
	Message: "no repository association",
	Name:    "no_association",
	Status:  http.StatusNotFound,
}

var ErrUnauthorized = ApplicationError{
	Message: "unauthorized",
	Name:    "unauthorized",
	Status:  http.StatusUnauthorized,
}

// ErrExists is a general error we can use when we violate specific integrity
// violations, hit duplicate keys, etc.
var ErrExists = ApplicationError{
	Message: "a duplicate key exists",
	Name:    "exists",
	Status:  http.StatusConflict,
}

// ErrIntegrity deals specifically with the issue of not referencing something
// that must be referenced. It's not for distinct or uniqueness, it's for more
// foreign key violations in the database / data model that need to propagate
// to the client with some additional information
var ErrIntegrity = ApplicationError{
	Message: "your request refers to an item that already exists",
	Name:    "integrity_constraint_violation",
	Status:  http.StatusBadRequest,
}

// pass-through to underlying errors package. not sure if we need this.
var As = errors.As
var Is = errors.Is
