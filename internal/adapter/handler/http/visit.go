package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/handler/http/dto"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type VisitHandler struct {
	svc port.VisitService
}

func NewVisitHandler(svc port.VisitService) *VisitHandler {
	return &VisitHandler{
		svc: svc,
	}
}

func (vh *VisitHandler) CreateVisit(ctx *gin.Context) {
	var req dto.CreateVisit

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

	v := &domain.Visit{
		VisitNo:      req.VisitNo,
		PatientID:    req.PatientID,
		DoctorID:     req.DoctorID,
		Status:       domain.Status(req.Status),
		RegisteredBy: userPayload.UserId,
	}

	vis, err := vh.svc.CreateVisit(ctx, v)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, vis)
}

func (vh *VisitHandler) GetVisitByID(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := uuid.Parse(id)

	if err != nil {
		parseError(err)
		return
	}

	res, err := vh.svc.GetVisitByID(ctx, uid)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := dto.VisitsResponse(res)

	handleSuccess(ctx, rsp)
}
