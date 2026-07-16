package http

import (

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/handler/http/dto"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type LabHandler struct {
	svc port.LabService
}

func NewLabHandler(svc port.LabService) *LabHandler {
	return &LabHandler{
		svc: svc,
	}
}

// InsertLab creates a new lab settings entry
// @Summary Create lab settings
// @Description Create new lab settings (admin only)
// @Tags Lab Settings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.LabRequest true "Lab settings details"
// @Success 200 {object} response{data=domain.LabSettings}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /admin/lab [post]
func (sh *LabHandler) InsertLab(ctx *gin.Context) {
	var req dto.LabRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

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

	lab := &domain.LabSettings{
		LabName:        req.LabName,
		Tagline:        &req.Tagline,
		Address:        &req.Address,
		Phone:          &req.Phone,
		Email:          &req.Email,
		RegistrationNo: &req.RegistrationNo,
		ReportFooter:   &req.ReportFooter,
		UpdatedBy:      userPayload.UserId,
	}

	s, err := sh.svc.InsertLab(ctx, lab)

	if err != nil {
		handleError(ctx, err)
		return
	}
	handleSuccess(ctx, s)
}

// GetLabByID returns lab settings by ID
// @Summary Get lab settings
// @Description Get lab settings by ID
// @Tags Lab Settings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Lab ID"
// @Success 200 {object} response{data=domain.LabSettings}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /lab/get/{id} [get]
func (sh *LabHandler) GetLabByID(ctx *gin.Context) {
	labID := ctx.Param("id")
	labUUID, err := uuid.Parse(labID)
	if err != nil {
		handleError(ctx, domain.ErrInvalidUUID)
		return
	}

	lab, err := sh.svc.GetLabByLabID(ctx, labUUID)

	if err != nil {
		handleError(ctx, err)
		return
	}
	handleSuccess(ctx, lab)
}

// UpdateLab updates lab settings
// @Summary Update lab settings
// @Description Update lab settings by ID (admin only)
// @Tags Lab Settings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Lab ID"
// @Param request body dto.LabRequest true "Lab settings details"
// @Success 200 {object} response
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /admin/lab/{id} [patch]
func (sh *LabHandler) UpdateLab(ctx *gin.Context) {
	labID := ctx.Param("id")
	labUUID, err := uuid.Parse(labID)
	if err != nil {
		handleError(ctx, domain.ErrInvalidUUID)
		return
	}

	var req dto.LabRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

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

	lab := &domain.LabSettings{
		ID:             labUUID,
		LabName:        req.LabName,
		Tagline:        &req.Tagline,
		Address:        &req.Address,
		Phone:          &req.Phone,
		Email:          &req.Email,
		RegistrationNo: &req.RegistrationNo,
		ReportFooter:   &req.ReportFooter,
		UpdatedBy:      userPayload.UserId,
	}

	err = sh.svc.UpdateLab(ctx, lab)
	if err != nil {
		handleError(ctx, err)
		return
	}
	handleSuccess(ctx, gin.H{"message": "Lab settings updated successfully"})
}

// GetAllLabs returns all lab settings
// @Summary List all lab settings
// @Description Get all lab settings (admin only)
// @Tags Lab Settings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response{data=[]dto.LabResponse}
// @Failure 401 {object} errorResponse
// @Router /admin/lab [get]
func (sh *LabHandler) GetAllLabs(ctx *gin.Context) {
	
	res, err := sh.svc.GetAllLabs(ctx)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := dto.LabsResponses(res)

	handleSuccess(ctx, rsp)
}
