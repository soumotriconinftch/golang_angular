package handlers

import (
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/szoumoc/golang_angular/internal/middleware"
	"github.com/szoumoc/golang_angular/internal/repository"
)

func SetupRoutes(repo *repository.Repository) *chi.Mux {
	userHandler := NewUserHandler(repo)
	contentHandler := NewContentHandler(repo)

	r := chi.NewRouter()
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(middleware.CORS)
	r.Use(chimiddleware.Timeout(30 * time.Second))

	r.Route("/user", func(r chi.Router) {
		r.Post("/sign-up", userHandler.SignUp)
		r.Post("/sign-in", userHandler.SignIn)
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth)
			r.Get("/me", userHandler.GetCurrentUser)
			r.Route("/me/content", func(r chi.Router) {
				r.Post("/", contentHandler.Create)
			})
		})
	})

	return r
}
