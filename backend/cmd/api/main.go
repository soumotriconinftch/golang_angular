package main

import (
	"log"
	"net/http"
	"time"

	cfg "github.com/szoumoc/golang_angular/internal/config"
	"github.com/szoumoc/golang_angular/internal/database"
	"github.com/szoumoc/golang_angular/internal/env"
	"github.com/szoumoc/golang_angular/internal/handlers"
	"github.com/szoumoc/golang_angular/internal/repository"
)

func main() {
	config := cfg.Config{
		Server: cfg.ServerConfig{
			Address: env.GetString("ADDR", ":9000"),
		},
		Database: cfg.DatabaseConfig{
			Address:      env.GetString("DB_ADDR", ""),
			MaxOpenConns: 30,
			MaxIdleConns: 30,
			MaxIdleTime:  "15m",
		},
		Secret: env.GetString("ACCESS_SECRET","abc"),
	}

	db, err := database.New(
		config.Database.Address,
		config.Database.MaxOpenConns,
		config.Database.MaxIdleConns,
		config.Database.MaxIdleTime,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("DB connected!")

	repo := repository.NewRepository(db)
	router := handlers.SetupRoutes(repo)

	srv := &http.Server{
		Addr:         config.Server.Address,
		Handler:      router,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started %s", config.Server.Address)
	log.Fatal(srv.ListenAndServe())
}
