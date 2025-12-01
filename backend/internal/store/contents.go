package store

import (
	"context"
	"database/sql"
	"time"
)

type ContentsStore struct {
	db *sql.DB
}

type Contents struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	Title     string `json:"username"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	Users     Users  `json:"user"`
}

func (s *ContentsStore) Create(ctx context.Context, contents *Contents) error {
	query := `
	INSERT INTO content (body, title, user_id)
	VALUES ($1, $2, $3) RETURNING id, created_at
	`
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()
	err := s.db.QueryRowContext(ctx,
		query,
		contents.Body,
		contents.Title,
		contents.UserID,
	).Scan(
		&contents.ID,
		&contents.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
