package auth

import (
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateToken(t *testing.T) {
	tokenStr, err := GenerateToken()
	fmt.Print(tokenStr)
	if err != nil {
		t.Fatalf("uexpected error: %v", err)
	}
	if tokenStr == "" {
		t.Fatalf("empty token")
	}

}

func TestValidateToken(t *testing.T) {
	tokenStr, err := GenerateToken()
	if err != nil {
		t.Fatalf("token generation failed: %v", err)
	}

	token, err := ValidateToken(tokenStr)
	if err != nil {
		t.Fatalf("validation failed: %v", err)
	}

	if !token.Valid {
		t.Fatalf("token marked invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		t.Fatalf("bad claims type")
	}

	if claims["role"] != "admin" {
		t.Fatalf("incorrect role claim")
	}
}
