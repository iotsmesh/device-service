package logger

import (
	log "github.com/sirupsen/logrus"
)

// Logger instance for the service
var Logger *log.Logger

// Set the default attributes for logger
func init() {
	Logger = log.New()

	Logger.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	Logger.SetLevel(log.DebugLevel)
}

// Set logger attribute according to the configuration
func Config() {

}
