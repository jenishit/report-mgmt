package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/storage/postgres"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type DoctorRepository struct {
	DB *postgres.DB
}

func NewDoctorRepository(db *postgres.DB) *DoctorRepository {
	return &DoctorRepository{
		DB: db,
	}
}

func (dr *DoctorRepository) CreateDoctor(ctx context.Context, doc *domain.Doctor) (*domain.Doctor, error) {
	query, args, err := sq.
		Insert("doctors").
		Columns(
			"first_name",
			"last_name",
			"phone",
			"email",
			"qualification",
			"registration_no",
		).
		Values(
			doc.FirstName,
			doc.LastName,
			doc.Phone,
			doc.Email,
			doc.Qualification,
			doc.RegistrationNo,
		).
		Suffix(`
		RETURNING
		id
		`).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("DoctorRepo.CreateDoctor build: %w", err)
	}

	err = dr.DB.QueryRow(ctx, query, args...).Scan(
		&doc.ID,
	)

	if err != nil {
		return nil, err
	}

	return doc, nil
}

func (dr *DoctorRepository) GetDoctorByID(ctx context.Context, id uuid.UUID) (*domain.Doctor, error) {
	query, args, err := sq.
		Select(
			"id",
			"first_name",
			"last_name",
			"phone",
			"email",
			"qualification",
			"registration_no",
			"created_at",
		).From("doctors").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	rows, err := dr.DB.Query(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	doc := &domain.Doctor{}

	err = dr.DB.QueryRow(ctx, query, args...).Scan(
		&doc.ID,
		&doc.FirstName,
		&doc.LastName,
		&doc.Phone,
		&doc.Email,
		&doc.Qualification,
		&doc.RegistrationNo,
		&doc.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return doc, nil
}

func (dr *DoctorRepository) GetDoctors(ctx context.Context) ([]*domain.Doctor, error) {
	query, args, err := sq.
		Select(
			"id",
			"first_name",
			"last_name",
			"phone",
			"email",
			"qualification",
			"registration_no",
			"created_at",
		).From("doctors").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := dr.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("querying doctors: %w", err)
	}
	defer rows.Close()

	var docs []*domain.Doctor

	for rows.Next() {
		doc := &domain.Doctor{}
		err = dr.DB.QueryRow(ctx, query, args...).Scan(
			&doc.ID,
			&doc.FirstName,
			&doc.LastName,
			&doc.Phone,
			&doc.Email,
			&doc.Qualification,
			&doc.RegistrationNo,
			&doc.CreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		docs = append(docs, doc)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating doctors rows: %w", err)
	}

	return docs, nil
}

func (dr *DoctorRepository) UpdateDoctor(ctx context.Context, doc *domain.Doctor) error {
	builder := sq.Update("doctors").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": doc.ID})

	if doc.FirstName != "" {
		builder = builder.Set("first_name", doc.FirstName)
	}

	if doc.LastName != "" {
		builder = builder.Set("last_name", doc.LastName)
	}

	if doc.Phone != nil {
		builder = builder.Set("phone", doc.Phone)
	}

	if doc.Email != "" {
		builder = builder.Set("email", doc.Email)
	}

	if doc.Qualification != "" {
		builder = builder.Set("qualification", doc.Qualification)
	}

	if doc.RegistrationNo != "" {
		builder = builder.Set("registration_no", doc.RegistrationNo)
	}

	builder = builder.Set("updated_by", doc.UpdatedBy)

	query, args, err := builder.ToSql()

	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	_, err = dr.DB.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update doctor: %w", err)
	}

	return nil
}