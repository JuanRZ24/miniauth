package repositories


import (
	"context"
	"database/sql"
	"errors"
	"miniauth/internal/models"
	"github.com/lib/pq"
)

type PostgresEmailVerificationRepository struct {
	db *sql.DB
}

func NewPostgresEmailVerificationRepository(db *sql.DB) *PostgresEmailVerificationRepository {
	return &PostgresEmailVerificationRepository{db: db}
}