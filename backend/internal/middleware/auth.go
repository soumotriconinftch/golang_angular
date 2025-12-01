package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/szoumoc/golang_angular/internal/auth"
	"github.com/szoumoc/golang_angular/internal/ctxkey"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessCookie, err := r.Cookie("accessToken")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tok, err := auth.ValidateToken(accessCookie.Value)
		if err == nil && tok.Valid {
			claims, ok := tok.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			userID := int64(claims["user_id"].(float64))
			ctx := context.WithValue(r.Context(), ctxkey.UserID, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		// Check if token is expired
		isExpired := false
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				isExpired = true
			} else {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}

		if isExpired {
			// Extract user_id from expired token
			claims, ok := tok.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			expiredUserID := int64(claims["user_id"].(float64))

			// Check Refresh Token
			refreshCookie, err := r.Cookie("refreshToken")
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			refreshTok, err := auth.ValidateRefresh(refreshCookie.Value)
			if err != nil || !refreshTok.Valid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			refreshClaims, ok := refreshTok.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			refreshUserID := int64(refreshClaims["user_id"].(float64))

			// Compare IDs
			if expiredUserID != refreshUserID {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Generate new access token
			newAccess, err := auth.GenerateAccessToken(refreshUserID)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:     "accessToken",
				Value:    newAccess,
				HttpOnly: true,
				SameSite: http.SameSiteStrictMode,
				MaxAge:   15 * 60,
			})

			// Proceed
			ctx := context.WithValue(r.Context(), ctxkey.UserID, refreshUserID)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		// Fallback
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
