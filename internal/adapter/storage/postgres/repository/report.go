package repository

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/storage/postgres"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type ReportRepository struct {
	DB *postgres.DB
}

func NewReportRepository(db *postgres.DB) *ReportRepository {
	return &ReportRepository{DB: db}
}

func (r *ReportRepository) CreateReport(ctx context.Context, report *domain.Report) (*domain.Report, error) {
	query, args, err := sq.
		Insert("reports").
		Columns(
			"report_no",
			"visit_id",
			"generated_by",
			"status",
			"pdf_path",
		).
		Values(
			report.ReportNo,
			nullUUID(report.VisitID),
			report.GeneratedBy,
			report.Status,
			nullString(report.PDFPath),
		).
		Suffix(`
			RETURNING
			id,
			generated_at
		`).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("ReportRepo.CreateReport build: %w", err)
	}

	err = r.DB.QueryRow(ctx, query, args...).Scan(
		&report.ID,
		&report.GeneratedAt,
	)
	if err != nil {
		return nil, err
	}

	return report, nil
}

func (r *ReportRepository) GetReportByID(ctx context.Context, id uuid.UUID) (*domain.Report, error) {
	var visitID uuid.NullUUID
	var pdfPath sql.NullString

	query, args, err := sq.
		Select("id", "report_no", "visit_id", "generated_by", "generated_at", "status", "pdf_path").
		From("reports").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("ReportRepo.GetReportByID build: %w", err)
	}

	var report domain.Report
	err = r.DB.QueryRow(ctx, query, args...).Scan(
		&report.ID,
		&report.ReportNo,
		&visitID,
		&report.GeneratedBy,
		&report.GeneratedAt,
		&report.Status,
		&pdfPath,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("report not found")
		}
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	if visitID.Valid {
		report.VisitID = visitID.UUID
	}
	if pdfPath.Valid {
		report.PDFPath = pdfPath.String
	}

	return &report, nil
}

func (r *ReportRepository) GetReportsByVisitID(ctx context.Context, visitID uuid.UUID) ([]*domain.Report, error) {
	var visitIDNull uuid.NullUUID
	var pdfPath sql.NullString

	query, args, err := sq.
		Select("id", "report_no", "visit_id", "generated_by", "generated_at", "status", "pdf_path").
		From("reports").
		Where(sq.Eq{"visit_id": visitID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("ReportRepo.GetReportsByVisitID build: %w", err)
	}

	rows, err := r.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []*domain.Report
	for rows.Next() {
		var report domain.Report
		err := rows.Scan(
			&report.ID,
			&report.ReportNo,
			&visitIDNull,
			&report.GeneratedBy,
			&report.GeneratedAt,
			&report.Status,
			&pdfPath,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		if visitIDNull.Valid {
			report.VisitID = visitIDNull.UUID
		}
		if pdfPath.Valid {
			report.PDFPath = pdfPath.String
		}

		reports = append(reports, &report)
	}

	return reports, nil
}

func (r *ReportRepository) UpdateReport(ctx context.Context, report *domain.Report) error {
	status := nullString(string(report.Status))
	pdfPath := nullString(report.PDFPath)

	query, args, err := sq.
		Update("reports").
		Set("status", sq.Expr("COALESCE(?, status)", status)).
		Set("pdf_path", sq.Expr("COALESCE(?, pdf_path)", pdfPath)).
		Where(sq.Eq{"id": report.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	_, err = r.DB.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update report: %w", err)
	}

	return nil
}

func (r *ReportRepository) CreateReportPrint(ctx context.Context, print *domain.ReportPrint) (*domain.ReportPrint, error) {
	query, args, err := sq.
		Insert("report_print").
		Columns(
			"report_id",
			"printed_by",
			"copy_number",
		).
		Values(
			print.ReportID,
			print.PrintedBy,
			print.CopyNumber,
		).
		Suffix(`
			RETURNING
			id,
			printed_at
		`).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("ReportRepo.CreateReportPrint build: %w", err)
	}

	err = r.DB.QueryRow(ctx, query, args...).Scan(
		&print.ID,
		&print.PrintedAt,
	)
	if err != nil {
		return nil, err
	}

	return print, nil
}

func (r *ReportRepository) GetReportPrintsByReportID(ctx context.Context, reportID uuid.UUID) ([]*domain.ReportPrint, error) {
	query, args, err := sq.
		Select("id", "report_id", "printed_by", "printed_at", "copy_number").
		From("report_print").
		Where(sq.Eq{"report_id": reportID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("ReportRepo.GetReportPrintsByReportID build: %w", err)
	}

	rows, err := r.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prints []*domain.ReportPrint
	for rows.Next() {
		var print domain.ReportPrint
		err := rows.Scan(
			&print.ID,
			&print.ReportID,
			&print.PrintedBy,
			&print.PrintedAt,
			&print.CopyNumber,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		prints = append(prints, &print)
	}

	return prints, nil
}
