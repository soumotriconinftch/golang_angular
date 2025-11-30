package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/szoumoc/golang_angular/internal/store"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
	db   dbConfig
}
type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(app.corsMiddleware)

	r.Use(middleware.Timeout(30 * time.Second))
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})
	r.Route("/user", func(r chi.Router) {
		r.Post("/sign-up", app.createUserHandler)
		r.Post("/sign-in", app.loginHandler)
		r.Group(func(r chi.Router) {
			r.Use(app.AuthTokenMiddleware)
			r.Get("/me", app.getCurrentUserHandler)
		})
	})
	return r
}

func (app *application) run(mux *chi.Mux) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started %s", app.config.addr)
	return srv.ListenAndServe()
}
