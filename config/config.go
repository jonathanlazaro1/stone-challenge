package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config defines the config variables gotten from the environment
type Config struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPass     string
	AuthSecret string
}

// Load attempts to load .env var from the projects's root folder
func Load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// GetConfig generates a pointer to a Config object, filling it with values taken from the environment
func GetConfig() *Config {
	config := &Config{}

	config.DBUser = os.Getenv("DB_USER")
	config.DBPass = os.Getenv("DB_PASS")
	config.DBHost = os.Getenv("DB_HOST")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBName = os.Getenv("DB_NAME")
	config.AuthSecret = os.Getenv("AUTH_SECRET")

	return config
}
