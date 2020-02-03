package usecases

import (
	"context"
	"mongodb-example/src/domain"
	"time"

	"mongodb-example/src/domain/dto"
	"mongodb-example/src/usecases/adapter"
)

const gameServiceStoreTimeout = 1 * time.Second

type GameService struct {
	gameRepository GameRepository
	validator      Validator
}

func NewGameService(gameRepository GameRepository, validator Validator) domain.GameService {
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
