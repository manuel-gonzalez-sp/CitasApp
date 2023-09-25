package utils

import (
	"encoding/json"
	"net/http"
)

// TODO: HTTP standarized response, HTTP filters/sorting

func GetRequestBody[T any](r *http.Request) (T, error) {
	var value T
	err := json.NewDecoder(r.Body).Decode(value)
	return value, err
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(statusCode)

	if payload != nil {
		data, _ := json.Marshal(payload)
		w.Write(data)
	}
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}
