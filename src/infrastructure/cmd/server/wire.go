//+build wireinject
//go:generate wire

package main

import (
	"mongodb-example/src/infrastructure"
	"mongodb-example/src/infrastructure/db/mongo"
	"mongodb-example/src/interfaces/api"
	"mongodb-example/src/interfaces/mongo/repository"
	"mongodb-example/src/usecases"
	"net/http"

	"github.com/google/wire"
)

func Init() (http.Handler, error) {
	panic(wire.Build(
		api.New,
		infrastructure.NewStdLogger,
		usecases.NewUserService,
		usecases.NewValidator,
		repository.NewMongoUserRepository,
		mongo.NewDbConnection,
		mongo.DefaultConfig,
	))

}
