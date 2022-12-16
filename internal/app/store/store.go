package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	config *Config
	db     *sql.DB
	repo   *Repo
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *Repo {
	if s.repo != nil {
		return s.repo
	}
	s.repo = &Repo{
		store: s,
	}

	return s.repo
}
