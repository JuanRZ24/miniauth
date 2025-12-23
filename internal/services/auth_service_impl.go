package services


import(
	"context"
	"errors"


	"miniauth/internal/models"
	"miniauth/internal/repositories"
	"miniauth/internal/security"
)


type authService struct {
	userRepo repositories.UserRepository
}



func NewAuthService(userRepo repositories.UserRepository) AuthService{
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) Register(ctx context.Context,email, password string) (*models.User, error){
	if email == "" || password == "" {
		return nil, errors.New("invalid credentials")
	}

	existing, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil,err
	}

	if existing != nil {
		return nil, errors.New("USER ALREADY EXISTS")
	}

	hash, err := security.HashPassword(password)
	if err != nil {
		return nil, err
	}

		user := &models.User{
		Email:        email,
		PasswordHash: hash,
		Role:         "user",
		IsActive:     true,
	}
		if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}


