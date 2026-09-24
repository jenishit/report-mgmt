package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain/valueobjects"
)

type mockTokenRepository struct {
	createOTPFn func(ctx context.Context, userID uuid.UUID, otp string, expiresAt time.Time) error
	getOTPFn    func(ctx context.Context, email, otp string) (*domain.Token, error)
	markUsedFn  func(ctx context.Context, id uuid.UUID) error
}

func (m *mockTokenRepository) CreateOTP(ctx context.Context, userID uuid.UUID, otp string, expiresAt time.Time) error {
	return m.createOTPFn(ctx, userID, otp, expiresAt)
}

func (m *mockTokenRepository) GetOTP(ctx context.Context, email, otp string) (*domain.Token, error) {
	return m.getOTPFn(ctx, email, otp)
}

func (m *mockTokenRepository) MarkUsed(ctx context.Context, id uuid.UUID) error {
	return m.markUsedFn(ctx, id)
}

type mockTokenService struct {
	createAccessTokenFn func(user *domain.BasicDetails, sessionID uuid.UUID) (string, error)
	verifyAccessTokenFn func(token string) (*domain.TokenPayload, error)
}

func (m *mockTokenService) CreateAccessToken(user *domain.BasicDetails, sessionID uuid.UUID) (string, error) {
	return m.createAccessTokenFn(user, sessionID)
}

func (m *mockTokenService) VerifyAccessToken(token string) (*domain.TokenPayload, error) {
	return m.verifyAccessTokenFn(token)
}

func mustPassword(t *testing.T, plaintext string) valueobjects.Password {
	t.Helper()
	pwd, err := valueobjects.NewPassword(plaintext)
	if err != nil {
		t.Fatalf("failed to create password: %v", err)
	}
	return *pwd
}

