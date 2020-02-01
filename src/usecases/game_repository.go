package usecases

import "mongodb-example/src/domain"

type GameRepository interface {
	Store(user domain.Game) error
}
