package dto

type Game struct {
	PointsGained stringInt `json:"points_gained" validate:"required"`
	WinStatus    stringInt `json:"win_status" validate:"required,alphanum"`
	GameType     stringInt `json:"game_type" validate:"required,alphanum"`
	Created      Date      `json:"created" validate:"required"`
}