func TestAuthService_Login_Success(t *testing.T) {
	userID := uuid.New()
	userRepo := &mockUserRepository{
		getUserByEmailFn: func(ctx context.Context, login *domain.Login) (*domain.BasicDetails, error) {
			return &domain.BasicDetails{
				ID:       userID,
				UserRole: "ROLE_ADMIN",
				Password: mustPassword(t, "correct-password"),
			}, nil
		},
	}
	tokenSvc := &mockTokenService{
		createAccessTokenFn: func(user *domain.BasicDetails, sessionID uuid.UUID) (string, error) {
			return "signed-token", nil
		},
	}

	svc := NewAuthService(userRepo, &mockTokenRepository{}, tokenSvc, "", 0)

	res, err := svc.Login(context.Background(), &domain.Login{Email: "user@example.com", Password: "correct-password"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.AccessToken != "signed-token" {
		t.Fatalf("expected access token %q, got %q", "signed-token", res.AccessToken)
	}
	if res.UserID != userID {
		t.Fatalf("expected user id %v, got %v", userID, res.UserID)
	}
	if res.UserRole != "ROLE_ADMIN" {
		t.Fatalf("expected role ROLE_ADMIN, got %v", res.UserRole)
	}
}

func TestAuthService_Login_WrongPassword(t *testing.T) {
	userRepo := &mockUserRepository{
		getUserByEmailFn: func(ctx context.Context, login *domain.Login) (*domain.BasicDetails, error) {
			return &domain.BasicDetails{
				ID:       uuid.New(),
				Password: mustPassword(t, "correct-password"),
			}, nil
		},
	}
	tokenSvc := &mockTokenService{
		createAccessTokenFn: func(user *domain.BasicDetails, sessionID uuid.UUID) (string, error) {
			t.Fatal("CreateAccessToken should not be called when credentials are invalid")
			return "", nil
		},
	}

	svc := NewAuthService(userRepo, &mockTokenRepository{}, tokenSvc, "", 0)

	_, err := svc.Login(context.Background(), &domain.Login{Email: "user@example.com", Password: "wrong-password"})
	if !errors.Is(err, domain.ErrInvalidCredentials) {
		t.Fatalf("expected ErrInvalidCredentials, got %v", err)
	}
}

func TestAuthService_Login_UserNotFound(t *testing.T) {
	userRepo := &mockUserRepository{
		getUserByEmailFn: func(ctx context.Context, login *domain.Login) (*domain.BasicDetails, error) {
			return nil, domain.ErrDataNotFound
		},
	}

	svc := NewAuthService(userRepo, &mockTokenRepository{}, &mockTokenService{}, "", 0)

	_, err := svc.Login(context.Background(), &domain.Login{Email: "missing@example.com", Password: "whatever1"})
	if !errors.Is(err, domain.ErrDataNotFound) {
		t.Fatalf("expected ErrDataNotFound, got %v", err)
	}
}

func TestAuthService_ForgotPassword_Success(t *testing.T) {
	userID := uuid.New()
	userRepo := &mockUserRepository{
		getUserByEmailFn: func(ctx context.Context, login *domain.Login) (*domain.BasicDetails, error) {
			return &domain.BasicDetails{ID: userID}, nil
		},
	}
	var storedOTP string
	tokenRepo := &mockTokenRepository{
		createOTPFn: func(ctx context.Context, uid uuid.UUID, otp string, expiresAt time.Time) error {
			if uid != userID {
				t.Fatalf("expected user id %v, got %v", userID, uid)
			}
			storedOTP = otp
			return nil
		},
	}

	svc := NewAuthService(userRepo, tokenRepo, &mockTokenService{}, "http://app/reset?token={token}", 15*time.Minute)

	res, err := svc.ForgotPassword(context.Background(), &domain.ForgotPasswordRequest{Email: "user@example.com"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(res.OTP) != otpLength {
		t.Fatalf("expected OTP of length %d, got %q", otpLength, res.OTP)
	}
	if res.OTP != storedOTP {
		t.Fatalf("expected returned OTP to match stored OTP")
	}
	wantURL := "http://app/reset?token=" + storedOTP
	if res.ResetURL != wantURL {
		t.Fatalf("expected reset url %q, got %q", wantURL, res.ResetURL)
	}
}

func TestAuthService_ForgotPassword_UnknownEmail(t *testing.T) {
	userRepo := &mockUserRepository{
		getUserByEmailFn: func(ctx context.Context, login *domain.Login) (*domain.BasicDetails, error) {
			return nil, errors.New("no rows")
		},
	}

	svc := NewAuthService(userRepo, &mockTokenRepository{}, &mockTokenService{}, "", 0)

	_, err := svc.ForgotPassword(context.Background(), &domain.ForgotPasswordRequest{Email: "missing@example.com"})
	if !errors.Is(err, domain.ErrDataNotFound) {
		t.Fatalf("expected ErrDataNotFound, got %v", err)
	}
}

func TestAuthService_ResetPassword_Success(t *testing.T) {
	tokenID := uuid.New()
	userID := uuid.New()

	tokenRepo := &mockTokenRepository{
		getOTPFn: func(ctx context.Context, email, otp string) (*domain.Token, error) {
			return &domain.Token{ID: tokenID, UserID: userID}, nil
		},
		markUsedFn: func(ctx context.Context, id uuid.UUID) error {
			if id != tokenID {
				t.Fatalf("expected token id %v, got %v", tokenID, id)
			}
			return nil
		},
	}
	var capturedReset *domain.PasswordReset
	userRepo := &mockUserRepository{
		updatePasswordFn: func(ctx context.Context, reset *domain.PasswordReset) error {
			capturedReset = reset
			return nil
		},
	}

	svc := NewAuthService(userRepo, tokenRepo, &mockTokenService{}, "", 0)

	err := svc.ResetPassword(context.Background(), &domain.ResetPasswordRequest{
		Email:       "user@example.com",
		OTP:         "123456",
		NewPassword: "brand-new-password",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if capturedReset == nil || capturedReset.ID != userID {
		t.Fatalf("expected password reset for user %v, got %+v", userID, capturedReset)
	}
}

func TestAuthService_ResetPassword_InvalidOTP(t *testing.T) {
	tokenRepo := &mockTokenRepository{
		getOTPFn: func(ctx context.Context, email, otp string) (*domain.Token, error) {
			return nil, errors.New("not found")
		},
	}

	svc := NewAuthService(&mockUserRepository{}, tokenRepo, &mockTokenService{}, "", 0)

	err := svc.ResetPassword(context.Background(), &domain.ResetPasswordRequest{
		Email:       "user@example.com",
		OTP:         "000000",
		NewPassword: "brand-new-password",
	})
	if !errors.Is(err, domain.ErrInvalidOTP) {
		t.Fatalf("expected ErrInvalidOTP, got %v", err)
	}
}
