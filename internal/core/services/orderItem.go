package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type OrderService struct {
	repo port.OrderRepository
}

func NewOrderService(repo port.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	return s.repo.CreateOrder(ctx, order)
}

func (s *OrderService) GetOrderByID(ctx context.Context, id uuid.UUID) (*domain.Order, error) {
	return s.repo.GetOrderByID(ctx, id)
}

func (s *OrderService) GetOrdersByVisitID(ctx context.Context, visitID uuid.UUID) ([]*domain.Order, error) {
	return s.repo.GetOrdersByVisitID(ctx, visitID)
}

func (s *OrderService) UpdateOrder(ctx context.Context, order *domain.Order) error {
	return s.repo.UpdateOrder(ctx, order)
}
