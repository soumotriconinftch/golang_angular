package repository

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/szoumoc/golang_angular/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	query := `
	INSERT INTO users (username, email, password)
	VALUES ($1, $2, $3) RETURNING id
	`
	ctx1, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	log.Println("Executing create user query")
	err := r.db.QueryRowContext(
		ctx1,
		query,
		user.Username,
		user.Email,
		user.Password.Hash,
	).Scan(
		&user.ID,
	)

	if err != nil {
		log.Printf("Database error: %v", err)
		errStr := err.Error()
		switch {
		case strings.Contains(errStr, "duplicate"):
			return models.ErrDuplicateUsername
		default:
			return err
		}
	}

	log.Printf("User created: %d", user.ID)
	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	query := `
	SELECT id, username, email
	FROM users
	WHERE id = $1
	`
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
	SELECT id, username, email, password
	FROM users
	WHERE email = $1
	`
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password.Hash,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
