package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	User interface {
		GetByID(context.Context, int64) (*User, error)
		Create(context.Context) error
		Delete(context.Context, int64) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Users: &UsersStore{db},
	}
}
