package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type ResultRepository interface {
	CreateResult(ctx context.Context, result *domain.Result) (*domain.Result, error)
	CreateResults(ctx context.Context, results []*domain.Result) ([]*domain.Result, error)
	GetResultByID(ctx context.Context, id uuid.UUID) (*domain.Result, error)
	GetResultsByOrderID(ctx context.Context, orderID uuid.UUID) ([]*domain.Result, error)
	UpdateResult(ctx context.Context, result *domain.Result) error
}

type ResultService interface {
	CreateResult(ctx context.Context, result *domain.Result) (*domain.Result, error)
	CreateResults(ctx context.Context, results []*domain.Result) ([]*domain.Result, error)
	GetResultByID(ctx context.Context, id uuid.UUID) (*domain.Result, error)
	GetResultsByOrderID(ctx context.Context, orderID uuid.UUID) ([]*domain.Result, error)
	UpdateResult(ctx context.Context, result *domain.Result) error
}
