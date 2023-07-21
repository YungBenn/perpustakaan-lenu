package repository

import (
	"perpustakaan-lenu/pkg/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	sqlc.Querier
}

type SQLRepository struct {
	db *pgxpool.Pool
	*sqlc.Queries
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &SQLRepository{
		db:      db,
		Queries: sqlc.New(db),
	}
}