package main

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var accessSecret = []byte("your-access-token-secret")
var refreshSecret = []byte("your-refresh-token-secret")

type user struct {
	ID       int
	Username string
	Password string
}

var users = []user{
	{ID: 1, Username: "user", Password: "password"},
}

// func abc() {
// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/login", loginnHandler)
// 	mux.HandleFunc("/protected", protectedHandler)
// 	mux.HandleFunc("/refresh", refreshHandler)

// 	http.ListenAndServe(":3000", mux)
// }

func loginnHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	var u *user
	for _, x := range users {
		if x.Username == username && x.Password == password {
			u = &x
			break
		}
	}
	if u == nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	accessToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": u.ID,
		"exp":    time.Now().Add(15 * time.Minute).Unix(),
	}).SignedString(accessSecret)

	refreshToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": u.ID,
		"exp":    time.Now().Add(7 * 24 * time.Hour).Unix(),
	}).SignedString(refreshSecret)

	http.SetCookie(w, &http.Cookie{
		Name:     "accessToken",
		Value:    accessToken,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   15 * 60,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   7 * 24 * 60 * 60,
	})

	w.Write([]byte(`{"message":"Login successful"}`))
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("accessToken")
	if err != nil {
		http.Error(w, "No access token", http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(c.Value, func(t *jwt.Token) (interface{}, error) {
		return accessSecret, nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "Invalid or expired access token", http.StatusUnauthorized)
		return
	}

	w.Write([]byte(`{"message":"Access granted"}`))
}

func refreshHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("refreshToken")
	if err != nil {
		http.Error(w, "No refresh token", http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(c.Value, func(t *jwt.Token) (interface{}, error) {
		return refreshSecret, nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "Invalid or expired refresh token", http.StatusUnauthorized)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	uid := claims["userId"]

	newAccess, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": uid,
		"exp":    time.Now().Add(15 * time.Minute).Unix(),
	}).SignedString(accessSecret)

	http.SetCookie(w, &http.Cookie{
		Name:     "accessToken",
		Value:    newAccess,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   15 * 60,
	})

	w.Write([]byte(`{"message":"Access token refreshed"}`))
}
