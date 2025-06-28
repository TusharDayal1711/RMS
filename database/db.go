package database

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Init(connectionStr string) {
	var err error
	DB, err = sqlx.Connect("postgres", connectionStr)
	if err != nil {
		log.Fatalf("Cannot connect to Postgres: %+v", err)
	}
	fmt.Println("Connected to PostgreSQL...")

	if err := migrateUp(DB); err != nil {
		log.Fatalf("Migration failed: %+v", err)
	}
}

func migrateUp(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://database/migrations", "postgres", driver)
	if err != nil {
		return err
	}

	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return fmt.Errorf("version check failed: %w", err)
	}
	fmt.Printf("Before migration: version=%v dirty=%v\n", version, dirty)

	if err := m.Up(); err != nil && err.Error() != "no change" {
		return err
	}

	version, dirty, _ = m.Version()
	fmt.Printf("After migration: version=%v dirty=%v\n", version, dirty)

	fmt.Println("Migration completed.")
	return nil
}
