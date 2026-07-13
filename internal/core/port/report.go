package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type ReportRepository interface {
	CreateReport(ctx context.Context, report *domain.Report) (*domain.Report, error)
	GetReportByID(ctx context.Context, id uuid.UUID) (*domain.Report, error)
	GetReportsByVisitID(ctx context.Context, visitID uuid.UUID) ([]*domain.Report, error)
	UpdateReport(ctx context.Context, report *domain.Report) error
	CreateReportPrint(ctx context.Context, print *domain.ReportPrint) (*domain.ReportPrint, error)
	GetReportPrintsByReportID(ctx context.Context, reportID uuid.UUID) ([]*domain.ReportPrint, error)
}

type ReportService interface {
	CreateReport(ctx context.Context, report *domain.Report) (*domain.Report, error)
	GetReportByID(ctx context.Context, id uuid.UUID) (*domain.Report, error)
	GetReportsByVisitID(ctx context.Context, visitID uuid.UUID) ([]*domain.Report, error)
	UpdateReport(ctx context.Context, report *domain.Report) error
	CreateReportPrint(ctx context.Context, print *domain.ReportPrint) (*domain.ReportPrint, error)
	GetReportPrintsByReportID(ctx context.Context, reportID uuid.UUID) ([]*domain.ReportPrint, error)
}
