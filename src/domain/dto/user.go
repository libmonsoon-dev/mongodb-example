package dto

type User struct {
	Email     string    `json:"email"  validate:"required,email"`
	LastName  string    `json:"last_name" validate:"required"`
	Country   string    `json:"country" validate:"required"`
	City      string    `json:"city" validate:"required"`
	Gender    string    `json:"gender" validate:"required"`
	BirthDate BirthDate `json:"birth_date" validate:"required,alphanum"`
}
