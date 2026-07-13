package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/handler/http/dto"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type ReportHandler struct {
	svc port.ReportService
}

func NewReportHandler(svc port.ReportService) *ReportHandler {
	return &ReportHandler{svc: svc}
}

func (h *ReportHandler) CreateReport(ctx *gin.Context) {
	var req dto.CreateReportRequest

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

	report := &domain.Report{
		ReportNo:    req.ReportNo,
		VisitID:     req.VisitID,
		GeneratedBy: userPayload.UserId,
		Status:      domain.ReportDraft,
	}

	res, err := h.svc.CreateReport(ctx, report)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, dto.ReportResponseFromDomain(res))
}

func (h *ReportHandler) GetReportByID(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		handleError(ctx, domain.ErrInvalidUUID)
		return
	}

	res, err := h.svc.GetReportByID(ctx, uid)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, dto.ReportResponseFromDomain(res))
}

func (h *ReportHandler) GetReportsByVisitID(ctx *gin.Context) {
	id := ctx.Param("visit_id")

	uid, err := uuid.Parse(id)
	if err != nil {
		handleError(ctx, domain.ErrInvalidUUID)
		return
	}

	res, err := h.svc.GetReportsByVisitID(ctx, uid)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, dto.ReportsResponseFromDomain(res))
}

func (h *ReportHandler) UpdateReport(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		handleError(ctx, domain.ErrInvalidUUID)
		return
	}

	var req dto.UpdateReportRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	report := &domain.Report{
		ID:      uid,
		Status:  req.Status,
		PDFPath: req.PDFPath,
	}

	err = h.svc.UpdateReport(ctx, report)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, gin.H{"message": "Report updated successfully"})
}

func (h *ReportHandler) CreateReportPrint(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		handleError(ctx, domain.ErrInvalidUUID)
		return
	}

	var req dto.ReportPrintRequest

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

	print := &domain.ReportPrint{
		ReportID:   uid,
		PrintedBy:  userPayload.UserId,
		CopyNumber: req.CopyNumber,
	}

	res, err := h.svc.CreateReportPrint(ctx, print)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, dto.ReportPrintResponseFromDomain(res))
}

func (h *ReportHandler) GetReportPrints(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		handleError(ctx, domain.ErrInvalidUUID)
		return
	}

	res, err := h.svc.GetReportPrintsByReportID(ctx, uid)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, dto.ReportPrintsResponseFromDomain(res))
}
