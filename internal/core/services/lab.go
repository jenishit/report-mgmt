package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type LabService struct {
	repo port.LabRepository
}

func NewLabService(sr port.LabRepository) *LabService {
	return &LabService{
		repo: sr,
	}
}

func (ss *LabService) InsertLab(ctx context.Context, s *domain.LabSettings) (*domain.LabSettings, error) {
	return ss.repo.InsertLab(ctx, s)
}

func (ss *LabService) GetLabByLabID(ctx context.Context, labID uuid.UUID) (*domain.LabSettings, error) {
	return ss.repo.GetLabByLabID(ctx, labID)
}

func (ss *LabService) UpdateLab(ctx context.Context, s *domain.LabSettings) error {
	return ss.repo.UpdateLab(ctx, s)
}

func (ss *LabService) GetAllLabs(ctx context.Context) ([]*domain.LabSettings, error) {
	return ss.repo.GetAllLabs(ctx)
}
