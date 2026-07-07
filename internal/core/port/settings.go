package port

import (
	"context"

	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type SettingsRepository interface {
	UpsertSettings(ctx context.Context, s *domain.LabSettings) error
	GetSettings(ctx context.Context) (*domain.LabSettings, error)
}

type SettingsService interface {
	UpsertSettings(ctx context.Context, s *domain.LabSettings) error
	GetSettings(ctx context.Context) (*domain.LabSettings, error)
}