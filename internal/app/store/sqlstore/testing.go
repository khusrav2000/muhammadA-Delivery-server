package sqlstore

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			log.Println("Clean atble", tables)
			if _, err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				log.Println("CLEANT TABLE ERROR")
				log.Println(err)
				t.Fatal(err)
			}
		}

		db.Close()
	}

}
