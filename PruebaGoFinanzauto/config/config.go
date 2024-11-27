package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	ServerPort string
	ServerHost string

	JWTSecret          string
	JWTExpirationHours int

	// Environment
	Environment string
}

func LoadConfig() (*Config, error) {
	godotenv.Load()

	config := &Config{}

	// Database Config
	config.DBHost = getEnv("DB_HOST", "localhost")
	config.DBPort = getEnv("DB_PORT", "5432")
	config.DBName = getEnv("DB_NAME", "")
	config.DBUser = getEnv("DB_USER", "")
	config.DBPassword = getEnv("DB_PASSWORD", "")
	config.DBSSLMode = getEnv("DB_SSL_MODE", "disable")

	// Server Config
	config.ServerHost = getEnv("SERVER_HOST", "0.0.0.0")
	config.ServerPort = getEnv("SERVER_PORT", "8080")

	// JWT Config
	config.JWTSecret = getEnv("JWT_SECRET", "")

	// Environment
	config.Environment = getEnv("ENVIRONMENT", "development")

	// Validar configuración requerida
	if err := config.validate(); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) validate() error {
	if c.DBName == "" {
		return errors.New("DB_NAME es requerido")
	}
	if c.DBUser == "" {
		return errors.New("DB_USER es requerido")
	}
	if c.DBPassword == "" {
		return errors.New("DB_PASSWORD es requerido")
	}
	if c.JWTSecret == "" {
		return errors.New("JWT_SECRET es requerido")
	}
	return nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func (c *Config) IsDevelopment() bool {
	return c.Environment == "development"
}

func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

func (c *Config) GetDSN() string {
	return "host=" + c.DBHost +
		" user=" + c.DBUser +
		" password=" + c.DBPassword +
		" dbname=" + c.DBName +
		" port=" + c.DBPort +
		" sslmode=" + c.DBSSLMode +
		" TimeZone=America/Bogota"
}
