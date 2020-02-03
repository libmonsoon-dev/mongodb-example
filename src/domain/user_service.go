package domain

import "mongodb-example/src/domain/dto"

type UserService interface {
	Store(dto dto.User) (string, error)
}
