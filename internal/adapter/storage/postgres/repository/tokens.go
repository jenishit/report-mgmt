package repository

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/storage/postgres"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type TokenRepository struct {
	DB *postgres.DB
}

func NewTokensRepository(db *postgres.DB) *TokenRepository {
	return &TokenRepository{
		DB: db,
	}
}

func (tr *TokenRepository) CreateOTP(ctx context.Context, userID uuid.UUID, otp string, expiresAt time.Time) error {
	query, args, err := sq.
		Insert("tokens").
		Columns(
			"userid",
			"token",
			"otp",
			"expires_at",
		).
		Values(
			userID,
			uuid.New(),
			otp,
			expiresAt,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("TokenRepo.CreateOTP build: %w", err)
	}

	_, err = tr.DB.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("TokenRepo.CreateOTP exec: %w", err)
	}

	return nil
}

func (tr *TokenRepository) GetOTP(ctx context.Context, email, otp string) (*domain.Token, error) {
	var t domain.Token
	query, args, err := sq.
		Select(
			"t.id",
			"t.userid",
			"t.token",
			"t.otp",
			"t.expires_at",
			"t.created_at",
			"t.used_at",
		).
		From("tokens t").
		Join("users u on u.id = t.userid").
		Where(sq.Eq{"u.email": email}).
		Where(sq.Eq{"t.otp": otp}).
		Where("t.used_at IS NULL").
		Where(sq.Gt{"t.expires_at": time.Now().UTC()}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("TokenRepo.GetOTP build: %w", err)
	}

	err = tr.DB.QueryRow(ctx, query, args...).Scan(
		&t.ID,
		&t.UserID,
		&t.Token,
		&t.OTP,
		&t.ExpiresAt,
		&t.CreatedAt,
		&t.UsedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("TokenRepo.GetOTP scan: %w", err)
	}

	return &t, nil
}

func (tr *TokenRepository) MarkUsed(ctx context.Context, id uuid.UUID) error {
	query, args, err := sq.
		Update("tokens").
		Set("used_at", time.Now()).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("TokenRepo.MarkUsed build: %w", err)
	}

	_, err = tr.DB.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("TokenRepo.MarkUsed exec: %w", err)
	}

	return nil
}
