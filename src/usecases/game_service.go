package usecases

import (
	"context"
	"time"

	"mongodb-example/src/usecases/adapter"
	"mongodb-example/src/usecases/dto"
)

const gameServiceStoreTimeout = 1 * time.Second

type GameService struct {
	gameRepository GameRepository
	validator      Validator
}

func (s *GameService) Store(dto dto.GameDto) error {
	err := s.validator.Struct(dto)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), gameServiceStoreTimeout)
	defer cancel()

	if err := s.gameRepository.Store(ctx, adapter.DtoToGame(dto)); err != nil {
		return err
	}
	return ctx.Err()
}
