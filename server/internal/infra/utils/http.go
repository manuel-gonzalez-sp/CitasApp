package utils

import (
	"encoding/json"
	"net/http"
	"time"

	"go.uber.org/multierr"
)

type Response[T any] struct {
	Data   []T      `json:"data,omitempty"`
	Errors []string `json:"errors,omitempty"`
	Meta   Meta     `json:"meta"`
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

func GetRequestBody[T interface{}](r *http.Request) (value T, err error) {
	decoder := json.NewDecoder(r.Body)
	decodeErr := decoder.Decode(&value)
	if decodeErr != nil {
		err = multierr.Append(err, decodeErr)
	}
	validationErr := ValidateStruct(value)
	if validationErr != nil {
		err = multierr.Append(err, validationErr)
	}
	return value, err
}

func WriteResponse[T interface{}](w http.ResponseWriter, statusCode int, data ...T) {
	response := &Response[T]{
		Data: data,
		Meta: newMeta(statusCode),
	}
	WriteJSON(w, statusCode, response)
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	var responseErrs []string
	errs := multierr.Errors(err)
	for _, err := range errs {
		responseErrs = append(responseErrs, err.Error())
	}
	response := &Response[any]{
		Errors: responseErrs,
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
