package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func setupRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/test/{name}", Validator)
	return r
}

func TestValidator_MissingName(t *testing.T) {
	req := httptest.NewRequest("POST", "/test/", nil)
	rr := httptest.NewRecorder()

	router := chi.NewRouter()
	router.Post("/test/", Validator)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestValidator_InvalidJSON(t *testing.T) {
	body := bytes.NewBuffer([]byte(`{invalid json`))
	req := httptest.NewRequest("POST", "/test/soumo", body)
	rr := httptest.NewRecorder()

	router := setupRouter()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestValidator_ValidationError(t *testing.T) {
	body := bytes.NewBuffer([]byte(`{"name":"ab","age":0}`))
	req := httptest.NewRequest("POST", "/test/soumo", body)
	rr := httptest.NewRecorder()

	router := setupRouter()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestValidator_Success(t *testing.T) {
	sample := CreateUserPayload{
		Name: "abcd",
		Age:  20,
	}
	data, _ := json.Marshal(sample)

	req := httptest.NewRequest("POST", "/test/soumo", bytes.NewBuffer(data))
	rr := httptest.NewRecorder()

	router := setupRouter()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, rr.Code)
	}
}
