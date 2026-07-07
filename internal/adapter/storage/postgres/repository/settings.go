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

type SettingsRepository struct {
	DB *postgres.DB
}

func NewSettingsRepository(db *postgres.DB) *SettingsRepository {
	return &SettingsRepository{
		DB: db,
	}
}

func (sr *SettingsRepository) UpsertSettings(ctx context.Context, s *domain.LabSettings, updatedBy uuid.UUID) error {
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
		return fmt.Errorf("upserting lab settings: %w", err)
	}

	return nil
}

func (sr *SettingsRepository) GetSettings(ctx context.Context) (*domain.LabSettings, error) {
	var phone, email, tagline, address, registrationNo, reportFooter sql.NullString

	query, args, err := sq.
		Select(
			"ID",
			"lab_name",
			"tagline",
			"address",
			"phone",
			"email",
			"registration_no",
			"report_footer",
		).From("lab_settings").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	rows, err := sr.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close() //Release resources after reading the rows

	setting := &domain.LabSettings{}
	err = sr.DB.QueryRow(ctx, query, args...).Scan(
		&setting.ID,
		&setting.LabName,
		&tagline,
		&address,
		&phone,
		&email,
		&registrationNo,
		&reportFooter,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	if tagline.Valid {
		setting.Tagline = &tagline.String
	}
	if address.Valid {
		setting.Address = &address.String
	}
	if email.Valid {
		setting.Email = &email.String
	}
	if phone.Valid {
		setting.Phone = &phone.String
	}
	if registrationNo.Valid {
		setting.RegistrationNo = &registrationNo.String
	}
	if reportFooter.Valid {
		setting.ReportFooter = &reportFooter.String
	}

	return setting, nil
}
