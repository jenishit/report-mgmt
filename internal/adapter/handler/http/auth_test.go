package http

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

func TestAuthHandler_Logout_RevokesCallerSession(t *testing.T) {
	sessionID := uuid.New()
	var revoked uuid.UUID
	svc := &mockAuthService{
		logoutFn: func(ctx context.Context, sid uuid.UUID) error {
			revoked = sid
			return nil
		},
	}
	h := NewAuthHandler(svc)

	caller := &domain.TokenPayload{UserId: uuid.New(), RoleName: "ROLE_TECHNICIAN", SessionID: sessionID}
	ctx, w := newJSONTestContext(t, http.MethodPost, "/api/auth/logout", nil, caller)

	h.Logout(ctx)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
	if revoked != sessionID {
		t.Fatalf("expected session %v to be revoked, got %v", sessionID, revoked)
	}
}

func TestAuthHandler_Logout_RequiresAuthentication(t *testing.T) {
	svc := &mockAuthService{
		logoutFn: func(ctx context.Context, sid uuid.UUID) error {
			t.Fatal("Logout should not be reached with no authenticated caller")
			return nil
		},
	}
	h := NewAuthHandler(svc)

	ctx, w := newJSONTestContext(t, http.MethodPost, "/api/auth/logout", nil, nil)

	h.Logout(ctx)

	if w.Code == http.StatusOK {
		t.Fatalf("expected an error status, got 200: %s", w.Body.String())
	}
}
