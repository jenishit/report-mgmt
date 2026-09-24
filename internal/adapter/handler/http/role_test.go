package http

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/handler/http/dto"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

func TestRoleHandler_CreateRole_NonAdminRoleNeedsNoElevatedCaller(t *testing.T) {
	svc := &mockRoleService{
		createRoleFn: func(ctx context.Context, role *domain.Role) (*domain.Role, error) {
			return &domain.Role{ID: uuid.New(), RoleName: role.RoleName}, nil
		},
	}
	h := NewRoleHandler(svc)

	// No caller payload at all - a plain admin creating an ordinary role
	// doesn't need to be checked against RoleSuperAdmin.
	ctx, w := newJSONTestContext(t, http.MethodPost, "/api/admin/role/create", dto.CreateRole{RoleName: "ROLE_TECHNICIAN"}, nil)

	h.CreateRole(ctx)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
}

func TestRoleHandler_CreateRole_AdminRoleRequiresSuperAdmin(t *testing.T) {
	svc := &mockRoleService{
		createRoleFn: func(ctx context.Context, role *domain.Role) (*domain.Role, error) {
			t.Fatal("CreateRole should not be reached when caller is not super admin")
			return nil, nil
		},
	}
	h := NewRoleHandler(svc)

	caller := &domain.TokenPayload{UserId: uuid.New(), RoleName: RoleAdmin}
	ctx, w := newJSONTestContext(t, http.MethodPost, "/api/admin/role/create", dto.CreateRole{RoleName: RoleAdmin}, caller)

	h.CreateRole(ctx)

	if w.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d: %s", w.Code, w.Body.String())
	}
}

func TestRoleHandler_CreateRole_SuperAdminCanCreateAdminRole(t *testing.T) {
	var created *domain.Role
	svc := &mockRoleService{
		createRoleFn: func(ctx context.Context, role *domain.Role) (*domain.Role, error) {
			created = role
			return &domain.Role{ID: uuid.New(), RoleName: role.RoleName}, nil
		},
	}
	h := NewRoleHandler(svc)

	caller := &domain.TokenPayload{UserId: uuid.New(), RoleName: RoleSuperAdmin}
	ctx, w := newJSONTestContext(t, http.MethodPost, "/api/admin/role/create", dto.CreateRole{RoleName: RoleAdmin}, caller)

	h.CreateRole(ctx)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
	if created == nil || created.RoleName != RoleAdmin {
		t.Fatalf("expected ROLE_ADMIN to be created, got %+v", created)
	}
}

func TestRoleHandler_CreateRole_MissingPayloadForAdminRole(t *testing.T) {
	svc := &mockRoleService{
		createRoleFn: func(ctx context.Context, role *domain.Role) (*domain.Role, error) {
			t.Fatal("CreateRole should not be reached with no caller payload")
			return nil, nil
		},
	}
	h := NewRoleHandler(svc)

	// Requesting an admin-level role with no authenticated caller at all
	// must be rejected, not silently allowed.
	ctx, w := newJSONTestContext(t, http.MethodPost, "/api/admin/role/create", dto.CreateRole{RoleName: RoleSuperAdmin}, nil)

	h.CreateRole(ctx)

	if w.Code == http.StatusOK {
		t.Fatalf("expected an error status, got 200: %s", w.Body.String())
	}
}
