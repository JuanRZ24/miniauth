package services 


import(
	"context"

	"miniauth/internal/models"
)


type AuthService interface {
	Register (ctx context.Context, email, password string) (*models.User, error)
}


