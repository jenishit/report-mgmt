package services

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/handler/http/dto"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type mockUserRepository struct {
	createUserFn     func(ctx context.Context, user *domain.User) (*domain.User, error)
	getUserByEmailFn func(ctx context.Context, login *domain.Login) (*domain.BasicDetails, error)
	updatePasswordFn func(ctx context.Context, reset *domain.PasswordReset) error
}

func (m *mockUserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return m.createUserFn(ctx, user)
}

func (m *mockUserRepository) GetUserByEmail(ctx context.Context, login *domain.Login) (*domain.BasicDetails, error) {
	return m.getUserByEmailFn(ctx, login)
}

func (m *mockUserRepository) UpdatePassword(ctx context.Context, reset *domain.PasswordReset) error {
	return m.updatePasswordFn(ctx, reset)
}

type mockRoleService struct {
	getRoleIDByNameFn func(ctx context.Context, name string) (*uuid.UUID, error)
}

func (m *mockRoleService) CreateRole(ctx context.Context, role *domain.Role) (*domain.Role, error) {
	panic("not used in these tests")
}

func (m *mockRoleService) GetRoleIDByRoleName(ctx context.Context, name string) (*uuid.UUID, error) {
	return m.getRoleIDByNameFn(ctx, name)
}

func (m *mockRoleService) GetRoleNameByRoleID(ctx context.Context, id uuid.UUID) (*string, error) {
	panic("not used in these tests")
}

type mockProfileService struct {
	createProfileFn func(ctx context.Context, profile *domain.Profile) (*domain.Profile, error)
}

func (m *mockProfileService) CreateProfile(ctx context.Context, profile *domain.Profile) (*domain.Profile, error) {
	return m.createProfileFn(ctx, profile)
}

func (m *mockProfileService) GetProfileByID(ctx context.Context, id uuid.UUID) (*domain.GetProfileDetails, error) {
	panic("not used in these tests")
}

func (m *mockProfileService) GetProfiles(ctx context.Context) ([]*domain.GetProfileDetails, error) {
	panic("not used in these tests")
}

func (m *mockProfileService) UpdateProfileByUserID(ctx context.Context, prof *domain.GetProfileDetails) error {
	panic("not used in these tests")
}

func newCreateUserRequest() *dto.CreateUser {
	return &dto.CreateUser{
		FirstName: "Ada",
		LastName:  "Lovelace",
		Email:     "ada@example.com",
		Password:  "longenoughpassword",
		Phone:     "1234567890",
		RoleName:  "ROLE_ADMIN",
	}
}

func TestUserService_CreateUser_Success(t *testing.T) {
	roleID := uuid.New()
	userID := uuid.New()

	roleSvc := &mockRoleService{
		getRoleIDByNameFn: func(ctx context.Context, name string) (*uuid.UUID, error) {
			return &roleID, nil
		},
	}
	userRepo := &mockUserRepository{
		createUserFn: func(ctx context.Context, user *domain.User) (*domain.User, error) {
			if user.RoleID != roleID {
				t.Fatalf("expected role id %v, got %v", roleID, user.RoleID)
			}
			user.ID = userID
			return user, nil
		},
	}
	var capturedProfile *domain.Profile
	profileSvc := &mockProfileService{
		createProfileFn: func(ctx context.Context, profile *domain.Profile) (*domain.Profile, error) {
			capturedProfile = profile
			return profile, nil
		},
	}

	svc := NewUserService(userRepo, roleSvc, profileSvc)

	user, err := svc.CreateUser(context.Background(), newCreateUserRequest())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.ID != userID {
		t.Fatalf("expected user id %v, got %v", userID, user.ID)
	}
	if capturedProfile == nil {
		t.Fatal("expected profile to be created")
	}
	if capturedProfile.UserID != userID {
		t.Fatalf("expected profile user id %v, got %v", userID, capturedProfile.UserID)
	}
}

