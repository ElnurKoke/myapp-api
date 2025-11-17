package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	AuthRepo
	UserRepo
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		AuthRepo: NewAuth(db),
		UserRepo: NewUser(db),
	}
}
