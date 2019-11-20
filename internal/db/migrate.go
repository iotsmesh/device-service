package db

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	logger "github.com/iotsmesh/device-service/internal/logger"
)

// Migrate Database
func Migrate() bool {

	// TODO: Use configuration for DB URL
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:iotsmesh@localhost:5432/devicesvc?sslmode=disable")

	if err != nil {
		logger.Logger.Fatal(err)
		return false
	}

	if err := m.Up(); err != nil {
		logger.Logger.Fatal(err)
		return false
	}

	return true
}
