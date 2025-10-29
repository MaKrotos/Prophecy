package config

import (
	"os"
)

// Config структура для хранения конфигурации приложения
type Config struct {
	ServerPort      string
	DBHost          string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBName          string
	JWTSecret       string
	AdminTelegramID string
}

// GetConfig возвращает конфигурацию приложения
func GetConfig() *Config {
	return &Config{
		ServerPort:      getEnv("SERVER_PORT", "8080"),
		DBHost:          getEnv("DB_HOST", "localhost"),
		DBPort:          getEnv("DB_PORT", "5432"),
		DBUser:          getEnv("DB_USER", "user"),
		DBPassword:      getEnv("DB_PASSWORD", "password"),
		DBName:          getEnv("DB_NAME", "prophecy"),
		JWTSecret:       getEnv("JWT_SECRET", "prophecy_jwt_secret_key"),
		AdminTelegramID: getEnv("ADMIN_TELEGRAM_ID", ""),
	}
}

// getEnv возвращает значение переменной окружения или значение по умолчанию
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

