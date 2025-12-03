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
			r.With(middleware.Authorization("is_admin")).Get("/", userHandler.GetAllUsers)

			// r.Get("/", userHandler.GetAllUsers)

			r.Get("/me", userHandler.GetCurrentUser)

			r.Route("/me/content", func(r chi.Router) {
				r.Post("/", contentHandler.Create)
				r.Get("/", contentHandler.GetAll)
				r.Get("/{id}", contentHandler.GetByID)
			})
		})

		//TODO
		// adminaccesskey := "is_admin"
		// r.Group(func(r chi.Router) {
		// 	r.Use(middleware.Authorization(adminaccesskey))
		// 	r.Get("/", userHandler.GetAllUsers)
		// })
		// r.Use(middleware.Authorization).
		// r.Get("/{id}", userHandler.GetUserDataAdmin)
		// r.Get("/", userHandler.GetAllUsers)

		// With adds inline middlewares for an endpoint handler.
		// With(middlewares ...func(http.Handler) http.Handler) Router
		// r.With(middleware.Authorization("is_admin")).Get("/", userHandler.GetAllUsers)

	})

	return r
}
