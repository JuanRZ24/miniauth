package repositories


import (
	"context"
	"github.com/google/uuid"

	"miniauth/internal/models"
)



type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*models.User, error)
}

