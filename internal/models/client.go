package models

import (
	"time"
)

type ClientAddress struct {
	City   string
	Street string
	House  string
	Flat   string
	Desc   string
}

type ClientPet struct {
	ID          string
	Name        string
	Sex         string
	Type        string
	Breed       string
	Weight      int
	DateOfBirth time.Time
	CreatedAt   time.Time
	Desc        string
}

type Client struct {
	ID          string
	FirstName   string
	LastName    string
	Phone       string
	Email       string
	DateOfBirth time.Time
	CreatedAt   time.Time
	Pets        []ClientPet
	Address     ClientAddress
}
