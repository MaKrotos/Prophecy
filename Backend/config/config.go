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
	SSLCertPath     string
	SSLKeyPath      string
	UseHTTPS        bool
	AdminTelegramID string
}

// GetConfig возвращает конфигурацию приложения
func GetConfig() *Config {
	sslCertPath := getEnv("SSL_CERT_PATH", "/etc/letsencrypt/live/prophecy.makrotos.ru/fullchain.pem")
	sslKeyPath := getEnv("SSL_KEY_PATH", "/etc/letsencrypt/live/prophecy.makrotos.ru/privkey.pem")

	// Проверяем наличие сертификатов для определения, использовать ли HTTPS
	useHTTPS := fileExists(sslCertPath) && fileExists(sslKeyPath)

	return &Config{
		ServerPort:      getEnv("SERVER_PORT", "8080"),
		DBHost:          getEnv("DB_HOST", "localhost"),
		DBPort:          getEnv("DB_PORT", "5432"),
		DBUser:          getEnv("DB_USER", "user"),
		DBPassword:      getEnv("DB_PASSWORD", "password"),
		DBName:          getEnv("DB_NAME", "prophecy"),
		JWTSecret:       getEnv("JWT_SECRET", "prophecy_jwt_secret_key"),
		AdminTelegramID: getEnv("ADMIN_TELEGRAM_ID", ""),
		SSLCertPath:     sslCertPath,
		SSLKeyPath:      sslCertPath,
		UseHTTPS:        useHTTPS,
	}
}

// getEnv возвращает значение переменной окружения или значение по умолчанию
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// fileExists проверяет, существует ли файл
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
