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
	City         string
	Country      string
	CreatedAt    time.Time
	PasswordHash string
	Photos       []Photo
	//Interests    []string

}
