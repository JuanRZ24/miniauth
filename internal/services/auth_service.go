package services 


import(
	"context"
	"github.com/google/uuid"

	"miniauth/internal/models"
)


type AuthService interface {
	Register (ctx context.Context, email, password string) (*models.User, error)
	Login (ctx context.Context, email, password string) (*models.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
}


