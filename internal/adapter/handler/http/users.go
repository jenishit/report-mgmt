package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/handler/http/dto"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type UserHandler struct {
	usvc port.UserService
}

func NewUsersHandler(usvc port.UserService) *UserHandler {
	return &UserHandler{
		usvc: usvc,
	}
}

// CreateUser creates a new user
// @Summary Create user
// @Description Create a new user with role and profile
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.CreateUser true "User details"
// @Success 200 {object} response
// @Failure 400 {object} errorResponse
// @Failure 409 {object} errorResponse
// @Router /user/create [post]
func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var req dto.CreateUser

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}
	user, err := uh.usvc.CreateUser(ctx, &req)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, user)
}
