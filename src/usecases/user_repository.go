package usecases

import "mongodb-example/src/domain"

type UserRepository interface {
	Store(user domain.User) error
}
