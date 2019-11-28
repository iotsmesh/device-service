package db

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/iotsmesh/device-service/internal/config"
	logger "github.com/iotsmesh/device-service/internal/logger"
)

// Migrate Database
func Migrate() bool {
	dbURL := databaseURL()

	m, err := migrate.New("file://migrations", dbURL)

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

func databaseURL() string {
	dbURL := "postgres://" + config.Config.Database.UserName + ":" +
		config.Config.Database.Password + "@" + config.Config.Database.HostName + ":" +
		config.Config.Database.Port + "/" + config.Config.Database.DBName + "?sslmode=disable"

	return dbURL
}
