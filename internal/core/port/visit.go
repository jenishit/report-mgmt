package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type VisitRepository interface {
	CreateVisit(ctx context.Context, visit *domain.Visit) (*domain.Visit, error)
	GetVisitByID(ctx context.Context, id uuid.UUID) ([]*domain.ListVisit, error)
}

type VisitService interface {
	CreateVisit(ctx context.Context, visit *domain.Visit) (*domain.Visit, error)
	GetVisitByID(ctx context.Context, id uuid.UUID) ([]*domain.ListVisit, error)
}
