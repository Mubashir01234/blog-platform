package config

import (
	"github.com/spf13/viper" // Importing the viper package for configuration handling
)

type Config struct {
	MongoURL   string `mapstructure:"MONGO_URL"`   // Configuration field for MongoDB URL
	ServerPort string `mapstructure:"SERVER_PORT"` // Configuration field for server port
	JwtSecret  string `mapstructure:"JWT_SECRET"`  // Configuration field for JWT secret
}

var Cfg Config // Global variable to hold the configuration

func LoadConfig() error {
	viper.AddConfigPath("./")   // Adding the configuration path
	viper.SetConfigName(".env") // Setting the configuration file name
	viper.SetConfigType("env")  // Setting the configuration file type
	err := viper.ReadInConfig() // Reading the configuration file
	viper.AutomaticEnv()        // Enabling automatic environment variable binding
	if err != nil {
		return err // Returning the error if there is any issue reading the configuration
	}
	err = viper.Unmarshal(&Cfg) // Unmarshaling the configuration into the Cfg variable
	return err
}
