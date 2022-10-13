package postgres

import (
	"context"
	"errors"
	"github.com/altercolt/auth/internal/core/auth"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TokenRepository struct {
	db *pgxpool.Pool
}

func NewTokenRepository(db *pgxpool.Pool) auth.TokenRepository {
	return TokenRepository{
		db: db,
	}
}

func (r TokenRepository) Create(ctx context.Context, m auth.Token) error {
	query := `INSERT INTO tokens (id, user_id, access_token, refresh_token)
			VALUES($1, $2, $3, $4);`

	_, err := r.db.Exec(ctx, query, &m.ID, &m.UserID, &m.AccessToken, &m.RefreshToken)
	if err != nil {
		return err
	}

	return nil
}

func (r TokenRepository) Fetch(ctx context.Context, filter auth.Filter) ([]auth.Token, error) {
	query := `SELECT * FROM tokens WHERE ($1 IS NULL OR id = ANY($1))
                     			     AND ($2::int[] IS NULL OR user_id = ANY($2)) 
                     			     AND ($3::varchar[] IS NULL OR acces_token = ANY($3))
                     			     AND ($4::varchar[] IS NULL OR refresh_token = ANY($4))`

	rows, err := r.db.Query(ctx, query, filter.IDs, filter.Users, filter.AccessTokens, filter.RefreshTokens)
	if err != nil {
		return nil, err
	}

	var tokens []auth.Token

	for rows.Next() {
		var token auth.Token
		if err = rows.Scan(&token.ID, &token.UserID, &token.AccessToken, &token.RefreshToken); err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}

	return tokens, nil
}

func (r TokenRepository) Delete(ctx context.Context, id uuid.UUID, userID int) error {
	query := `DELETE FROM tokens where id = $1 AND user_id = $2`

	trx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	res, err := r.db.Exec(ctx, query, id, userID)
	if err != nil {
		return err
	}

	if res.RowsAffected() > 1 {
		if err = trx.Rollback(ctx); err != nil {
			return err
		}
		return errors.New("transaction rolled back, more than one row affected")
	}

	return nil
}

func (r TokenRepository) DeleteAll(ctx context.Context, userID int) error {
	query := `DELETE FROM tokens WHERE user_id = $1`

	_, err := r.db.Exec(ctx, query, userID)

	return err
}
