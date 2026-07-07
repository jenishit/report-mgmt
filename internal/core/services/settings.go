package services

import (
	"context"

	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type SettingsService struct {
	repo port.SettingsRepository
}

func NewSettingsService(sr port.SettingsRepository) *SettingsService {
	return &SettingsService{
		repo: sr,
	}
}

func (ss *SettingsService) UpsertSettings(ctx context.Context, s *domain.LabSettings) error {
	return ss.repo.UpsertSettings(ctx, s)
}