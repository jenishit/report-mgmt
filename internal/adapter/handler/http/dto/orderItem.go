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

type ListOrders struct {
	ID          uuid.UUID          `json:"id"`
	VisitNo     string             `json:"visit_no"`
	PatientName string             `json:"patinet_name"`
	TestName    string             `json:"test_name"`
	TestCode    string             `json:"test_code"`
	TestPrice   float64            `json:"test_price"`
	PanelName   string             `json:"panel_name"`
	PanelCode   string             `json:"panel_code"`
	PanelPrice  float64            `json:"panel_price"`
	Status      domain.OrderStatus `json:"status"`
	Price       float64            `json:"price"`
	CollectedBy string             `json:"collected_by"`
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

func ListOrderRes(lo *domain.ListOrders) *ListOrders {
	return &ListOrders{
		ID:          lo.ID,
		VisitNo:     lo.VisitNo,
		PatientName: lo.PtFirstName + lo.PtLastName,
		TestName:    lo.TestName,
		TestCode:    lo.TestCode,
		TestPrice:   lo.TestPrice,
		PanelName:   lo.PanelName,
		PanelCode:   lo.PanelCode,
		PanelPrice:  lo.PanelPrice,
		Status:      lo.Status,
		Price:       lo.Price,
		CollectedBy: lo.CollectorFirstName + lo.CollectorLastName,
	}
}

func ListOrderResponse(listOrders []*domain.ListOrders) []*ListOrders {
	orders := make([]*ListOrders, 0, len(listOrders))

	for _, lo := range listOrders {
		orders = append(orders, &ListOrders{
			ID:          lo.ID,
			VisitNo:     lo.VisitNo,
			PatientName: lo.PtFirstName + lo.PtLastName,
			TestName:    lo.TestName,
			TestCode:    lo.TestCode,
			TestPrice:   lo.TestPrice,
			PanelName:   lo.PanelName,
			PanelCode:   lo.PanelCode,
			PanelPrice:  lo.PanelPrice,
			Status:      lo.Status,
			Price:       lo.Price,
			CollectedBy: lo.CollectorFirstName + lo.CollectorLastName,
		})
	}
	return orders
}
