package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/handler/http/dto"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

func init() {
	gin.SetMode(gin.TestMode)
}

// newJSONTestContext builds a gin context with the given JSON body and,
// optionally, an authenticated caller's token payload already set (as
// authMiddleware would have done).
func newJSONTestContext(t *testing.T, method, path string, body any, caller *domain.TokenPayload) (*gin.Context, *httptest.ResponseRecorder) {
	t.Helper()

	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			t.Fatalf("failed to encode request body: %v", err)
		}
	}

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest(method, path, &buf)
	ctx.Request.Header.Set("Content-Type", "application/json")

	if caller != nil {
		ctx.Set(authorizationPayloadKey, caller)
	}

	return ctx, w
}

type mockRoleService struct {
	createRoleFn      func(ctx context.Context, role *domain.Role) (*domain.Role, error)
	getRoleIDByNameFn func(ctx context.Context, name string) (*uuid.UUID, error)
	getRoleNameByIDFn func(ctx context.Context, id uuid.UUID) (*string, error)
}

func (m *mockRoleService) CreateRole(ctx context.Context, role *domain.Role) (*domain.Role, error) {
	return m.createRoleFn(ctx, role)
}

func (m *mockRoleService) GetRoleIDByRoleName(ctx context.Context, name string) (*uuid.UUID, error) {
	return m.getRoleIDByNameFn(ctx, name)
}

func (m *mockRoleService) GetRoleNameByRoleID(ctx context.Context, id uuid.UUID) (*string, error) {
	return m.getRoleNameByIDFn(ctx, id)
}

type mockUserService struct {
	createUserFn     func(ctx context.Context, data *dto.CreateUser) (*domain.User, error)
	getUserByEmailFn func(ctx context.Context, login *domain.Login) (*domain.BasicDetails, error)
}

func (m *mockUserService) CreateUser(ctx context.Context, data *dto.CreateUser) (*domain.User, error) {
	return m.createUserFn(ctx, data)
}

func (m *mockUserService) GetUserByEmail(ctx context.Context, login *domain.Login) (*domain.BasicDetails, error) {
	return m.getUserByEmailFn(ctx, login)
}

type mockProfileService struct {
	createProfileFn         func(ctx context.Context, profile *domain.Profile) (*domain.Profile, error)
	getProfileByIDFn        func(ctx context.Context, id uuid.UUID) (*domain.GetProfileDetails, error)
	getProfilesFn           func(ctx context.Context) ([]*domain.GetProfileDetails, error)
	updateProfileByUserIDFn func(ctx context.Context, prof *domain.GetProfileDetails) error
}

func (m *mockProfileService) CreateProfile(ctx context.Context, profile *domain.Profile) (*domain.Profile, error) {
	return m.createProfileFn(ctx, profile)
}

func (m *mockProfileService) GetProfileByID(ctx context.Context, id uuid.UUID) (*domain.GetProfileDetails, error) {
	return m.getProfileByIDFn(ctx, id)
}

func (m *mockProfileService) GetProfiles(ctx context.Context) ([]*domain.GetProfileDetails, error) {
	return m.getProfilesFn(ctx)
}

func (m *mockProfileService) UpdateProfileByUserID(ctx context.Context, prof *domain.GetProfileDetails) error {
	return m.updateProfileByUserIDFn(ctx, prof)
}

type mockAuthService struct {
	loginFn          func(ctx context.Context, details *domain.Login) (*domain.LoginResponse, error)
	forgotPasswordFn func(ctx context.Context, req *domain.ForgotPasswordRequest) (*domain.ForgotPasswordResponse, error)
	resetPasswordFn  func(ctx context.Context, req *domain.ResetPasswordRequest) error
	logoutFn         func(ctx context.Context, sessionID uuid.UUID) error
}

func (m *mockAuthService) Login(ctx context.Context, details *domain.Login) (*domain.LoginResponse, error) {
	return m.loginFn(ctx, details)
}

func (m *mockAuthService) ForgotPassword(ctx context.Context, req *domain.ForgotPasswordRequest) (*domain.ForgotPasswordResponse, error) {
	return m.forgotPasswordFn(ctx, req)
}

func (m *mockAuthService) ResetPassword(ctx context.Context, req *domain.ResetPasswordRequest) error {
	return m.resetPasswordFn(ctx, req)
}

func (m *mockAuthService) Logout(ctx context.Context, sessionID uuid.UUID) error {
	return m.logoutFn(ctx, sessionID)
}
