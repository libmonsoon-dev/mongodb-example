package apiV1

import (
	"mongodb-example/src/domain"
	"mongodb-example/src/usecases"
	"net/http"
)

func New(logger usecases.Logger, userService domain.UserService) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/user", NewUserHandler(logger, userService))

	return mux
}
