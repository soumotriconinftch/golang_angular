package main

import (
	"net/http"
)

func (app *application) checkUsernameMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("X-Username")
		if username != "szoumo" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (app *application) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from Angular dev server (any localhost port)
		origin := r.Header.Get("Origin")
		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Username")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
