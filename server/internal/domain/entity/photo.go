package entity

import (
	"github.com/google/uuid"
)

type Photo struct {
	ID     uuid.UUID
	URL    string
	IsMain bool

	UserID uuid.UUID
}
