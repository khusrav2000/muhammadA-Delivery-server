package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	log.Println(s.config.DatabaseURL)
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		log.Println("ERROR CONECTION!")
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	log.Println("Connect with DB!")
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}
