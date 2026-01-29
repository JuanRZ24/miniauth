package security


import (
	"os"
	"time"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	jwt.RegisteredClaims
}


func GenerateToken(userID uuid.UUID) (string,error){
	claims := Claims{
		RegisteredClaims : jwt.RegisteredClaims{
			Subject: userID.String(),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15*time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}



func ParseToken(tokenString string) (uuid.UUID, error){
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},

	)

	if err != nil || !token.Valid {
		return uuid.Nil , errors.New("invalid token")
	}

	claims := token.Claims.(*Claims)
	return uuid.Parse(claims.Subject)
}