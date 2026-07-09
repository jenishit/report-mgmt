package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/storage/postgres"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type LabRepository struct {
	DB *postgres.DB
}

func NewLabRepository(db *postgres.DB) *LabRepository {
	return &LabRepository{
		DB: db,
	}
}

func (sr *LabRepository) InsertLab(ctx context.Context, s *domain.LabSettings) error {
	now := time.Now()

	query, args, err := sq.
		Insert("lab_settings").
		Columns(
			"lab_name",
			"tagline",
			"address",
			"phone",
			"email",
			"registration_no",
			"report_footer",
			"updated_at",
			"updated_by ",
		).
		Values(
			s.LabName,
			s.Tagline,
			s.Address,
			s.Phone,
			s.Email,
			s.RegistrationNo,
			s.ReportFooter,
			now,
			s.UpdatedBy,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	_, err = sr.DB.Exec(ctx, query, args...)

	if err != nil {
		return fmt.Errorf("inserting lab preference: %w", err)
	}

	return nil
}

func (sr *LabRepository) GetLabByLabID(ctx context.Context, labID uuid.UUID) (*domain.LabSettings, error) {
	var phone, email, tagline, address, registrationNo, reportFooter sql.NullString

	query, args, err := sq.
		Select(
			"id",
			"lab_name",
			"tagline",
			"address",
			"phone",
			"email",
			"registration_no",
			"report_footer",
			"updated_by",
		).From("lab_settings").
		Where(sq.Eq{"id": labID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	rows, err := sr.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close() //Release resources after reading the rows

	lab := &domain.LabSettings{}
	err = sr.DB.QueryRow(ctx, query, args...).Scan(
		&lab.ID,
		&lab.LabName,
		&tagline,
		&address,
		&phone,
		&email,
		&registrationNo,
		&reportFooter,
		&lab.UpdatedBy,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	if tagline.Valid {
		lab.Tagline = &tagline.String
	}
	if address.Valid {
		lab.Address = &address.String
	}
	if email.Valid {
		lab.Email = &email.String
	}
	if phone.Valid {
		lab.Phone = &phone.String
	}
	if registrationNo.Valid {
		lab.RegistrationNo = &registrationNo.String
	}
	if reportFooter.Valid {
		lab.ReportFooter = &reportFooter.String
	}

	return lab, nil
}

func (sr *LabRepository) UpdateLab(ctx context.Context, s *domain.LabSettings) error {
	now := time.Now()

	builder := sq.Update("lab_settings").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": s.ID})

	if s.LabName != "" {
		builder = builder.Set("lab_name", s.LabName)
	}
	if s.Tagline != nil {
		builder = builder.Set("tagline", s.Tagline)
	}
	if s.Address != nil {
		builder = builder.Set("address", s.Address)
	}
	if s.Phone != nil {
		builder = builder.Set("phone", s.Phone)
	}

	if s.ReportFooter != nil {
		builder = builder.Set("report_footer", s.ReportFooter)
	}

	builder = builder.Set("updated_at", now)
	builder = builder.Set("updated_by", s.UpdatedBy)

	query, args, err := builder.ToSql()

	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	_, err = sr.DB.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update profile: %w", err)
	}

	return nil
}
