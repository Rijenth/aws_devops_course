package domain

import "time"

type User struct {
	ID             int
	Username       string
	Password       string
	Email          string
	FirstName      string
	LastName       string
	DateOfBirth    time.Time
	Address        string
	PhoneNumber    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	LastLogin      *time.Time
	IsActive       bool
	IsAdmin        bool
	ProfilePicture string
	Bio            string
}
