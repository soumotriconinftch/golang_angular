package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/szoumoc/golang_angular/internal/models"
)

type ContentRepository struct {
	db *sql.DB
}

func NewContentRepository(db *sql.DB) *ContentRepository {
	return &ContentRepository{db: db}
}

func (r *ContentRepository) Create(ctx context.Context, content *models.Content) error {
	query := `
	INSERT INTO content (body, title, user_id)
	VALUES ($1, $2, $3) RETURNING id, created_at
	`
	ctx1, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	err := r.db.QueryRowContext(ctx1,
		query,
		content.Body,
		content.Title,
		content.UserID,
	).Scan(
		&content.ID,
		&content.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
