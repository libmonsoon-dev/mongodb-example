package usecases

import (
	"context"
	"time"

	"mongodb-example/src/domain"
	"mongodb-example/src/domain/dto"
	"mongodb-example/src/usecases/adapter"
)

const userServiceStoreTimeout = 1 * time.Second

type UserService struct {
	userRepository UserRepository
	validator      Validator
}

func NewUserService(userRepository UserRepository, validator Validator) domain.UserService {
	return &UserService{userRepository, validator}
}

func (s *UserService) Store(dto dto.User) (string, error) {
	err := s.validator.Struct(dto)
	if err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(context.Background(), userServiceStoreTimeout)
	defer cancel()

	return s.userRepository.Store(ctx, adapter.DtoToUser(dto))
}
