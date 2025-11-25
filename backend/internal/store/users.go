package store

import (
	"context"
	"database/sql"
)

type UsersStore struct {
	db *sql.DB
}

type Users struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (s *UsersStore) Create(ctx context.Context, users *Users) error {
	query := `
	INSERT INTO users (name, content)
	VALUES ($1, $2) RETURNING id
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		users.Name,
		users.Content,
	).Scan(
		&users.ID,
	)
	if err != nil {
		return err
	}
	return nil
}
