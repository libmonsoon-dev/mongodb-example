package usecases

import (
	"mongodb-example/src/usecases/adapter"
	"mongodb-example/src/usecases/dto"
)

type UserService struct {
	userRepository UserRepository
	validator      Validator
}

func (s *UserService) Store(dto dto.UserDto) error {
	err := s.validator.Struct(dto)
	if err != nil {
		return err
	}
	return s.userRepository.Store(adapter.DtoToUser(dto))
}
