package http

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/handler/http/dto"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

func newCreateUserRequest(roleName string) dto.CreateUser {
	return dto.CreateUser{
		FirstName: "Ada",
		LastName:  "Lovelace",
		Email:     "ada@example.com",
		Password:  "longenoughpassword",
		Phone:     "1234567890",
		RoleName:  roleName,
	}
}

func TestUserHandler_CreateUser_NonAdminRoleNeedsNoElevatedCaller(t *testing.T) {
	svc := &mockUserService{
		createUserFn: func(ctx context.Context, data *dto.CreateUser) (*domain.User, error) {
			return &domain.User{ID: uuid.New()}, nil
		},
	}
	h := NewUsersHandler(svc)

	ctx, w := newJSONTestContext(t, http.MethodPost, "/api/admin/user/create", newCreateUserRequest("ROLE_TECHNICIAN"), nil)

	h.CreateUser(ctx)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
}

func TestUserHandler_CreateUser_AdminRoleRequiresSuperAdmin(t *testing.T) {
	svc := &mockUserService{
		createUserFn: func(ctx context.Context, data *dto.CreateUser) (*domain.User, error) {
			t.Fatal("CreateUser should not be reached when caller is not super admin")
			return nil, nil
		},
	}
	h := NewUsersHandler(svc)

	caller := &domain.TokenPayload{UserId: uuid.New(), RoleName: RoleAdmin}
	ctx, w := newJSONTestContext(t, http.MethodPost, "/api/admin/user/create", newCreateUserRequest(RoleAdmin), caller)

	h.CreateUser(ctx)

	if w.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d: %s", w.Code, w.Body.String())
	}
}

func TestUserHandler_CreateUser_SuperAdminCanCreateAdmin(t *testing.T) {
	var createdRole string
	svc := &mockUserService{
		createUserFn: func(ctx context.Context, data *dto.CreateUser) (*domain.User, error) {
			createdRole = data.RoleName
			return &domain.User{ID: uuid.New()}, nil
		},
	}
	h := NewUsersHandler(svc)

	caller := &domain.TokenPayload{UserId: uuid.New(), RoleName: RoleSuperAdmin}
	ctx, w := newJSONTestContext(t, http.MethodPost, "/api/admin/user/create", newCreateUserRequest(RoleAdmin), caller)

	h.CreateUser(ctx)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
	if createdRole != RoleAdmin {
		t.Fatalf("expected user to be created with role %q, got %q", RoleAdmin, createdRole)
	}
}

func TestUserHandler_CreateUser_MissingPayloadForAdminRole(t *testing.T) {
	svc := &mockUserService{
		createUserFn: func(ctx context.Context, data *dto.CreateUser) (*domain.User, error) {
			t.Fatal("CreateUser should not be reached with no caller payload")
			return nil, nil
		},
	}
	h := NewUsersHandler(svc)

	ctx, w := newJSONTestContext(t, http.MethodPost, "/api/admin/user/create", newCreateUserRequest(RoleSuperAdmin), nil)

	h.CreateUser(ctx)

	if w.Code == http.StatusOK {
		t.Fatalf("expected an error status, got 200: %s", w.Body.String())
	}
}
