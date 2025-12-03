package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/szoumoc/golang_angular/internal/auth"
	"github.com/szoumoc/golang_angular/internal/ctxkey"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessCookie, err := r.Cookie("accessToken")
		if err != nil {
			log.Print("failed to get accessToken")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tok, err := auth.ValidateToken(accessCookie.Value)
		if err == nil && tok.Valid {
			claims, ok := tok.Claims.(jwt.MapClaims)
			if !ok {
				log.Print("failed to get claims")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			userID := int64(claims["user_id"].(float64))
			isAdmin := claims["is_admin"].(bool)
			ctx := context.WithValue(r.Context(), ctxkey.UserID, userID)
			ctx1 := context.WithValue(ctx, ctxkey.IsAdmin, isAdmin)
			next.ServeHTTP(w, r.WithContext(ctx1))
			return
		}

		isExpired := false
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				isExpired = true
			} else {
				log.Print("Wrong JWT")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}

		if isExpired {

			claims, ok := tok.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			expiredUserID := int64(claims["user_id"].(float64))
			expiredisAdmin := claims["is_admin"].(bool)

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

			if expiredUserID != refreshUserID {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			newAccess, err := auth.GenerateAccessToken(refreshUserID, expiredisAdmin)
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

			// isAdmin := claims["is_admin"].(bool)
			ctx := context.WithValue(r.Context(), ctxkey.UserID, refreshUserID)
			ctx1 := context.WithValue(ctx, ctxkey.IsAdmin, expiredisAdmin)
			next.ServeHTTP(w, r.WithContext(ctx1))
			return
		}

		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func Authorization(checker string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			isAdmin, ok := r.Context().Value(ctxkey.IsAdmin).(bool)
			if !ok {
				log.Print("Failed to fetch value from the context")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			if isAdmin {
				h.ServeHTTP(w, r)
				return
			}
			log.Print("HERE IT FAILS")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		})
	}
}
