package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/szoumoc/golang_angular/internal/auth"
	"github.com/szoumoc/golang_angular/internal/store"
)

type NewUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("========== CREATE USER REQUEST STARTED ==========")
	var payload NewUserPayload

	log.Println("Step 1: Decoding JSON payload from request body")
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("Step 1 FAILED: failed to decode JSON: %v", err)
		http.Error(w, "invalid json format", http.StatusBadRequest)
		return
	}
	log.Printf("Step 1 SUCCESS: JSON decoded - Username: %s, Email: %s", payload.Username, payload.Email)

	log.Println("Step 2: Validating user payload")
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

	log.Println("Step 3: Creating user object")
	user := &store.Users{
		Username: payload.Username,
		Email:    payload.Email,
	}
	log.Printf("Step 3 SUCCESS: User object created for username: %s", payload.Username)

	log.Println("Step 4: Hashing password")
	if err := user.Password.Set(payload.Password); err != nil {
		log.Printf("Step 4 FAILED: failed to hash password: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Println("Step 4 SUCCESS: Password hashed successfully")

	log.Println("Step 5: Saving user to database")
	if err := app.store.Users.Create(r.Context(), user); err != nil {
		log.Printf("Step 5 FAILED: Failed to create user in database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Step 5 SUCCESS: User created successfully with ID: %d", user.ID)

	log.Printf("Step 6: Generating JWT token for user ID: %d", user.ID)
	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		log.Printf("Step 6 FAILED: Failed to generate token: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Println("Step 6 SUCCESS: Token generated successfully")

	log.Println("Step 7: Sending success response")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"user":  user,
		"token": token,
	}); err != nil {
		log.Printf("Step 7 FAILED: Failed to encode response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Printf("Step 7 SUCCESS: User registration completed for: %s", user.Username)
	log.Println("========== CREATE USER REQUEST COMPLETED ==========")
}

func (app *application) getCurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(userIDKey).(int64)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := app.store.Users.GetByID(r.Context(), userID)
	if err != nil {
		log.Printf("Failed to get user by ID: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
