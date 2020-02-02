package dto

type Game struct {
	PointsGained stringInt `json:"points_gained" validate:"required"`
	WinStatus    `json:"win_status" validate:"required,alphanum"`
	GameType     `json:"game_type" validate:"required,alphanum"`
	Created      Date `json:"created" validate:"required"`
}
