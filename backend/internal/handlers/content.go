package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/szoumoc/golang_angular/internal/ctxkey"
	"github.com/szoumoc/golang_angular/internal/models"
	"github.com/szoumoc/golang_angular/internal/repository"
	"github.com/szoumoc/golang_angular/internal/validator"
)

type ContentHandler struct {
	repo *repository.Repository
}

func NewContentHandler(repo *repository.Repository) *ContentHandler {
	return &ContentHandler{repo: repo}
}

type CreateContentPayload struct {
	Title string `json:"title" validate:"required,max=255"`
	Body  string `json:"body" validate:"required"`
}

func (h *ContentHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("========== CREATE CONTENT REQUEST STARTED ==========")
	var payload CreateContentPayload

	log.Println("Step 1: Decoding JSON payload from request body")
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("Step 1 FAILED: failed to decode JSON: %v", err)
		http.Error(w, "invalid json format", http.StatusBadRequest)
		return
	}
	log.Printf("Step 1 SUCCESS: JSON decoded - Title: %s", payload.Title)

	log.Println("Step 2: Validating content payload")
	if err := validator.Validate.Struct(payload); err != nil {
		log.Printf("Step 2 FAILED: Validation failed: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Step 2 SUCCESS: Payload validation passed")

	log.Println("Step 3: Extracting user ID from context")
	userID, ok := r.Context().Value(ctxkey.UserID).(int64)
	if !ok {
		log.Println("Step 3 FAILED: User ID not found in context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	log.Printf("Step 3 SUCCESS: User ID extracted: %d", userID)

	log.Println("Step 4: Creating content object")
	content := &models.Content{
		UserID: userID,
		Title:  payload.Title,
		Body:   payload.Body,
	}
	log.Printf("Step 4 SUCCESS: Content object created for user: %d", userID)

	log.Println("Step 5: Saving content to database")
	if err := h.repo.Content.Create(r.Context(), content); err != nil {
		log.Printf("Step 5 FAILED: Failed to create content in database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Step 5 SUCCESS: Content created successfully with ID: %d", content.ID)

	log.Println("Step 6: Sending success response")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(content); err != nil {
		log.Printf("Step 6 FAILED: Failed to encode response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Printf("Step 6 SUCCESS: Content creation completed for: %s", content.Title)
	log.Println("CREATE CONTENT REQUEST COMPLETED")
}
