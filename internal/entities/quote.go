package entities

import (
	"time"

	"github.com/google/uuid"
)

type Quote struct {
	ID     uuid.UUID
	Quote  string
	Author string
	CreatedAt time.Time
	DeletedAt time.Time
}