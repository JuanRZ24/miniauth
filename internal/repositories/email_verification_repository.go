
package repositories

import (
	"context"
	"miniauth/internal/models"
)

type EmailVerificationRepository interface {
	Create(ctx context.Context, token *models.EmailVerificationToken) error
	FindActiveByTokenHash(ctx context.Context, hash string) (*models.EmailVerificationToken, error)
	MarkUsed(ctx context.Context, tokenID string) error
	DeleteExpired(ctx context.Context) error
}
