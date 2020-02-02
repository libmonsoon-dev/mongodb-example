package api

import (
	"net/http"

	apiV1 "mongodb-example/src/interfaces/api/v1"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/v1/", http.StripPrefix("/v1", apiV1.New()))
	return mux
}
