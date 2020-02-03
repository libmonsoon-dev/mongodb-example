package adapter

import (
	"time"

	"mongodb-example/src/domain"
	"mongodb-example/src/domain/dto"
)

func DtoToUser(dto dto.User) domain.User {
	return domain.User{
		dto.Email,
		dto.LastName,
		dto.Country,
		dto.City,
		dto.Gender,
		time.Time(dto.BirthDate),
	}
}
