package port

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type TokenRepository interface {
	CreateOTP(ctx context.Context, userID uuid.UUID, otp string, expiresAt time.Time) error
	GetOTP(ctx context.Context, email, otp string) (*domain.Token, error)
	MarkUsed(ctx context.Context, id uuid.UUID) error
}
