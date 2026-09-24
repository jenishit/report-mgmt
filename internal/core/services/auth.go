package services

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain/valueobjects"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

const otpLength = 6

type AuthService struct {
	repo     port.UserRepository
	tr       port.TokenRepository
	ts       port.TokenService
	resetURL string
	otpTTL   time.Duration
}

func NewAuthService(userRepo port.UserRepository, tokenRepo port.TokenRepository, tokenService port.TokenService, resetURL string, otpTTL time.Duration) *AuthService {
	return &AuthService{
		repo:     userRepo,
		tr:       tokenRepo,
		ts:       tokenService,
		resetURL: resetURL,
		otpTTL:   otpTTL,
	}
}

func (as *AuthService) Login(ctx context.Context, details *domain.Login) (*domain.LoginResponse, error) {
	sessionID := uuid.New()
	user, err := as.repo.GetUserByEmail(ctx, details)

	if err != nil {
		return nil, err
	}

	passwordVO, err := valueobjects.NewPasswordFromHash(user.Password.Hash())
	if err != nil {
		return nil, domain.ErrInvalidCredentials
	}
	if err := passwordVO.Verify(details.Password); err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	accessToken, err := as.ts.CreateAccessToken(user, sessionID)

	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		AccessToken: accessToken,
		SessionID:   sessionID,
		UserID:      user.ID,
		UserRole:    string(user.UserRole),
	}, nil
}

func (as *AuthService) ForgotPassword(ctx context.Context, req *domain.ForgotPasswordRequest) (*domain.ForgotPasswordResponse, error) {
	user, err := as.repo.GetUserByEmail(ctx, &domain.Login{Email: req.Email})
	if err != nil {
		return nil, domain.ErrDataNotFound
	}

	otp, err := generateOTP(otpLength)
	if err != nil {
		return nil, err
	}

	expiresAt := time.Now().UTC().Add(as.otpTTL)

	if err := as.tr.CreateOTP(ctx, user.ID, otp, expiresAt); err != nil {
		return nil, err
	}

	return &domain.ForgotPasswordResponse{
		OTP:       otp,
		ExpiresAt: expiresAt,
		ResetURL:  buildResetURL(as.resetURL, otp),
	}, nil
}

func (as *AuthService) ResetPassword(ctx context.Context, req *domain.ResetPasswordRequest) error {
	token, err := as.tr.GetOTP(ctx, req.Email, req.OTP)
	if err != nil {
		return domain.ErrInvalidOTP
	}

	newPassword, err := valueobjects.NewPassword(req.NewPassword)
	if err != nil {
		return err
	}

	err = as.repo.UpdatePassword(ctx, &domain.PasswordReset{
		ID:       token.UserID,
		Password: *newPassword,
	})
	if err != nil {
		return err
	}

	return as.tr.MarkUsed(ctx, token.ID)
}

func generateOTP(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("invalid otp length")
	}
	limit := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length)), nil)
	n, err := rand.Int(rand.Reader, limit)
	if err != nil {
		return "", fmt.Errorf("failed to generate otp: %w", err)
	}
	return fmt.Sprintf("%0*d", length, n.Int64()), nil
}

func buildResetURL(baseURL, otp string) string {
	if baseURL == "" {
		return ""
	}
	return strings.ReplaceAll(baseURL, "{token}", otp)
}
