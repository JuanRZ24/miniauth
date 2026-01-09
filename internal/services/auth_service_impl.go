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

var ErrInvalidCredentials = errors.New("invalid credentials")
func (s *authService) Login(ctx context.Context, email, password string) (*models.User, error){
	if email == "" || password == "" {
		return nil, ErrInvalidCredentials
	}


	existing, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil,ErrInvalidCredentials
	}

	err = security.Compare(existing.PasswordHash, password); 
	if err != nil{
		return nil, ErrInvalidCredentials
	}

	if !existing.IsActive{
		return nil, ErrInvalidCredentials
	}

	return existing, nil
	
}

func (s *authService) Register(ctx context.Context,email, password string) (*models.User, error){
	if email == "" || password == "" {
		return nil, errors.New("invalid credentials")
	}

	err := security.Validate(password)

	if err != nil{
		return nil, err
	}

	hash, err := security.Hash(password)
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


