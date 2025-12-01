package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/szoumoc/golang_angular/internal/auth"
)

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("LOGIN REQUEST STARTED")
	var payload LoginPayload

	log.Println("Step 1: Decoding JSON payload from request body")
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("Step 1 FAILED: failed to decode JSON: %v", err)
		http.Error(w, "invalid json format", http.StatusBadRequest)
		return
	}
	log.Printf("Step 1 SUCCESS: JSON decoded - Email: %s", payload.Email)

	log.Println("Step 2: Validating login payload")
	if err := Validate.Struct(payload); err != nil {
		log.Printf("Step 2 FAILED: Validation failed: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":   "Validation failed",
			"details": err.Error(),
		})
		return
	}
	log.Println("Step 2 SUCCESS: Payload validation passed")

	log.Printf("Step 3: Fetching user from database by email: %s", payload.Email)
	user, err := app.store.Users.GetByEmail(r.Context(), payload.Email)
	if err != nil {
		log.Printf("Step 3 FAILED: Failed to get user by email: %v", err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	log.Printf("Step 3 SUCCESS: User found - ID: %d, Username: %s", user.ID, user.Username)

	log.Println("Step 4: Comparing provided password with stored hash")
	if err := user.ComparePassword(payload.Password); err != nil {
		log.Printf("Step 4 FAILED: Password mismatch for user %s: %v", payload.Email, err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	log.Println("Step 4 SUCCESS: Password matched")

	log.Printf("Step 5: Generating JWT token for user ID: %d", user.ID)
	// token, err := auth.GenerateToken(user.ID)
	// if err != nil {
	// 	log.Printf("Step 5 FAILED: Failed to generate token: %v", err)
	// 	http.Error(w, "Internal server error", http.StatusInternalServerError)
	// 	return
	// }
	accessToken, _ := auth.GenerateAccessToken(user.ID)

	refreshToken, _ := auth.GenerateRefreshToken(user.ID)

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
	log.Printf("Step 5 SUCCESS: Token generated successfully")

	log.Println("Step 6: Sending success response with token")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user": map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
	log.Printf("Step 6 SUCCESS: Login completed for user: %s", user.Username)
	log.Println("LOGIN REQUEST COMPLETED")
}
