package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

/*
	The Config struct is a simple struct that holds the configuration values for the application.
	The initConfig function is a constructor function that returns a new Config struct.
	The getEnv function is a helper function that gets the value of an environment variable or returns a fallback value.
*/
type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPass     string
	DBAddr     string
	DBName     string
}

var Envs = initConfig()

// NewConfig is a constructor function that returns a new Config struct
func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port: getEnv("PORT", "8080"),
		DBUser: getEnv("DB_USER", "root"),
		DBPass: getEnv("DB_PASS", "Alexa4ndru1234"),
		DBAddr: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"),getEnv("DB_PORT", "3306")),
		DBName: getEnv("DB_NAME", "ecomserver"),
	}
}

// getEnv is a helper function to get the value of an environment variable or return a fallback value
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}