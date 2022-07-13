package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	fmt.Println("ERROR CONECTION1!")
	fmt.Println(s.config.DatabaseURL)
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		fmt.Println("ERROR CONECTION!")
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}