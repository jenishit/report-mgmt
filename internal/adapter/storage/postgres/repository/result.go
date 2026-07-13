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

type ResultRepository struct {
	DB *postgres.DB
}

func NewResultRepository(db *postgres.DB) *ResultRepository {
	return &ResultRepository{DB: db}
}

func (r *ResultRepository) CreateResult(ctx context.Context, result *domain.Result) (*domain.Result, error) {
	query, args, err := sq.
		Insert("result").
		Columns(
			"order_id",
			"parameter_id",
			"result_value",
			"flag",
			"performed_by",
			"verified_by",
			"remarks",
		).
		Values(
			result.OrderID,
			result.ParameterID,
			result.ResultValue,
			result.Flag,
			result.PerformedBy,
			result.VerifiedBy,
			nullString(result.Remarks),
		).
		Suffix(`
			RETURNING
			id,
			performed_at
		`).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("ResultRepo.CreateResult build: %w", err)
	}

	err = r.DB.QueryRow(ctx, query, args...).Scan(
		&result.ID,
		&result.PerformedAt,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *ResultRepository) CreateResults(ctx context.Context, results []*domain.Result) ([]*domain.Result, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("ResultRepo.CreateResults begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	for _, result := range results {
		query, args, err := sq.
			Insert("result").
			Columns(
				"order_id",
				"parameter_id",
				"result_value",
				"flag",
				"performed_by",
				"verified_by",
				"remarks",
			).
			Values(
				result.OrderID,
				result.ParameterID,
				result.ResultValue,
				result.Flag,
				result.PerformedBy,
				result.VerifiedBy,
				nullString(result.Remarks),
			).
			Suffix(`
				RETURNING
				id,
				performed_at
			`).
			PlaceholderFormat(sq.Dollar).
			ToSql()

		if err != nil {
			return nil, fmt.Errorf("ResultRepo.CreateResults build: %w", err)
		}

		err = tx.QueryRow(ctx, query, args...).Scan(
			&result.ID,
			&result.PerformedAt,
		)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("ResultRepo.CreateResults commit: %w", err)
	}

	return results, nil
}

func (r *ResultRepository) GetResultByID(ctx context.Context, id uuid.UUID) (*domain.Result, error) {
	var remarks sql.NullString

	query, args, err := sq.
		Select("id", "order_id", "parameter_id", "result_value", "flag", "performed_by", "performed_at", "verified_by", "remarks").
		From("result").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("ResultRepo.GetResultByID build: %w", err)
	}

	var result domain.Result
	err = r.DB.QueryRow(ctx, query, args...).Scan(
		&result.ID,
		&result.OrderID,
		&result.ParameterID,
		&result.ResultValue,
		&result.Flag,
		&result.PerformedBy,
		&result.PerformedAt,
		&result.VerifiedBy,
		&remarks,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("result not found")
		}
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	if remarks.Valid {
		result.Remarks = remarks.String
	}

	return &result, nil
}

func (r *ResultRepository) GetResultsByOrderID(ctx context.Context, orderID uuid.UUID) ([]*domain.Result, error) {
	var remarks sql.NullString

	query, args, err := sq.
		Select("id", "order_id", "parameter_id", "result_value", "flag", "performed_by", "performed_at", "verified_by", "remarks").
		From("result").
		Where(sq.Eq{"order_id": orderID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("ResultRepo.GetResultsByOrderID build: %w", err)
	}

	rows, err := r.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*domain.Result
	for rows.Next() {
		var result domain.Result
		err := rows.Scan(
			&result.ID,
			&result.OrderID,
			&result.ParameterID,
			&result.ResultValue,
			&result.Flag,
			&result.PerformedBy,
			&result.PerformedAt,
			&result.VerifiedBy,
			&remarks,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		if remarks.Valid {
			result.Remarks = remarks.String
		}

		results = append(results, &result)
	}

	return results, nil
}

func (r *ResultRepository) UpdateResult(ctx context.Context, result *domain.Result) error {
	resultValue := nullString(result.ResultValue)
	flag := nullString(string(result.Flag))
	remarks := nullString(result.Remarks)
	verifiedBy := nullUUID(result.VerifiedBy)

	query, args, err := sq.
		Update("result").
		Set("result_value", sq.Expr("COALESCE(?, result_value)", resultValue)).
		Set("flag", sq.Expr("COALESCE(?, flag)", flag)).
		Set("remarks", sq.Expr("COALESCE(?, remarks)", remarks)).
		Set("verified_by", sq.Expr("COALESCE(?, verified_by)", verifiedBy)).
		Set("updated_at", sq.Expr("NOW()")).
		Where(sq.Eq{"id": result.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	_, err = r.DB.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update result: %w", err)
	}

	return nil
}
