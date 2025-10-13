package auth

import (
	"time"

	"prophecy/backend/config"
	"prophecy/backend/models"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims структура для хранения пользовательских данных в токене
type JWTClaims struct {
	UserID        int    `json:"user_id"`
	TelegramID    int64  `json:"telegram_id"`
	GeneratedName string `json:"generated_name"`
	IsAdmin       bool   `json:"is_admin"`
	Role          string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateJWT генерирует JWT токен для пользователя
func GenerateJWT(user *models.TelegramUser) (string, error) {
	cfg := config.GetConfig()

	// Создание claims с пользовательскими данными
	claims := &JWTClaims{
		UserID:        user.ID,
		TelegramID:    user.TelegramID,
		GeneratedName: user.GeneratedName,
		IsAdmin:       user.IsAdmin,
		Role:          user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Токен действует 24 часа
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "prophecy-backend",
		},
	}

	// Создание токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подпись токена
	tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT проверяет JWT токен и возвращает claims
func ValidateJWT(tokenString string) (*JWTClaims, error) {
	cfg := config.GetConfig()

	// Парсинг токена
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// Проверка валидности токена
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrTokenInvalidClaims
}
