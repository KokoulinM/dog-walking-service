package models

import "time"

type Order struct {
	ID          string
	ClientID    string
	ClientPetID string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	StartDate   time.Time
	EndDate     time.Time
}
