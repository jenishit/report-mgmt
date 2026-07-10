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

func (sh *LabHandler) GetAllLabs(ctx *gin.Context) {
	
	res, err := sh.svc.GetAllLabs(ctx)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := dto.LabsResponses(res)

	handleSuccess(ctx, rsp)
}
