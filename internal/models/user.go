package models

import "time"

type User struct {
	ID          string
	FirstName   string
	LastName    string
	Roles       []string
	DateOfBirth time.Time
	CreatedAt   time.Time
}
