package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Users interface {
		// GetByID(context.Context, int64) (*Users, error)
		Create(context.Context, *Users) error
		// GetAll(context.Context) ([]*Users, error)
		// Delete(context.Context, int64) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Users: &UsersStore{db},
	}
}
