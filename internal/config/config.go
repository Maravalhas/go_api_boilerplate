package config

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	DatabaseDsn   = "DATABASE_DSN"
	OauthTokenUrl = "OAUTH_TOKEN_URL"
	OauthIdUrl    = "OAUTH_ID_URL"
	ClientId      = "CLIENT_ID"
	ClientSecret  = "CLIENT_SECRET"
	RedirectUri   = "REDIRECT_URI"
	Origin        = "ORIGIN"
	Port          = "PORT"
)

type Config struct {
	// GORM
	DatabaseDsn string

	// OAUTH
	OauthTokenUrl string
	OauthIdUrl    string
	ClientId      string
	ClientSecret  string
	RedirectUri   string

	// APP
	Origin string
	Port   string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	return &Config{
		DatabaseDsn:   getEnvOrPanic(DatabaseDsn),
		OauthTokenUrl: getEnvOrPanic(OauthTokenUrl),
		OauthIdUrl:    getEnvOrPanic(OauthIdUrl),
		ClientId:      getEnvOrPanic(ClientId),
		ClientSecret:  getEnvOrPanic(ClientSecret),
		RedirectUri:   getEnvOrPanic(RedirectUri),
		Origin:        getEnvOrPanic(Origin),
		Port:          getEnvOrDefault(Port, "8080"),
	}
}

func getEnvOrDefault(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func getEnvOrPanic(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic("Environment variable " + key + " is required.")
	}
	return value
}
