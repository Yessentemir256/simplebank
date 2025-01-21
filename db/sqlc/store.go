package db

import (
	"database/sql"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Querier
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
