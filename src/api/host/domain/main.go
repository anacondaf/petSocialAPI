package domain

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunDBMigration(db *sql.DB, MIGRATION_URL string) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	fmt.Println(driver.Version())

	m, err := migrate.NewWithDatabaseInstance(
		MIGRATION_URL,
		"pet-social",
		driver,
	)
	if err != nil {
		return err
	}

	return m.Up()
}

func UseSqlc(logger *zap.Logger) *Queries {
	const ConnectionString = "root:secret@tcp(localhost:3306)/pet-social?multiStatements=true"

	db, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("Connect to mysql success")

	err = RunDBMigration(db, "file://../api/migrations")
	if err != nil {
		logger.Fatal(fmt.Sprintf("Error when run db migration: %v", err))
	}

	logger.Info("Migration successfully")

	return New(db)
}
