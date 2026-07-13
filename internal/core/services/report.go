package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type ReportService struct {
	repo port.ReportRepository
}

func NewReportService(repo port.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) CreateReport(ctx context.Context, report *domain.Report) (*domain.Report, error) {
	return s.repo.CreateReport(ctx, report)
}

func (s *ReportService) GetReportByID(ctx context.Context, id uuid.UUID) (*domain.Report, error) {
	return s.repo.GetReportByID(ctx, id)
}

func (s *ReportService) GetReportsByVisitID(ctx context.Context, visitID uuid.UUID) ([]*domain.Report, error) {
	return s.repo.GetReportsByVisitID(ctx, visitID)
}

func (s *ReportService) UpdateReport(ctx context.Context, report *domain.Report) error {
	return s.repo.UpdateReport(ctx, report)
}

func (s *ReportService) CreateReportPrint(ctx context.Context, print *domain.ReportPrint) (*domain.ReportPrint, error) {
	return s.repo.CreateReportPrint(ctx, print)
}

func (s *ReportService) GetReportPrintsByReportID(ctx context.Context, reportID uuid.UUID) ([]*domain.ReportPrint, error) {
	return s.repo.GetReportPrintsByReportID(ctx, reportID)
}
