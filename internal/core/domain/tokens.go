package domain

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Token     uuid.UUID
	OTP       string
	ExpiresAt time.Time
	CreatedAt time.Time
	UsedAt    *time.Time
}
