package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/handler/http/dto"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type ProfileHandler struct {
	psvc port.ProfileService
}

func NewProfileHandler(psvc port.ProfileService) *ProfileHandler {
	return &ProfileHandler{
		psvc: psvc,
	}
}

// GetProfileByID returns the profile of the authenticated user
// @Summary Get my profile
// @Description Get the profile of the currently authenticated user
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response{data=dto.ProfileResponse}
// @Failure 401 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /profile/getme [get]
func (ph *ProfileHandler) GetProfileByID(ctx *gin.Context) {
	payload, exists := ctx.Get(authorizationPayloadKey)
	if !exists {
		validationError(ctx, domain.ErrEmptyAuthorizationHeader)
		return
	}

	userPayload, ok := payload.(*domain.TokenPayload)
	if !ok {
		validationError(ctx, domain.ErrInvalidAuthorizationHeader)
		return
	}

	res, err := ph.psvc.GetProfileByID(ctx, userPayload.UserId)
	if err != nil {
		handleError(ctx, err)
		return
	}
	rsp := dto.NewProfileResponse(res)

	handleSuccess(ctx, rsp)
}

// GetProfiles returns all profiles (admin only)
// @Summary List all profiles
// @Description Get all user profiles (admin only)
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response{data=[]dto.ProfileResponse}
// @Failure 401 {object} errorResponse
// @Failure 403 {object} errorResponse
// @Router /admin/profile/profile-details [get]
func (ph *ProfileHandler) GetProfiles(ctx *gin.Context) {
	res, err := ph.psvc.GetProfiles(ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}
	rsp := dto.NewProfileResponses(res)

	handleSuccess(ctx, rsp)
}

// UpdateProfileByUserID updates the profile of the authenticated user
// @Summary Update profile
// @Description Update the profile of the currently authenticated user
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Param request body dto.UpdateProfileRequest true "Profile details"
// @Success 200 {object} response
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /profile/update-profile/{id} [patch]
func (ph *ProfileHandler) UpdateProfileByUserID(ctx *gin.Context) {
	payload, exists := ctx.Get(authorizationPayloadKey)
	if !exists {
		validationError(ctx, domain.ErrEmptyAuthorizationHeader)
		return
	}

	userPayload, ok := payload.(*domain.TokenPayload)
	if !ok {
		validationError(ctx, domain.ErrInvalidAuthorizationHeader)
		return
	}

	var req dto.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	profile := &domain.GetProfileDetails{
		UserID:    userPayload.UserId,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
	}

	err := ph.psvc.UpdateProfileByUserID(ctx, profile)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, "profile updated successfully")
}
