package api

import (
	"mongodb-example/src/domain"
	"mongodb-example/src/usecases"
	"net/http"

	apiV1 "mongodb-example/src/interfaces/api/v1"
)

func New(logger usecases.Logger, userService domain.UserService) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/v1/", http.StripPrefix("/v1", apiV1.New(logger, userService)))
	return mux
}
