package main

import (
	"github.com/iotsmesh/device-service/internal/db"
	logger "github.com/iotsmesh/device-service/internal/logger"
)

func main() {
	if res := db.Migrate(); res {
		logger.Logger.Info("Database migration is done!")
	} else {
		logger.Logger.Info("Database migration is failed!")
	}
	logger.Logger.Info("Device Service is started!")
}
