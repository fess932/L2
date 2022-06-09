package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrMethodNotAllowed = errors.New("method not allowed")
)

type JSON map[string]interface{}

func JSONError(w http.ResponseWriter, code int, err error) {
	b, _ := json.Marshal(JSON{
		"error": err.Error(),
	})

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(b)
}
