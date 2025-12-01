package repository

import "database/sql"

type Repository struct {
	User    *UserRepository
	Content *ContentRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User:    NewUserRepository(db),
		Content: NewContentRepository(db),
	}
}
