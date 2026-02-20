// Package gameconfig provides functionality to load game configuration from YAML files and environment variables.
package gameconfig

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// GameConfig holds all application configuration grouped by concern.
type GameConfig struct {
	Title    string `yaml:"GAME_TITLE"`
	Host     string `yaml:"GAME_HOST"`
	Port     int    `yaml:"GAME_PORT"`
	Database DatabaseConfig
	JWT      JWTConfig
	Email    EmailConfig
	Log      LogConfig
}

// DatabaseConfig holds PostgreSQL connection settings.
type DatabaseConfig struct {
	Host     string `yaml:"DATABASE_HOST"`
	Port     int    `yaml:"DATABASE_PORT"`
	Name     string `yaml:"DATABASE_NAME"`
	User     string `yaml:"DATABASE_USER"`
	Password string `yaml:"DATABASE_PASSWORD"`
}

// JWTConfig holds JWT authentication settings.
type JWTConfig struct {
	Secret     string `yaml:"JWT_SECRET"`
	Expiration string `yaml:"JWT_EXPIRATION"`
}

// EmailConfig holds SMTP e-mail settings.
type EmailConfig struct {
	Host     string `yaml:"EMAIL_HOST"`
	Port     int    `yaml:"EMAIL_PORT"`
	Username string `yaml:"EMAIL_USERNAME"`
	Password string `yaml:"EMAIL_PASSWORD"`
}

// LogConfig holds logging settings.
type LogConfig struct {
	Level  string `yaml:"LOG_LEVEL"`
	Format string `yaml:"LOG_FORMAT"`
}

type GameConfigTypeEnum int

const (
	DEVELOPER GameConfigTypeEnum = iota
	PRODUCT
)

// InitGameConfigWithDefaults sets safe default values for all config fields.
func (gc *GameConfig) InitGameConfigWithDefaults() {
	fmt.Println("Initializing game config with default values...")
	gc.Title = "My Hobie 2D RPG Game"
	gc.Host = "localhost"
	gc.Port = 8080

	gc.Database = DatabaseConfig{
		Host:     "localhost",
		Port:     5432,
		Name:     "my_hobie_2d_rpg_game",
		User:     "postgres",
		Password: "postgres",
	}

	gc.JWT = JWTConfig{
		Secret:     "changeme",
		Expiration: "1h",
	}

	gc.Email = EmailConfig{
		Host:     "smtp.gmail.com",
		Port:     587,
		Username: "",
		Password: "",
	}

	gc.Log = LogConfig{
		Level:  "debug",
		Format: "json",
	}

	fmt.Printf("%+v\n", gc)
}

// GetConfig loads configuration from the appropriate YAML file.
// Falls back to defaults when the file cannot be read or parsed.
func (gc *GameConfig) GetConfig(gcte GameConfigTypeEnum) *GameConfig {
	var file []byte
	var fileErr error

	switch gcte {
	case DEVELOPER:
		file, fileErr = os.ReadFile("configuration.dev.yaml")
	case PRODUCT:
		file, fileErr = os.ReadFile("configuration.prod.yaml")
	default:
		file, fileErr = os.ReadFile("configuration.dev.yaml")
	}

	if fileErr != nil {
		gc.InitGameConfigWithDefaults()
		return gc
	}

	if ymlErr := yaml.Unmarshal(file, gc); ymlErr != nil {
		gc.InitGameConfigWithDefaults()
		return gc
	}

	return gc
}

// GetConfigWithEnvVars loads the YAML config and then overrides any field
// that has a corresponding environment variable set.
func (gc *GameConfig) GetConfigWithEnvVars(gcte GameConfigTypeEnum) *GameConfig {
	gc.GetConfig(gcte)

	// Game
	if v := os.Getenv("GAME_TITLE"); v != "" {
		gc.Title = v
	}
	if v := os.Getenv("GAME_HOST"); v != "" {
		gc.Host = v
	}
	if v := os.Getenv("GAME_PORT"); v != "" {
		fmt.Sscanf(v, "%d", &gc.Port)
	}

	// Database
	if v := os.Getenv("DATABASE_HOST"); v != "" {
		gc.Database.Host = v
	}
	if v := os.Getenv("DATABASE_PORT"); v != "" {
		fmt.Sscanf(v, "%d", &gc.Database.Port)
	}
	if v := os.Getenv("DATABASE_NAME"); v != "" {
		gc.Database.Name = v
	}
	if v := os.Getenv("DATABASE_USER"); v != "" {
		gc.Database.User = v
	}
	if v := os.Getenv("DATABASE_PASSWORD"); v != "" {
		gc.Database.Password = v
	}

	// JWT
	if v := os.Getenv("JWT_SECRET"); v != "" {
		gc.JWT.Secret = v
	}
	if v := os.Getenv("JWT_EXPIRATION"); v != "" {
		gc.JWT.Expiration = v
	}

	// Email
	if v := os.Getenv("EMAIL_HOST"); v != "" {
		gc.Email.Host = v
	}
	if v := os.Getenv("EMAIL_PORT"); v != "" {
		fmt.Sscanf(v, "%d", &gc.Email.Port)
	}
	if v := os.Getenv("EMAIL_USERNAME"); v != "" {
		gc.Email.Username = v
	}
	if v := os.Getenv("EMAIL_PASSWORD"); v != "" {
		gc.Email.Password = v
	}

	// Log
	if v := os.Getenv("LOG_LEVEL"); v != "" {
		gc.Log.Level = v
	}
	if v := os.Getenv("LOG_FORMAT"); v != "" {
		gc.Log.Format = v
	}

	return gc
}
