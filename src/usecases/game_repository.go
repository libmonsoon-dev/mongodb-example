package usecases

import (
	"context"

	"mongodb-example/src/domain"
)

type GameRepository interface {
	Store(ctx context.Context, user domain.Game) (id string, err error)
}
