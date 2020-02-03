package adapter

import (
	"time"

	"mongodb-example/src/domain"
	"mongodb-example/src/domain/dto"
)

func DtoToGame(dto dto.Game) domain.Game {
	return domain.Game{
		int(dto.PointsGained),
		domain.WinStatus(dto.WinStatus),
		domain.GameType(dto.GameType),
		time.Time(dto.Created),
	}
}
