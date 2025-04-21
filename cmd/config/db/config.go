package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type DBConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	SSLMode  string
}

func NewPostgresDB(config DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Database,
		config.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	if err := runMigrations(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("error running migrations: %v", err)
	}

	return db, nil
}

func runMigrations(db *sql.DB) error {

	goose.SetDialect("postgres")
	migrationsDir := "./cmd/config/db/migrations"

	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		return fmt.Errorf("migrations directory does not exist: %s", migrationsDir)
	}

	version, err := goose.GetDBVersion(db)
	if err != nil {
		return fmt.Errorf("error getting database version: %v", err)
	}

	if version > 0 {
		if err := goose.Reset(db, migrationsDir); err != nil {
			return fmt.Errorf("error resetting migrations: %v", err)
		}
	}

	if err := goose.Up(db, migrationsDir); err != nil {
		return fmt.Errorf("error running migrations: %v", err)
	}

	return nil
}
