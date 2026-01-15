package models

import "time"


type EmailVerificationToken struct {
	ID		  string
	UserID    string
	TokenHash string
	ExpiresAt time.Time
	UsedAt *time.Time
	CreatedAt time.Time

}