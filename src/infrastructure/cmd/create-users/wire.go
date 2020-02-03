//+build wireinject
//go:generate wire

package main

import (
	"mongodb-example/src/domain"
	"mongodb-example/src/infrastructure/db/mongo"
	"mongodb-example/src/interfaces/mongo/repository"
	"mongodb-example/src/usecases"

	"github.com/google/wire"
)

func Init() (domain.UserService, error) {
	panic(wire.Build(
		usecases.NewUserService,
		usecases.NewValidator,
		repository.NewMongoUserRepository,
		mongo.NewDbConnection,
		mongo.DefaultConfig,
	))
}
