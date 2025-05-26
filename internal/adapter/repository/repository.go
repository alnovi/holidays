package repository

import (
	"github.com/alnovi/holidays/pkg/database/sqlite"
)

type Repository struct {
	db *sqlite.Client
}

func NewRepository(db *sqlite.Client) *Repository {
	return &Repository{db: db}
}
