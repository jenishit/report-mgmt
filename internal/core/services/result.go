package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type ResultService struct {
	repo port.ResultRepository
}

func NewResultService(repo port.ResultRepository) *ResultService {
	return &ResultService{repo: repo}
}

func (s *ResultService) CreateResult(ctx context.Context, result *domain.Result) (*domain.Result, error) {
	return s.repo.CreateResult(ctx, result)
}

func (s *ResultService) CreateResults(ctx context.Context, results []*domain.Result) ([]*domain.Result, error) {
	return s.repo.CreateResults(ctx, results)
}

func (s *ResultService) GetResultByID(ctx context.Context, id uuid.UUID) (*domain.Result, error) {
	return s.repo.GetResultByID(ctx, id)
}

func (s *ResultService) GetResultsByOrderID(ctx context.Context, orderID uuid.UUID) ([]*domain.Result, error) {
	return s.repo.GetResultsByOrderID(ctx, orderID)
}

func (s *ResultService) UpdateResult(ctx context.Context, result *domain.Result) error {
	return s.repo.UpdateResult(ctx, result)
}
