// Package db provides simple database abstractions for common operations.
package db

import (
	"database/sql"
	"fmt"
	"time"
)

// NewPostgres returns a Postgres implementation given an postgres URL.
func NewPostgres(postgresURL string) Postgres {
	postgresDB, _ := sql.Open("postgres", postgresURL)

	return Postgres{
		db: postgresDB,
	}
}

// Postgres contains a pointer of sql.DB and represents the database
// abrastraction for Postgresql-compatible databases.
type Postgres struct {
	db *sql.DB
}

// CheckConnection runs a simple `SELECT 1` on the Postgres database
// connection instantied and returns an error if some.
func (pq Postgres) CheckConnection() error {
	if _, err := pq.db.Exec("SELECT 1;"); err != nil {
		fmt.Printf("%s - error querying database: %s\n", time.Now().Format(time.ANSIC), err.Error())
		return err
	}
	return nil
}
