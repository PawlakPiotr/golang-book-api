package utils

import (
	"encoding/json"
	"errors"
	"net/http"

	m "golang-book-api/model"

	"github.com/Sirupsen/logrus"
)

// Logger object with given log level
var Logger = NewLogger(logrus.DebugLevel)

// LoggerIn returns func entry log
func LoggerIn(function string) {
	Logger.Debugf("> [%s]", function)
}

// LoggerOut returns func exit log
func LoggerOut(function string) {
	Logger.Debugf("< [%s]", function)
}

// CreateBookParseRequest parses requeset of CreateBook function [API][CreateBook]
func CreateBookParseRequest(w http.ResponseWriter, r *http.Request) (*m.Book, error) {
	req := json.NewDecoder(r.Body)
	var input m.Book
	err := req.Decode(&input)
	if err != nil {
		ErrorResponse(w, ErrDecodeBody, "")
		return nil, errors.New(ErrDecodeBody.ErrorDescription)
	}
	return &input, nil
}
