package usecases

import (
	"context"
	"time"

	"mongodb-example/src/usecases/adapter"
	"mongodb-example/src/usecases/dto"
)

const gameServiceStoreTimeout = 1 * time.Second

type IGameService interface {
	Store(dto dto.Game) (string, error)
}

type GameService struct {
	gameRepository GameRepository
	validator      Validator
}

func NewGameService(gameRepository GameRepository, validator Validator) IGameService {
	return &GameService{gameRepository: gameRepository, validator: validator}
}

func (s *GameService) Store(dto dto.Game) (string, error) {
	err := s.validator.Struct(dto)
	if err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(context.Background(), gameServiceStoreTimeout)
	defer cancel()

	return s.gameRepository.Store(ctx, adapter.DtoToGame(dto))
}
