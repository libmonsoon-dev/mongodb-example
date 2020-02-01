package adapter

import (
	"time"

	"mongodb-example/src/domain"
	"mongodb-example/src/usecases/dto"
)

func DtoToGame(dto dto.GameDto) domain.Game {
	return domain.Game{
		int(dto.PointsGained),
		domain.WinStatus(dto.WinStatusDto),
		domain.GameType(dto.GameTypeDto),
		time.Time(dto.Created),
	}
}