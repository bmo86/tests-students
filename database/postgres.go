package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgressRepository struct {
	db *sql.DB
}

func NewPortgressRepository(url string) (*PostgressRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgressRepository{db}, nil
}
