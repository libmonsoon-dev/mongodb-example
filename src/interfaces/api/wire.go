//go:generate wire gen
//+build wireinject

package api

import (
	"mongodb-example/src/infrastructure"
	"mongodb-example/src/infrastructure/db/mongo"
	"mongodb-example/src/interfaces/mongo/repository"
	"mongodb-example/src/usecases"
	"net/http/httptest"

	"github.com/google/wire"
)

func InitTestServer() *httptest.Server {
	panic(wire.Build(
		httptest.NewServer,
		New,
		infrastructure.NewStdLogger,
		usecases.NewUserService,
		usecases.NewValidator,
		repository.NewMongoUserRepository,
		mongo.MustNewDbConnection,
		mongo.TestConfig,
	))
}
