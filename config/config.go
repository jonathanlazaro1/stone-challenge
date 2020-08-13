package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config defines the config variables gotten from the environment
type Config struct {
	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string
}

// GetConfig generates a ponter to a Config object, filling it with values taken from the environment
func GetConfig() *Config {
	config := &Config{}
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.DBUser = os.Getenv("DB_USER")
	config.DBPass = os.Getenv("DB_PASS")
	config.DBHost = os.Getenv("DB_HOST")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBName = os.Getenv("DB_NAME")

	return config
}
