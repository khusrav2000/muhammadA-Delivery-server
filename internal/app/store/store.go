package store

import (
	"database/sql"
	"log"

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
	log.Println("ERROR CONECTION1!")
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
