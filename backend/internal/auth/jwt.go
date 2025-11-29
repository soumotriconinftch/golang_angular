package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("abcdef")
var username = "soumo"

func GenerateToken() (string, error) {
	fmt.Printf("calling token func")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"role":     "admin",
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	fmt.Print(tokenString)
	return tokenString, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return secretKey, nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
}
