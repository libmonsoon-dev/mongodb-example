package domain

import "time"

type User struct {
	Email     string
	LastName  string
	Country   string
	City      string
	Gender    string
	BirthDate time.Time
}
