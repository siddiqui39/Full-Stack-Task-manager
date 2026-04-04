
package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte("your_secret_key_here") // replace with a strong secret

// GenerateToken generates a JWT token for a user
func GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(JwtSecret)
	return tokenString, err
}