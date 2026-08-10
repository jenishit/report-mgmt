package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain/valueobjects"
)

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type LoginResponse struct {
	AccessToken string    `json:"access_token"`
	SessionID   uuid.UUID `json:"session_id"`
	UserID      uuid.UUID `json:"user_id"`
	UserRole    string    `json:"user_type"`
}

type PasswordReset struct {
	ID       uuid.UUID
	Password valueobjects.Password
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email" binding:"required,email"`
	OTP         string `json:"otp" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type ForgotPasswordResponse struct {
	OTP       string    `json:"otp"`
	ExpiresAt time.Time `json:"expires_at"`
	ResetURL  string    `json:"reset_url"`
}
