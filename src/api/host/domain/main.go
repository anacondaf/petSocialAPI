package domain

import (
	"database/sql"
	"errors"
	"fmt"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunDBMigration(db *sql.DB, MigrationUrl string) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		MigrationUrl,
		"pet-social",
		driver,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

func UseSqlc(logger *zap.Logger) *Queries {
	const ConnectionString = "root:secret@tcp(localhost:3306)/pet-social?multiStatements=true"

	db, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("Connect to mysql success")

	err = RunDBMigration(db, "file://src/api/host/migrations")
	if err != nil {
		logger.Fatal(fmt.Sprintf("Error when run db migration: %v", err))
	}

	logger.Info("Migration successfully")

	return New(db)
}
