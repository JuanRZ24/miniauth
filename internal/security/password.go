package security

import (
	"errors" 
	"strings"
	"golang.org/x/crypto/bcrypt")


func Validate(password string) error {
	if strings.TrimSpace(password) == "" {
		return errors.New("Password Invalid")
	}

	if len(password) < 8 {
		return errors.New("Password too short")
	}

	if len(password) > 78{
		return errors.New("Password too long")
	}

	return nil

}


func Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func Compare(hash string, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}
