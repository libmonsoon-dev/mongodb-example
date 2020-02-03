package domain

import "mongodb-example/src/domain/dto"

type GameService interface {
	Store(dto dto.Game) (string, error)
}
