package main

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/szoumoc/golang_angular/internal/auth"
)

type contextKey string

const userIDKey contextKey = "user_id"

func (a *application) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		accessCookie, err := r.Cookie("accessToken")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		tok, err := auth.ValidateToken(accessCookie.Value)
		if err != nil {
			// Access & refresh
			// JWT is missing, JWT is invalid text, Not generated with secret key, jwt expired
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		if tok.Valid {
			next.ServeHTTP(w, r)
			return
		}

		// 1. Validate the access token
		// 2. If yes then serve with the write and reader
		// 3. if expired and valid, then validate the refresh token
		// 4. if refresh token is invalid then write http error for unauthorised
		// 5. if refresh token is valid, then generate new access token
		// 6. serve with write and reader

		// if err == nil {
		// 	tok, err := auth.ValidateToken(accessCookie.Value)
		// 	if err == nil && tok.Valid {
		// 		next.ServeHTTP(w, r)
		// 		return
		// 	}
		// }

		refreshCookie, err := r.Cookie("refreshToken")
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		refreshTok, err := auth.ValidateRefresh(refreshCookie.Value)
		if err != nil || !refreshTok.Valid {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		claims, ok := refreshTok.Claims.(jwt.MapClaims)
		if !ok {
			panic("invalid claims type")
		}

		raw, ok := claims["user_id"]
		if !ok {
			panic("user_id missing")
		}
		uid := raw.(int64)
		newAccess, err := auth.GenerateAccessToken(uid)
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

		next.ServeHTTP(w, r)
	})
}

var origins = "localhost, google.com, hello"

func (app *application) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origins)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Username")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
