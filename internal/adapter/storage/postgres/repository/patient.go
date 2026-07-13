package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/storage/postgres"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type PatientRepository struct {
	DB *postgres.DB
}

func NewPatientRepository(db *postgres.DB) *PatientRepository {
	return &PatientRepository{
		DB: db,
	}
}

func (pr *PatientRepository) CreatePatient(ctx context.Context, pt *domain.Patient) (*domain.Patient, error) {
	query, args, err := sq.
		Insert("patients").
		Columns(
			"first_name",
			"last_name",
			"phone",
			"email",
			"mrn",
			"dob",
			"gender",
			"p_address",
		).
		Values(
			pt.FirstName,
			pt.LastName,
			pt.Phone,
			pt.Email,
			pt.MRN,
			pt.DOB,
			pt.Gender,
			pt.Address,
		).
		Suffix(`
		RETURNING
		id
		`).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("PatientRepo.CreatePatient build: %w", err)
	}

	err = pr.DB.QueryRow(ctx, query, args...).Scan(
		&pt.ID,
	)

	if err != nil {
		return nil, err
	}

	return pt, nil
}

func (pr *PatientRepository) GetPatientByID(ctx context.Context, id uuid.UUID) (*domain.Patient, error) {
	query, args, err := sq.
		Select(
			"id",
			"first_name",
			"last_name",
			"phone",
			"email",
			"mrn",
			"dob",
			"gender",
			"p_address",
			"created_at",
		).From("patients").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	rows, err := pr.DB.Query(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	pt := &domain.Patient{}

	err = pr.DB.QueryRow(ctx, query, args...).Scan(
		&pt.ID,
		&pt.FirstName,
		&pt.LastName,
		&pt.Phone,
		&pt.Email,
		&pt.MRN,
		&pt.DOB,
		&pt.Gender,
		&pt.Address,
		&pt.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return pt, nil
}

func (pr *PatientRepository) GetPatients(ctx context.Context) ([]*domain.Patient, error) {
	query, args, err := sq.
		Select(
			"id",
			"first_name",
			"last_name",
			"phone",
			"email",
			"mrn",
			"dob",
			"gender",
			"p_address",
			"created_at",
		).From("patients").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := pr.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("querying patients: %w", err)
	}
	defer rows.Close()

	var pts []*domain.Patient

	for rows.Next() {
		pt := &domain.Patient{}
		err = rows.Scan(
			&pt.ID,
			&pt.FirstName,
			&pt.LastName,
			&pt.Phone,
			&pt.Email,
			&pt.MRN,
			&pt.DOB,
			&pt.Gender,
			&pt.Address,
			&pt.CreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		pts = append(pts, pt)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating patients rows: %w", err)
	}

	return pts, nil
}

func (pr *PatientRepository) UpdatePatient(ctx context.Context, pt *domain.Patient) error {
	builder := sq.Update("patients").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": pt.ID})

	if pt.FirstName != "" {
		builder = builder.Set("first_name", pt.FirstName)
	}

	if pt.LastName != "" {
		builder = builder.Set("last_name", pt.LastName)
	}

	if pt.Phone != nil {
		builder = builder.Set("phone", pt.Phone)
	}

	if pt.Email != "" {
		builder = builder.Set("email", pt.Email)
	}

	if pt.MRN != "" {
		builder = builder.Set("mrn", pt.MRN)
	}

	if !pt.DOB.IsZero() {
		builder = builder.Set("dob", pt.DOB)
	}

	if pt.Gender != "" {
		builder = builder.Set("gender", pt.Gender)
	}

	if pt.Address != "" {
		builder = builder.Set("p_address", pt.Address)
	}

	builder = builder.Set("updated_by", pt.UpdatedBy)

	query, args, err := builder.ToSql()

	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	_, err = pr.DB.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update patient: %w", err)
	}

	return nil
}
