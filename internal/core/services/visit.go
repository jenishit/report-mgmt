package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type VisitService struct {
	repo port.VisitRepository
}

func NewVisitService(vr port.VisitRepository) *VisitService {
	return &VisitService{
		repo: vr,
	}
}

func (v *VisitService) CreateVisit(ctx context.Context, visit *domain.Visit) (*domain.Visit, error) {
	return v.repo.CreateVisit(ctx, visit)
}

func (v *VisitService) GetVisitByID(ctx context.Context, id uuid.UUID) ([]*domain.ListVisit, error) {
	return v.repo.GetVisitByID(ctx, id)
}
