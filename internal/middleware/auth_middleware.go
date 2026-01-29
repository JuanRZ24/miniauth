package middleware

import(
	"github.com/gin-gonic/gin"
	"strings"
	"miniauth/internal/security"
)
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			c.AbortWithStatusJSON(401,gin.H{"error":"missing token"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer"{
			c.AbortWithStatusJSON(401,gin.H{"error": "invalida token format"})
			return
		}

		token := parts[1]

		userID, err := security.ParseToken(token)

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error":"invalid token"})
			return
		}

		c.Set("userID", userID)
		c.Next()
	}


}