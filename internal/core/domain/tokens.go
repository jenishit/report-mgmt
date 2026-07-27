package domain

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Token     uuid.UUID
	ExpiresAt time.Time
	CreatedAt time.Time
}
