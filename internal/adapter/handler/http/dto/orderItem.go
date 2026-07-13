package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type CreateOrderRequest struct {
	VisitID uuid.UUID `json:"visit_id" binding:"required"`
	TestID  uuid.UUID `json:"test_id" binding:"required"`
	PanelID uuid.UUID `json:"panel_id"`
	Price   float64   `json:"price"`
	Status  string    `json:"status"`
}

type UpdateOrderRequest struct {
	PanelID uuid.UUID          `json:"panel_id"`
	Status  domain.OrderStatus `json:"status"`
	Price   float64            `json:"price"`
}

type OrderResponse struct {
	ID          uuid.UUID          `json:"id"`
	VisitID     uuid.UUID          `json:"visit_id"`
	TestID      uuid.UUID          `json:"test_id"`
	PanelID     *uuid.UUID         `json:"panel_id"`
	Status      domain.OrderStatus `json:"status"`
	Price       *float64           `json:"price"`
	CollectedBy uuid.UUID          `json:"collected_by"`
	CollectedAt time.Time          `json:"collected_at"`
}

func OrderResponseFromDomain(o *domain.Order) *OrderResponse {
	r := &OrderResponse{
		ID:          o.ID,
		VisitID:     o.VisitID,
		TestID:      o.TestID,
		Status:      o.Status,
		CollectedBy: o.CollectedBy,
		CollectedAt: o.CollectedAt,
	}
	if o.PanelID != uuid.Nil {
		r.PanelID = &o.PanelID
	}
	if o.Price != 0 {
		r.Price = &o.Price
	}
	return r
}

func OrdersResponseFromDomain(orders []*domain.Order) []*OrderResponse {
	res := make([]*OrderResponse, 0, len(orders))
	for _, o := range orders {
		res = append(res, OrderResponseFromDomain(o))
	}
	return res
}
