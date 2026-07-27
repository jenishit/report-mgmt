package repository

import (
	"github.com/jenish-brainztechs/go-backend/internal/adapter/storage/postgres"
)

type TokenRepository struct {
	DB *postgres.DB
}

func NewTokensRepository(db *postgres.DB) *TokenRepository {
	return &TokenRepository{
		DB: db,
	}
}

// func (tr *TokenRepository) CreateToken(ctx context.Context, ) (*uuid.UUID, error) {
// 	query, args, err := sq.
// 	Insert("TOKENS").
// 	Columns(
// 		"userID",
// 		"token",
// 		"expires_at",
// 	).
// 	Values(
// 		userID,
// 		Token,
// 		time.Now().UTC().Add(15*time.Minute),
// 	).
// 	Where(sq.Eq{"userID": userID}).
// 	Suffix(`RETURNING
// 	token
// 	`).
// 	PlaceholderFormat(sq.Dollar).
// 	ToSql()
// }
