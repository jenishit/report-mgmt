package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type AuthHandler struct {
	authService port.AuthService
}

func NewAuthHandler(authService port.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Login authenticates a user and returns a JWT token
// @Summary Login
// @Description Authenticate a user with email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body domain.Login true "Login credentials"
// @Success 200 {object} response{data=domain.LoginResponse}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(ctx *gin.Context) {
	var req domain.Login
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	res, err := h.authService.Login(ctx, &req)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, res)
}

// ForgotPassword generates an OTP for password reset
// @Summary Forgot Password
// @Description Request a password reset OTP for an email
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body domain.ForgotPasswordRequest true "Email to reset password for"
// @Success 200 {object} response{data=domain.ForgotPasswordResponse}
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /auth/forgot-password [post]
func (h *AuthHandler) ForgotPassword(ctx *gin.Context) {
	var req domain.ForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	res, err := h.authService.ForgotPassword(ctx, &req)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, res)
}

// ResetPassword resets the password using a valid OTP
// @Summary Reset Password
// @Description Reset the password using email and the OTP
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body domain.ResetPasswordRequest true "Email, OTP and new password"
// @Success 200 {object} response
// @Failure 400 {object} errorResponse
// @Router /auth/reset-password [post]
func (h *AuthHandler) ResetPassword(ctx *gin.Context) {
	var req domain.ResetPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	err := h.authService.ResetPassword(ctx, &req)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, nil)
}
