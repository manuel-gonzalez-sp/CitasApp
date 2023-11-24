package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Username     string
	Fullname     string
	BirthDate    time.Time
	Gender       string
	Introduction string
	LookingFor   string
	//Interests    []string
	City    string
	Country string
	//Photos       []Photo
	CreatedAt    time.Time
	ActiveAt     time.Time
	PasswordHash []byte
}

func NewUser(username string, fullname string) *User {
	return &User{
		ID:       uuid.New(),
		Username: username,
		Fullname: fullname,
	}
}
