package apiV1

import (
	"mongodb-example/src/usecases"
	"net/http"
)

func New(logger usecases.Logger, userService usecases.IUserService) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/user/", http.StripPrefix("/user", NewUserHandler(logger, userService)))
	return mux
}
