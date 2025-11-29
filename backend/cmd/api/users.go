package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/szoumoc/golang_angular/internal/store"
)

type NewUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload NewUserPayload

	log.Println("create user handler entered")

	// Decoding payload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("failed to decode JSON: %v", err)
		http.Error(w, "invalid json format", http.StatusBadRequest)
		return
	}
	log.Println(" decoded into payload")

	// validate thee payload

	if err := Validate.Struct(payload); err != nil {
		log.Printf("Validation failed: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":   "Validation failed",
			"details": err.Error(),
		})
		return
	}
	log.Println("payload validated success")

	user := &store.Users{
		Username: payload.Username,
		Email:    payload.Email,
	}

	// hashed the password

	if err := user.Password.Set(payload.Password); err != nil {
		log.Printf("failed hash password: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Println("Password hashed successfully")

	// create userr

	if err := app.store.Users.Create(r.Context(), user); err != nil {
		log.Printf("Failed to create user in database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("User created successfully with ID: %d", user.ID)

	// s8uccess response

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// func (app *application) getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
// 	users, err := app.store.Users.GetAll(r.Context())
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	if err := json.NewEncoder(w).Encode(users); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }
