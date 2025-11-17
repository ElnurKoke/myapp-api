package repository

import (
	"context"
	"elestial/model"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepo interface {
	SaveRefreshToken(ctx context.Context, userID int, token string, expires time.Time) error
	RevokeRefreshToken(ctx context.Context, tokenID int) error
	GetRefreshToken(ctx context.Context, token string) (model.RefreshToken, error)
}

type authRepo struct {
	db *pgxpool.Pool
}

func NewAuth(db *pgxpool.Pool) *authRepo {
	return &authRepo{
		db: db,
	}
}

func (r *authRepo) SaveRefreshToken(ctx context.Context, userID int, token string, expires time.Time) error {
	query := `
		INSERT INTO refresh_tokens (user_id, token, expires_at, created_at)
		VALUES ($1, $2, $3, NOW())
	`
	_, err := r.db.Exec(ctx, query, userID, token, expires)
	if err != nil {
		return fmt.Errorf("SaveRefreshToken failed: %w", err)
	}
	return nil
}

func (r *authRepo) RevokeRefreshToken(ctx context.Context, tokenID int) error {
	query := `
		UPDATE refresh_tokens
		SET revoked = TRUE
		WHERE id = $1
	`
	_, err := r.db.Exec(ctx, query, tokenID)
	if err != nil {
		return fmt.Errorf("RevokeRefreshToken failed: %w", err)
	}
	return nil
}

func (r *authRepo) GetRefreshToken(ctx context.Context, token string) (model.RefreshToken, error) {
	var rt model.RefreshToken

	query := `
		SELECT id, user_id, token, expires_at, created_at, revoked
		FROM refresh_tokens
		WHERE token = $1
	`

	err := r.db.QueryRow(ctx, query, token).Scan(
		&rt.ID,
		&rt.UserID,
		&rt.Token,
		&rt.ExpiresAt,
		&rt.CreatedAt,
		&rt.Revoked,
	)
	if err != nil {
		return model.RefreshToken{}, err
	}

	return rt, nil
}
