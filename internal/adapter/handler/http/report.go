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

// CreateReport creates a new report
// @Summary Create report
// @Description Create a new report for a visit
// @Tags Reports
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateReportRequest true "Report details"
// @Success 200 {object} response{data=dto.ReportResponse}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /report [post]
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

// GetReportByID returns a report by ID
// @Summary Get report by ID
// @Description Get a report's details by ID
// @Tags Reports
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Report ID"
// @Success 200 {object} response{data=dto.ReportResponse}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /report/{id} [get]
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

// GetReportsByVisitID returns reports for a visit
// @Summary List reports by visit
// @Description Get all reports for a given visit
// @Tags Reports
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param visit_id path string true "Visit ID"
// @Success 200 {object} response{data=[]dto.ReportResponse}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /report/visit/{visit_id} [get]
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

// UpdateReport updates a report
// @Summary Update report
// @Description Update an existing report
// @Tags Reports
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Report ID"
// @Param request body dto.UpdateReportRequest true "Report details"
// @Success 200 {object} response
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /report/{id} [patch]
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

// CreateReportPrint creates a print record for a report
// @Summary Create report print
// @Description Create a print record for a report
// @Tags Reports
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Report ID"
// @Param request body dto.ReportPrintRequest true "Print details"
// @Success 200 {object} response{data=dto.ReportPrintResponse}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /report/{id}/print [post]
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

// GetReportPrints returns print records for a report
// @Summary List report prints
// @Description Get all print records for a given report
// @Tags Reports
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Report ID"
// @Success 200 {object} response{data=[]dto.ReportPrintResponse}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /report/{id}/prints [get]
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
