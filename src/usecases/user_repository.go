package usecases

import (
	"context"

	"mongodb-example/src/domain"
)

type UserRepository interface {
	Store(ctx context.Context, user domain.User) error
}
