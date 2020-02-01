package dto

type GameDto struct {
	PointsGained stringInt `json:"points_gained" validate:"required"`
	WinStatusDto `json:"win_status" validate:"required,alphanum"`
	GameTypeDto  `json:"game_type" validate:"required,alphanum"`
	Created      DateDto `json:"created" validate:"required"`
}
