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
	ctx1, cancel := context.WithTimeout(ctx, time.Second*5)
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

func (r *ContentRepository) GetAll(ctx context.Context, userID int64) ([]*models.Content, error) {
	query := `
	SELECT id, user_id, title, body, created_at
	FROM content
	WHERE user_id = $1
	ORDER BY created_at DESC
	`
	ctx1, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	rows, err := r.db.QueryContext(ctx1, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contents []*models.Content
	for rows.Next() {
		content := &models.Content{}
		if err := rows.Scan(
			&content.ID,
			&content.UserID,
			&content.Title,
			&content.Body,
			&content.CreatedAt,
		); err != nil {
			return nil, err
		}
		contents = append(contents, content)
	}
	return contents, nil
}

func (r *ContentRepository) GetByID(ctx context.Context, id int64) (*models.Content, error) {
	query := `
	SELECT id, user_id, title, body, created_at
	FROM content
	WHERE id = $1
	`
	ctx1, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	content := &models.Content{}
	err := r.db.QueryRowContext(ctx1, query, id).Scan(
		&content.ID,
		&content.UserID,
		&content.Title,
		&content.Body,
		&content.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return content, nil
}
