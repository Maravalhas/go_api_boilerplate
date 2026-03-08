package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Environment string

const (
	Development Environment = "development"
	Debug       Environment = "debug"
	Test        Environment = "test"
	Production  Environment = "production"
)

type Configvar string

const (
	DatabaseDsn     Configvar = "DATABASE_DSN"
	OauthTokenUrl   Configvar = "OAUTH_TOKEN_URL"
	OauthJWKSUrl    Configvar = "OAUTH_JWKS_URL"
	OauthIssuers    Configvar = "OAUTH_ISSUERS"
	ClientId        Configvar = "CLIENT_ID"
	ClientSecret    Configvar = "CLIENT_SECRET"
	RedirectUri     Configvar = "REDIRECT_URI"
	Origin          Configvar = "ORIGIN"
	Port            Configvar = "PORT"
	Env             Configvar = "ENV"
	IdTokenTTL      Configvar = "ID_TOKEN_TTL"
	RefreshTokenTTL Configvar = "REFRESH_TOKEN_TTL"
)

type Config struct {
	// GORM
	DatabaseDsn string

	// OAUTH
	OauthTokenUrl   string
	OauthJWKSUrl    string
	OauthIssuers    []string
	ClientId        string
	ClientSecret    string
	RedirectUri     string
	IdTokenTTL      int
	RefreshTokenTTL int

	// APP
	Origin string
	Port   string

	Env Environment
}

func IsProduction() bool {
	return Current.Env == Production
}

func IsDebug() bool {
	return Current.Env == Debug || Current.Env == Test
}

func IsDevelopment() bool {
	return Current.Env == Development
}

func IsTest() bool {
	return Current.Env == Test
}

var Current *Config

func LoadConfig() {
	if os.Getenv("ENV") != string(Production) {
		if err := godotenv.Load(); err != nil {
			panic("Error loading .env file")
		}
	}

	id_token_ttl, err := strconv.Atoi(getEnvOrDefault(IdTokenTTL, "3600"))
	if err != nil {
		panic("Invalid value for ID_TOKEN_TTL")
	}

	refresh_token_ttl, err := strconv.Atoi(getEnvOrDefault(RefreshTokenTTL, "86400"))
	if err != nil {
		panic("Invalid value for REFRESH_TOKEN_TTL")
	}

	Current = &Config{
		DatabaseDsn: getEnvOrPanic(DatabaseDsn),

		OauthTokenUrl:   getEnvOrPanic(OauthTokenUrl),
		OauthJWKSUrl:    getEnvOrPanic(OauthJWKSUrl),
		OauthIssuers:    strings.Split(getEnvOrPanic(OauthIssuers), ","),
		ClientId:        getEnvOrPanic(ClientId),
		ClientSecret:    getEnvOrPanic(ClientSecret),
		RedirectUri:     getEnvOrPanic(RedirectUri),
		IdTokenTTL:      id_token_ttl,
		RefreshTokenTTL: refresh_token_ttl,

		Origin: getEnvOrPanic(Origin),
		Port:   getEnvOrDefault(Port, "8080"),

		Env: Environment(getEnvOrDefault(Env, string(Development))),
	}
}

func getEnvOrDefault(key Configvar, defaultValue string) string {
	value, exists := os.LookupEnv(string(key))
	if !exists {
		return defaultValue
	}
	return value
}

func getEnvOrPanic(key Configvar) string {
	value, exists := os.LookupEnv(string(key))
	if !exists {
		panic("Environment variable " + string(key) + " is required.")
	}
	return value
}
