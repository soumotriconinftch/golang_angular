package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/szoumoc/golang_angular/internal/auth"
	"github.com/szoumoc/golang_angular/internal/ctxkey"
	"github.com/szoumoc/golang_angular/internal/models"
	"github.com/szoumoc/golang_angular/internal/repository"
	"github.com/szoumoc/golang_angular/internal/validator"
)

type UserHandler struct {
	repo *repository.Repository
}

func NewUserHandler(repo *repository.Repository) *UserHandler {
	return &UserHandler{repo: repo}
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignUpPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	log.Println("========== CREATE USER REQUEST STARTED ==========")
	var payload SignUpPayload

	log.Println("Step 1: Decoding JSON payload from request body")
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("Step 1 FAILED: failed to decode JSON: %v", err)
		http.Error(w, "invalid json format", http.StatusBadRequest)
		return
	}
	log.Printf("Step 1 SUCCESS: JSON decoded - Username: %s, Email: %s", payload.Username, payload.Email)

	log.Println("Step 2: Validating user payload")
	if err := validator.Validate.Struct(payload); err != nil {
		log.Printf("Step 2 FAILED: Validation failed: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Step 2 SUCCESS: Payload validation passed")

	log.Println("Step 3: Creating user object")
	user := &models.User{
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
	if err := h.repo.User.Create(r.Context(), user); err != nil {
		log.Printf("Step 5 FAILED: Failed to create user in database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Step 5 SUCCESS: User created successfully with ID: %d", user.ID)

	log.Printf("Step 6: Generating JWT token for user ID: %d", user.ID)
	accessToken, _ := auth.GenerateAccessToken(user.ID, user.IsAdmin)
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
	log.Printf("Step 6 SUCCESS: Token generated successfully")

	log.Println("Step 7: Sending success response with token")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
	log.Printf("Step 7 SUCCESS: User registration completed for: %s", user.Username)
	log.Println("CREATE USER REQUEST COMPLETED")
}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	log.Println("LOGIN REQUEST STARTED")
	var payload LoginPayload

	log.Println("Step 1: Decoding JSON payload from request body")
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("Step 1 FAILED: failed to decode JSON: %v", err)
		http.Error(w, "invalid json format", http.StatusBadRequest)
		return
	}
	log.Printf("Step 1 SUCCESS: JSON decoded - Email: %s", payload.Email)
	payload.Email = strings.TrimSpace(payload.Email)
	log.Println("Step 2: Validating login payload")
	if err := validator.Validate.Struct(payload); err != nil {
		log.Printf("Step 2 FAILED: Validation failed: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Step 2 SUCCESS: Payload validation passed")

	log.Printf("Step 3: Fetching user from database by email: %s", payload.Email)
	user, err := h.repo.User.GetByEmail(r.Context(), payload.Email)
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
	accessToken, _ := auth.GenerateAccessToken(user.ID, user.IsAdmin)
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
	json.NewEncoder(w).Encode(user)
	log.Printf("Step 6 SUCCESS: Login completed for user: %s", user.Username)
	log.Println("LOGIN REQUEST COMPLETED")
}

func (h *UserHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	// log.Printf("Value: %s", r.Context().Value("user_id"))
	// user_ID := "user_id"
	// log.Printf("Value: %s", r.Context().Value(user_ID))
	// log.Printf("Value: %s", user_ID)
	// type hello string
	// var key hello = "user_id"

	log.Printf("Value: %s", ctxkey.UserID)
	log.Printf("Value: %s", r.Context().Value(ctxkey.UserID))
	userID, ok := r.Context().Value(ctxkey.UserID).(int64)
	if !ok {
		log.Print("Failed to fetch value from the context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.repo.User.GetByID(r.Context(), userID)
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

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.repo.User.GetAll(r.Context())
	if err != nil {
		log.Printf("Failed to get all users: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
