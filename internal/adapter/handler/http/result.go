package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/handler/http/dto"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type ResultHandler struct {
	svc port.ResultService
}

func NewResultHandler(svc port.ResultService) *ResultHandler {
	return &ResultHandler{svc: svc}
}

// CreateResult creates batch results for an order
// @Summary Create results
// @Description Create batch test results for an order
// @Tags Results
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.BatchCreateResultRequest true "Batch result details"
// @Success 200 {object} response{data=[]dto.ResultResponse}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /result [post]
func (h *ResultHandler) CreateResult(ctx *gin.Context) {
	var req dto.BatchCreateResultRequest

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

	verifiedBy := req.VerifiedBy
	if verifiedBy == uuid.Nil {
		verifiedBy = userPayload.UserId
	}

	results := make([]*domain.Result, 0, len(req.Results))
	for _, item := range req.Results {
		results = append(results, &domain.Result{
			OrderID:     req.OrderID,
			ParameterID: item.ParameterID,
			ResultValue: item.ResultValue,
			Flag:        item.Flag,
			PerformedBy: userPayload.UserId,
			VerifiedBy:  verifiedBy,
			Remarks:     item.Remarks,
		})
	}

	res, err := h.svc.CreateResults(ctx, results)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, dto.ResultsResponseFromDomain(res))
}

// GetResultByID returns a result by ID
// @Summary Get result by ID
// @Description Get a result's details by ID
// @Tags Results
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Result ID"
// @Success 200 {object} response{data=dto.ResultResponse}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /result/{id} [get]
func (h *ResultHandler) GetResultByID(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		handleError(ctx, domain.ErrInvalidUUID)
		return
	}

	res, err := h.svc.GetResultByID(ctx, uid)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, dto.ResultResponseFromDomain(res))
}

// GetResultsByOrderID returns results for an order
// @Summary List results by order
// @Description Get all results for a given order
// @Tags Results
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param order_id path string true "Order ID"
// @Success 200 {object} response{data=[]dto.ResultResponse}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /result/order/{order_id} [get]
func (h *ResultHandler) GetResultsByOrderID(ctx *gin.Context) {
	id := ctx.Param("order_id")

	uid, err := uuid.Parse(id)
	if err != nil {
		handleError(ctx, domain.ErrInvalidUUID)
		return
	}

	res, err := h.svc.GetResultsByOrderID(ctx, uid)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, dto.ResultsResponseFromDomain(res))
}

// UpdateResult updates a result
// @Summary Update result
// @Description Update an existing result
// @Tags Results
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Result ID"
// @Param request body dto.UpdateResultRequest true "Result details"
// @Success 200 {object} response
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /result/{id} [patch]
func (h *ResultHandler) UpdateResult(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		handleError(ctx, domain.ErrInvalidUUID)
		return
	}

	var req dto.UpdateResultRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	result := &domain.Result{
		ID:          uid,
		ResultValue: req.ResultValue,
		Flag:        req.Flag,
		VerifiedBy:  req.VerifiedBy,
		Remarks:     req.Remarks,
	}

	err = h.svc.UpdateResult(ctx, result)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, gin.H{"message": "Result updated successfully"})
}
