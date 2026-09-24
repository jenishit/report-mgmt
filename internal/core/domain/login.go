package domain

import (
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

// ForgotPasswordResponse deliberately carries no secret (OTP/reset URL): the
// endpoint always returns the same shape whether or not the email exists, and
// the OTP is delivered out-of-band (currently logged server-side), never in
// the API response, to avoid handing out a password-reset code to anyone who
// knows an email address.
type ForgotPasswordResponse struct {
	Message string `json:"message"`
}
