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
