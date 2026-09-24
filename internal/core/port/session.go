package port

import (
	"context"

	"github.com/google/uuid"
)

// SessionRepository tracks revoked JWT sessions (logout, forced sign-out)
// since access tokens are otherwise stateless and can't be invalidated on
// their own before they expire.
type SessionRepository interface {
	Revoke(ctx context.Context, sessionID uuid.UUID) error
	IsRevoked(ctx context.Context, sessionID uuid.UUID) (bool, error)
}
