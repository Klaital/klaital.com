package utilities

import (
	"database/sql"
	"embed"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func ConnectAndMigratePostgres(connectFmt string, targetDbName string, migrationsFS embed.FS) (*sql.DB, error) {
	// Attempt to connect to the project DB
	projectConnStr := fmt.Sprintf(connectFmt, targetDbName)
	projectDb, err := sql.Open("postgres", projectConnStr)
	if err != nil {
		return nil, fmt.Errorf("connecting to project db: %w", err)
	}
	err = projectDb.Ping()
	if err == nil {
		slog.Debug("Successfully connected to exiting project DB", "conn", projectConnStr, "db", targetDbName)
		err = MigrateSql(projectDb, migrationsFS)
		return projectDb, err
	}

	// If we failed to connect to the project DB, try
	// connecting to the system DB and creating the DB
	// TODO: inspect the error to see if this is worth trying
	slog.Debug("Failed to connect to project DB", "err", err)

	// connect to master DB and ensure the DB exists
	rootDb, err := sql.Open("postgres", fmt.Sprintf(connectFmt, "postgres"))
	defer rootDb.Close()
	if err != nil {
		return nil, fmt.Errorf("connect to root db: %w", err)
	} else {
		slog.Debug("Connected to root DB")
	}
	_, err = rootDb.Exec(fmt.Sprintf(`CREATE DATABASE %s`, targetDbName))
	if err != nil {
		return nil, fmt.Errorf("creating project db: %w", err)
	} else {
		slog.Debug("Created project db", "db", targetDbName)
	}

	// connect to the target DB and run migrations
	projectDb, err = sql.Open("postgres", projectConnStr)
	if err != nil {
		slog.Error("Still unable to connect to project DB", "err", err)
		return nil, err
	}
	err = MigrateSql(projectDb, migrationsFS)
	return projectDb, err
}

func MigrateSql(db *sql.DB, migrationsFS embed.FS) error {
	goose.SetBaseFS(migrationsFS)
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("set dialect to postgres: %w", err)
	}
	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("migrating db schema: %w", err)
	}
	// Success!
	return nil
}
