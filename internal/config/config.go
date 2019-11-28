package config

import (
	"github.com/spf13/viper"

	logger "github.com/iotsmesh/device-service/internal/logger"
)

// Config provides all the configurations of the applications
var Config Configurations

// Configurations contains all the configurations
type Configurations struct {
	Database DatabaseConfigurations
}

// DatabaseConfigurations contains Database Configurations
type DatabaseConfigurations struct {
	HostName string
	Port     string
	DBName   string
	UserName string
	Password string
}

// LoadConfigurations will load all the configurations from config file
func LoadConfigurations() {
	viper.SetConfigName("config-dev")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		logger.Logger.Errorf("Error reading config file: %s", err)
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		logger.Logger.Errorf("Unable to decode into struct, %v", err)
	}
}