func TestUserService_CreateUser_UnknownRole(t *testing.T) {
	roleSvc := &mockRoleService{
		getRoleIDByNameFn: func(ctx context.Context, name string) (*uuid.UUID, error) {
			return nil, domain.ErrDataNotFound
		},
	}
	// repo/profile must never be called when the role lookup fails.
	userRepo := &mockUserRepository{
		createUserFn: func(ctx context.Context, user *domain.User) (*domain.User, error) {
			t.Fatal("CreateUser should not be called when role lookup fails")
			return nil, nil
		},
	}
	profileSvc := &mockProfileService{
		createProfileFn: func(ctx context.Context, profile *domain.Profile) (*domain.Profile, error) {
			t.Fatal("CreateProfile should not be called when role lookup fails")
			return nil, nil
		},
	}

	svc := NewUserService(userRepo, roleSvc, profileSvc)

	_, err := svc.CreateUser(context.Background(), newCreateUserRequest())
	if !errors.Is(err, domain.ErrDataNotFound) {
		t.Fatalf("expected ErrDataNotFound, got %v", err)
	}
}

func TestUserService_CreateUser_InvalidPassword(t *testing.T) {
	roleID := uuid.New()
	roleSvc := &mockRoleService{
		getRoleIDByNameFn: func(ctx context.Context, name string) (*uuid.UUID, error) {
			return &roleID, nil
		},
	}
	userRepo := &mockUserRepository{
		createUserFn: func(ctx context.Context, user *domain.User) (*domain.User, error) {
			t.Fatal("CreateUser should not be called when password is invalid")
			return nil, nil
		},
	}
	profileSvc := &mockProfileService{}

	svc := NewUserService(userRepo, roleSvc, profileSvc)

	req := newCreateUserRequest()
	req.Password = "short"

	_, err := svc.CreateUser(context.Background(), req)
	if err == nil {
		t.Fatal("expected an error for a too-short password")
	}
}

func TestUserService_CreateUser_RepoErrorPropagates(t *testing.T) {
	roleID := uuid.New()
	repoErr := errors.New("duplicate email")

	roleSvc := &mockRoleService{
		getRoleIDByNameFn: func(ctx context.Context, name string) (*uuid.UUID, error) {
			return &roleID, nil
		},
	}
	userRepo := &mockUserRepository{
		createUserFn: func(ctx context.Context, user *domain.User) (*domain.User, error) {
			return nil, repoErr
		},
	}
	profileSvc := &mockProfileService{
		createProfileFn: func(ctx context.Context, profile *domain.Profile) (*domain.Profile, error) {
			t.Fatal("CreateProfile should not be called when user creation fails")
			return nil, nil
		},
	}

	svc := NewUserService(userRepo, roleSvc, profileSvc)

	_, err := svc.CreateUser(context.Background(), newCreateUserRequest())
	if !errors.Is(err, repoErr) {
		t.Fatalf("expected %v, got %v", repoErr, err)
	}
}

func TestUserService_CreateUser_ProfileErrorPropagates(t *testing.T) {
	roleID := uuid.New()
	profileErr := errors.New("profile insert failed")

	roleSvc := &mockRoleService{
		getRoleIDByNameFn: func(ctx context.Context, name string) (*uuid.UUID, error) {
			return &roleID, nil
		},
	}
	userRepo := &mockUserRepository{
		createUserFn: func(ctx context.Context, user *domain.User) (*domain.User, error) {
			user.ID = uuid.New()
			return user, nil
		},
	}
	profileSvc := &mockProfileService{
		createProfileFn: func(ctx context.Context, profile *domain.Profile) (*domain.Profile, error) {
			return nil, profileErr
		},
	}

	svc := NewUserService(userRepo, roleSvc, profileSvc)

	_, err := svc.CreateUser(context.Background(), newCreateUserRequest())
	if !errors.Is(err, profileErr) {
		t.Fatalf("expected %v, got %v", profileErr, err)
	}
}
