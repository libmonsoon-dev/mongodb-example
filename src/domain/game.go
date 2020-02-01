package domain

import "time"

type Game struct {
	PointsGained int
	WinStatus
	GameType
	Created time.Time
}
