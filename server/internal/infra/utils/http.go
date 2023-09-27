package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

// TODO:  HTTP filters/sorting
type Response[T interface{}] struct {
	Data   []T      `json:"data,omitempty"`
	Errors []string `json:"errors,omitempty"`
	Meta   Meta     `json:"meta,omitempty"`
}
type Meta struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	Version   string `json:"version"`
}

func newMeta(statusCode int) Meta {
	return Meta{
		Status:    statusCode,
		Message:   http.StatusText(statusCode),
		Timestamp: time.Now().Unix(),
		Version:   "1.0.0",
	}
}

func GetRequestBody[T any](r *http.Request) (T, error) {
	var value T
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&value)
	return value, err
}

func WriteResponse[T interface{}](w http.ResponseWriter, statusCode int, data ...T) {
	response := &Response[T]{
		Data: data,
		Meta: newMeta(statusCode),
	}
	WriteJSON(w, statusCode, response)
}

func WriteError(w http.ResponseWriter, statusCode int, errors ...error) {
	var responseErrors []string
	for _, err := range errors {
		responseErrors = append(responseErrors, err.Error())
	}
	response := &Response[any]{
		Errors: responseErrors,
		Meta:   newMeta(statusCode),
	}
	WriteJSON(w, statusCode, response)
}

func WriteJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(statusCode)

	if payload != nil {
		data, _ := json.Marshal(payload)
		w.Write(data)
	}
}
