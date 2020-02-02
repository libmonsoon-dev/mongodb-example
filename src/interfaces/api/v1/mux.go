package apiV1

import "net/http"

func New() *http.ServeMux {
	mux := http.NewServeMux()
	// mux.Handle()
	return mux
}
