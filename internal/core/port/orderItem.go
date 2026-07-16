package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	GetOrderByID(ctx context.Context, id uuid.UUID) (*domain.Order, error)
	GetOrdersByVisitID(ctx context.Context, visitID uuid.UUID) ([]*domain.Order, error)
	UpdateOrder(ctx context.Context, order *domain.Order) error
}

type OrderService interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
	GetOrderByID(ctx context.Context, id uuid.UUID) (*domain.Order, error)
	GetOrdersByVisitID(ctx context.Context, visitID uuid.UUID) ([]*domain.Order, error)
	UpdateOrder(ctx context.Context, order *domain.Order) error
}
