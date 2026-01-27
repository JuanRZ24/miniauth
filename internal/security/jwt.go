package security


import (
	"os"
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