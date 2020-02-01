package adapter

import (
	"time"

	"mongodb-example/src/domain"
	"mongodb-example/src/usecases/dto"
)

func DtoToUser(dto dto.UserDto) domain.User {
	return domain.User{
		dto.Email,
		dto.LastName,
		dto.Country,
		dto.City,
		dto.Gender,
		time.Time(dto.BirthDate),
	}
}
