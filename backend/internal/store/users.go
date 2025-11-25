package store

import (
	"context"
	"database/sql"
)

type UsersStore struct {
	db *sql.DB
}

type Users struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Content  string `json:"content"`
}

func (s *UsersStore) Create(ctx context.Context, users *Users) error {
	query := `
	INSERT INTO users (username, content)
	VALUES ($1, $2) RETURNING id
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		users.Username,
		users.Content,
	).Scan(
		&users.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *UsersStore) GetByID(ctx context.Context, id int64) (*Users, error) {
	query := `
	SELECT id, username, content
	FROM users
	WHERE id = $1
	`
	user := &Users{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Content,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UsersStore) GetAll(ctx context.Context) ([]*Users, error) {
	query := `
	SELECT id, username, content
	FROM users
	`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*Users
	for rows.Next() {
		var user Users
		if err := rows.Scan(&user.ID, &user.Username, &user.Content); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
