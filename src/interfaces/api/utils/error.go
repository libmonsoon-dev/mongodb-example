package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiError struct {
	StatusCode  int
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

func (apiErr ApiError) WriteJson(w http.ResponseWriter) error {
	w.Header().Set(ContentType, ApplicationJson)
	w.WriteHeader(apiErr.StatusCode)
	raw, err := json.Marshal(apiErr)

	if err != nil {
		return err
	}
	_, err = w.Write(raw)

	return err
}

func NewInternalServerError() *ApiError {
	return &ApiError{
		http.StatusInternalServerError,
		"Internal Server Error",
		nil,
	}
}

func NewUnsupportedMediaType() *ApiError {
	return &ApiError{
		http.StatusUnsupportedMediaType,
		"Unsupported Media Type",
		nil,
	}
}

func NewMethodNotAllowed(method, path string) *ApiError {
	description := fmt.Sprintf("Cannot %v %v", method, path)
	return &ApiError{
		http.StatusMethodNotAllowed,
		"Method Not Allowed",
		&description,
	}
}
