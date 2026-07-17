package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/handler/http/dto"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type OrderHandler struct {
	svc port.OrderService
}

func NewOrderHandler(svc port.OrderService) *OrderHandler {
	return &OrderHandler{svc: svc}
}

// CreateOrder creates a new order
// @Summary Create order
// @Description Create a new test order for a visit
// @Tags Orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateOrderRequest true "Order details"
// @Success 200 {object} response{data=dto.OrderResponse}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /order [post]
func (h *OrderHandler) CreateOrder(ctx *gin.Context) {
	var req dto.CreateOrderRequest

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

	order := &domain.Order{
		VisitID:     req.VisitID,
		TestID:      req.TestID,
		PanelID:     req.PanelID,
		Price:       req.Price,
		Status:      domain.OrderStatus(req.Status),
		CollectedBy: userPayload.UserId,
	}

	res, err := h.svc.CreateOrder(ctx, order)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, dto.OrderResponseFromDomain(res))
}

// GetOrderByID returns an order by ID
// @Summary Get order by ID
// @Description Get an order's details by ID
// @Tags Orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Success 200 {object} response{data=dto.OrderResponse}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /order/{id} [get]
func (h *OrderHandler) GetOrderByID(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		handleError(ctx, domain.ErrInvalidUUID)
		return
	}

	res, err := h.svc.GetOrderByID(ctx, uid)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, dto.OrderResponseFromDomain(res))
}

// GetOrdersByVisitID returns orders for a visit
// @Summary List orders by visit
// @Description Get all orders for a given visit
// @Tags Orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param visit_id path string true "Visit ID"
// @Success 200 {object} response{data=[]dto.OrderResponse}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /order/visit/{visit_id} [get]
func (h *OrderHandler) GetOrdersByVisitID(ctx *gin.Context) {
	id := ctx.Param("visit_id")

	uid, err := uuid.Parse(id)
	if err != nil {
		handleError(ctx, domain.ErrInvalidUUID)
		return
	}

	res, err := h.svc.GetOrdersByVisitID(ctx, uid)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, dto.OrdersResponseFromDomain(res))
}

// UpdateOrder updates an order
// @Summary Update order
// @Description Update an existing order
// @Tags Orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Param request body dto.UpdateOrderRequest true "Order details"
// @Success 200 {object} response
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /order/{id} [patch]
func (h *OrderHandler) UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		handleError(ctx, domain.ErrInvalidUUID)
		return
	}

	var req dto.UpdateOrderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	order := &domain.Order{
		ID:      uid,
		PanelID: req.PanelID,
		Status:  req.Status,
		Price:   req.Price,
	}

	err = h.svc.UpdateOrder(ctx, order)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, gin.H{"message": "Order updated successfully"})
}
