package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type LabRepository interface {
	InsertLab(ctx context.Context, s *domain.LabSettings) error
	GetLabByLabID(ctx context.Context, labID uuid.UUID) (*domain.LabSettings, error)
	UpdateLab(ctx context.Context, s *domain.LabSettings) error
}

type LabService interface {
	InsertLab(ctx context.Context, s *domain.LabSettings) error
	GetLabByLabID(ctx context.Context, labID uuid.UUID) (*domain.LabSettings, error)
	UpdateLab(ctx context.Context, s *domain.LabSettings) error
}
