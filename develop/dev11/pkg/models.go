package pkg

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Event struct {
	ID    string    `json:"id"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}

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

func JSONResponse(w http.ResponseWriter, code int, data interface{}) {
	b, _ := json.Marshal(JSON{
		"result": data,
	})

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(b)
}
