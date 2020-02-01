package usecases

import (
	"mongodb-example/src/usecases/adapter"
	"mongodb-example/src/usecases/dto"
)

type GameService struct {
	gameRepository GameRepository
	validator      Validator
}

func (s *GameService) Store(dto dto.GameDto) error {
	err := s.validator.Struct(dto)
	if err != nil {
		return err
	}
	return s.gameRepository.Store(adapter.DtoToGame(dto))
}
