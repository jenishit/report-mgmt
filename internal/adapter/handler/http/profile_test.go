package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/handler/http/dto"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

func newUpdateProfileContext(t *testing.T, targetID uuid.UUID, caller *domain.TokenPayload) (*gin.Context, *httptest.ResponseRecorder) {
	t.Helper()
	ctx, w := newJSONTestContext(t, http.MethodPatch, "/api/profile/update-profile/"+targetID.String(), dto.UpdateProfileRequest{
		FirstName: "New",
		LastName:  "Name",
	}, caller)
	ctx.Params = gin.Params{{Key: "id", Value: targetID.String()}}
	return ctx, w
}

func TestProfileHandler_UpdateProfileByUserID_OwnerCanUpdateSelf(t *testing.T) {
	userID := uuid.New()
	var updated bool
	svc := &mockProfileService{
		updateProfileByUserIDFn: func(ctx context.Context, prof *domain.GetProfileDetails) error {
			updated = true
			if prof.UserID != userID {
				t.Fatalf("expected update for user %v, got %v", userID, prof.UserID)
			}
			return nil
		},
	}
	h := NewProfileHandler(svc)

	caller := &domain.TokenPayload{UserId: userID, RoleName: "ROLE_TECHNICIAN"}
	ctx, w := newUpdateProfileContext(t, userID, caller)

	h.UpdateProfileByUserID(ctx)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
	if !updated {
		t.Fatal("expected the profile service to be called")
	}
}

func TestProfileHandler_UpdateProfileByUserID_NonOwnerForbidden(t *testing.T) {
	targetID := uuid.New()
	svc := &mockProfileService{
		updateProfileByUserIDFn: func(ctx context.Context, prof *domain.GetProfileDetails) error {
			t.Fatal("UpdateProfileByUserID should not be reached for a non-owner, non-admin caller")
			return nil
		},
	}
	h := NewProfileHandler(svc)

	caller := &domain.TokenPayload{UserId: uuid.New(), RoleName: "ROLE_TECHNICIAN"}
	ctx, w := newUpdateProfileContext(t, targetID, caller)

	h.UpdateProfileByUserID(ctx)

	if w.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d: %s", w.Code, w.Body.String())
	}
}

func TestProfileHandler_UpdateProfileByUserID_AdminCanUpdateAnyone(t *testing.T) {
	targetID := uuid.New()
	var updated bool
	svc := &mockProfileService{
		updateProfileByUserIDFn: func(ctx context.Context, prof *domain.GetProfileDetails) error {
			updated = true
			return nil
		},
	}
	h := NewProfileHandler(svc)

	caller := &domain.TokenPayload{UserId: uuid.New(), RoleName: RoleAdmin}
	ctx, w := newUpdateProfileContext(t, targetID, caller)

	h.UpdateProfileByUserID(ctx)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
	if !updated {
		t.Fatal("expected the profile service to be called")
	}
}
