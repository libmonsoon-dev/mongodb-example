package usecases

import (
	"context"
	"time"

	"mongodb-example/src/usecases/adapter"
	"mongodb-example/src/usecases/dto"
)

const userServiceStoreTimeout = 1 * time.Second

type UserService struct {
	userRepository UserRepository
	validator      Validator
}

func (s *UserService) Store(dto dto.UserDto) (string, error) {
	err := s.validator.Struct(dto)
	if err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(context.Background(), userServiceStoreTimeout)
	defer cancel()

	return s.userRepository.Store(ctx, adapter.DtoToUser(dto))
}
