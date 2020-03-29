package utils

import (
	"net/http"
)

// Error structure for errors expected by clients (REST API).
type Error struct {
	Status           int
	ErrorDescription string
}

// Error objects
var (
	ErrArgsMissing = Error{Status: http.StatusBadRequest, ErrorDescription: "Argument value missing: "}
	ErrDecodeBody  = Error{Status: http.StatusBadRequest, ErrorDescription: "Error decoding body."}
	ErrInsert      = Error{Status: http.StatusBadRequest, ErrorDescription: "Error inserting new document"}
	ErrNotFound    = Error{Status: http.StatusNotFound, ErrorDescription: "Object not found."}
	ErrQuery       = Error{Status: http.StatusInternalServerError, ErrorDescription: "Decoding documents into results."}
	ErrDatabase    = Error{Status: http.StatusInternalServerError, ErrorDescription: "Error returnning a cursor over the matching documents in the collection."}
)

// ErrorResponse raises error
func ErrorResponse(w http.ResponseWriter, err Error, value string) {
	w.WriteHeader(err.Status)
	w.Write([]byte(err.ErrorDescription + value))
	Logger.Errorf("%s", err.ErrorDescription+value)
}
