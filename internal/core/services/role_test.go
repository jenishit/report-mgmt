package services

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type mockRoleRepository struct {
	createRoleFn      func(ctx context.Context, role *domain.Role) (*domain.Role, error)
	getRoleIDByNameFn func(ctx context.Context, name string) (*uuid.UUID, error)
	getRoleNameByIDFn func(ctx context.Context, id uuid.UUID) (*string, error)
}

func (m *mockRoleRepository) CreateRole(ctx context.Context, role *domain.Role) (*domain.Role, error) {
	return m.createRoleFn(ctx, role)
}

func (m *mockRoleRepository) GetRoleIDByRoleName(ctx context.Context, name string) (*uuid.UUID, error) {
	return m.getRoleIDByNameFn(ctx, name)
}

func (m *mockRoleRepository) GetRoleNameByRoleID(ctx context.Context, id uuid.UUID) (*string, error) {
	return m.getRoleNameByIDFn(ctx, id)
}

func TestRoleService_CreateRole(t *testing.T) {
	want := &domain.Role{ID: uuid.New(), RoleName: "ROLE_ADMIN"}
	repo := &mockRoleRepository{
		createRoleFn: func(ctx context.Context, role *domain.Role) (*domain.Role, error) {
			return want, nil
		},
	}
	svc := NewRoleService(repo)

	got, err := svc.CreateRole(context.Background(), &domain.Role{RoleName: "ROLE_ADMIN"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestRoleService_CreateRole_PropagatesRepoError(t *testing.T) {
	repoErr := errors.New("insert failed")
	repo := &mockRoleRepository{
		createRoleFn: func(ctx context.Context, role *domain.Role) (*domain.Role, error) {
			return nil, repoErr
		},
	}
	svc := NewRoleService(repo)

	_, err := svc.CreateRole(context.Background(), &domain.Role{RoleName: "ROLE_ADMIN"})
	if !errors.Is(err, repoErr) {
		t.Fatalf("expected %v, got %v", repoErr, err)
	}
}

func TestRoleService_GetRoleIDByRoleName(t *testing.T) {
	id := uuid.New()
	repo := &mockRoleRepository{
		getRoleIDByNameFn: func(ctx context.Context, name string) (*uuid.UUID, error) {
			if name != "ROLE_ADMIN" {
				t.Fatalf("expected role name ROLE_ADMIN, got %s", name)
			}
			return &id, nil
		},
	}
	svc := NewRoleService(repo)

	got, err := svc.GetRoleIDByRoleName(context.Background(), "ROLE_ADMIN")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if *got != id {
		t.Fatalf("expected %v, got %v", id, *got)
	}
}

func TestRoleService_GetRoleIDByRoleName_NotFound(t *testing.T) {
	repo := &mockRoleRepository{
		getRoleIDByNameFn: func(ctx context.Context, name string) (*uuid.UUID, error) {
			return nil, domain.ErrDataNotFound
		},
	}
	svc := NewRoleService(repo)

	_, err := svc.GetRoleIDByRoleName(context.Background(), "ROLE_UNKNOWN")
	if !errors.Is(err, domain.ErrDataNotFound) {
		t.Fatalf("expected ErrDataNotFound, got %v", err)
	}
}
