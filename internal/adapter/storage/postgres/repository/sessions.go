package repository

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/storage/postgres"
)

type SessionRepository struct {
	DB *postgres.DB
}

func NewSessionRepository(db *postgres.DB) *SessionRepository {
	return &SessionRepository{
		DB: db,
	}
}

func (sr *SessionRepository) Revoke(ctx context.Context, sessionID uuid.UUID) error {
	query, args, err := sq.
		Insert("revoked_sessions").
		Columns("session_id").
		Values(sessionID).
		Suffix("ON CONFLICT (session_id) DO NOTHING").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("SessionRepo.Revoke build: %w", err)
	}

	_, err = sr.DB.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("SessionRepo.Revoke exec: %w", err)
	}

	return nil
}

func (sr *SessionRepository) IsRevoked(ctx context.Context, sessionID uuid.UUID) (bool, error) {
	query, args, err := sq.
		Select("1").
		From("revoked_sessions").
		Where(sq.Eq{"session_id": sessionID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return false, fmt.Errorf("SessionRepo.IsRevoked build: %w", err)
	}

	var exists int
	err = sr.DB.QueryRow(ctx, query, args...).Scan(&exists)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("SessionRepo.IsRevoked scan: %w", err)
	}

	return true, nil
}
