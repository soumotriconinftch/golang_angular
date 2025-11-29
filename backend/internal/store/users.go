package store

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UsersStore struct {
	db *sql.DB
}

type Users struct {
	ID       int64    `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password password `json:"-"`
}

type password struct {
	text *string
	hash []byte
}

var (
	ErrorDuplicateEmail    = errors.New("a user with that email already exists")
	ErrorDuplicateUsername = errors.New("a user with that username already exists")
)

func (p *password) Set(pass string) error {
	log.Println("entering hashing")
	cost := bcrypt.DefaultCost
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return err
	}
	p.text = &pass
	p.hash = bytes
	log.Println("hashed successfully")
	return nil
}

func (s *UsersStore) Create(ctx context.Context, users *Users) error {
	query := `
	INSERT INTO users (username, email, password)
	VALUES ($1, $2, $3) RETURNING id
	`
	ctx1, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	log.Println("Executing create user query")
	err := s.db.QueryRowContext(
		ctx1,
		query,
		users.Username,
		users.Email,
		users.Password.hash,
	).Scan(
		&users.ID,
	)

	if err != nil {
		log.Printf("Database error: %v", err)
		errStr := err.Error()
		switch {
		case strings.Contains(errStr, "duplicate"):
			return ErrorDuplicateUsername
		case strings.Contains(errStr, "duplicate"):
			return ErrorDuplicateEmail
		default:
			return err
		}
	}

	log.Printf("User created: %d", users.ID)
	return nil
}

// func (s *UsersStore) GetByID(ctx context.Context, id int64) (*Users, error) {
// 	query := `
// 	SELECT id, username, content
// 	FROM users
// 	WHERE id = $1
// 	`
// 	user := &Users{}
// 	err := s.db.QueryRowContext(ctx, query, id).Scan(
// 		&user.ID,
// 		&user.Username,
// 		&user.Password.hash,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// func (s *UsersStore) GetAll(ctx context.Context) ([]*Users, error) {
// 	query := `
// 	SELECT id, username, content
// 	FROM users
// 	`
// 	rows, err := s.db.QueryContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var users []*Users
// 	for rows.Next() {
// 		var user Users
// 		if err := rows.Scan(&user.ID, &user.Username, &user.Content); err != nil {
// 			return nil, err
// 		}
// 		users = append(users, &user)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }
