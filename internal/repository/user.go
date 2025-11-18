package repository

import (
	"context"
	"elestial/model"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user model.RegisterRequest) error
	UpdateUser(ctx context.Context, user model.User) error
	DeleteUser(ctx context.Context, user model.User) error

	GetUserById(ctx context.Context, userID int) (model.User, error)
	GetUserByName(ctx context.Context, username string) (model.User, error)
}

type userRepo struct {
	db *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user model.RegisterRequest) error {
	query := `
		INSERT INTO users (name, email, password_hash, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id
	`
	var id int
	err := r.db.QueryRow(ctx, query,
		user.Name,
		user.Email,
		user.Password,
	).Scan(&id)
	if err != nil {
		return fmt.Errorf("CreateUser failed: %w", err)
	}

	return nil
}

func (r *userRepo) UpdateUser(ctx context.Context, user model.User) error {
	query := `
		UPDATE users
		SET name=$1, email=$2, password_hash=$3, updated_at=NOW()
		WHERE id=$4
	`
	_, err := r.db.Exec(ctx, query,
		user.Name,
		user.Email,
		user.Password,
		user.ID,
	)
	if err != nil {
		return fmt.Errorf("UpdateUser failed: %w", err)
	}
	return nil
}

func (r *userRepo) DeleteUser(ctx context.Context, user model.User) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := r.db.Exec(ctx, query, user.ID)
	if err != nil {
		return fmt.Errorf("DeleteUser failed: %w", err)
	}
	return nil
}

func (r *userRepo) GetUserById(ctx context.Context, userID int) (model.User, error) {
	var user model.User
	query := `
		SELECT id, name, email, password_hash,created_at, updated_at
		FROM users
		WHERE id=$1
	`
	row := r.db.QueryRow(ctx, query, userID)
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return model.User{}, fmt.Errorf("GetUserById failed: %w", err)
	}
	return user, nil
}

func (r *userRepo) GetUserByName(ctx context.Context, username string) (model.User, error) {
	var user model.User
	query := `
		SELECT id, name, email, password_hash, created_at, updated_at
		FROM users
		WHERE name=$1
	`
	row := r.db.QueryRow(ctx, query, username)
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return model.User{}, fmt.Errorf("GetUserByName failed: %w", err)
	}
	return user, nil
}
