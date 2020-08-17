package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

// Config defines the config variables gotten from the environment
type Config struct {
	DBURL         string
	DBHost        string
	DBPort        string
	DBName        string
	DBUser        string
	DBPass        string
	DBSSLMode     string
	AppPort       string
	AppAuthSecret string
}

// Load attempts to load .env var from the projects's root folder
func Load() {
	// If this env var is present, other vars must be present on the env host (e.g Docker, Heroku)
	fetchVarsFromOSEnv := os.Getenv("OS_ENV_VARS")
	if fetchVarsFromOSEnv != "" {
		return
	}

	re := regexp.MustCompile(`^(.*` + "stone-challenge" + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatalf("Error loading .env vars: %v", err)
	}
}

// GetConfig generates a pointer to a Config object, filling it with values taken from the environment
func GetConfig() *Config {
	config := &Config{}

	config.DBURL = os.Getenv("DATABASE_URL")
	config.DBUser = os.Getenv("DB_USER")
	config.DBPass = os.Getenv("DB_PASS")
	config.DBHost = os.Getenv("DB_HOST")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBName = os.Getenv("DB_NAME")
	config.DBSSLMode = os.Getenv("DB_SSL_MODE")
	config.AppAuthSecret = os.Getenv("APP_AUTH_SECRET")
	if config.AppPort = os.Getenv("PORT"); config.AppPort == "" {
		log.Fatalf("Couldn't find App Port to run")
	}

	return config
}
