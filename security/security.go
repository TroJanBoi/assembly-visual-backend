package security

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Security middleware logic goes here
		// For example, you can set security headers or check authentication

		// Call the next handler in the chain
		c.Next()
	}
}

func GenerateToken(userId int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func HashPassword(password string) string {
	return password
}
