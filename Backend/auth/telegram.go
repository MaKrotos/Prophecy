// telegram_webapp_validator.go
package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// TelegramWebAppData структура для хранения данных из Telegram WebApp
type TelegramWebAppData struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	PhotoURL  string `json:"photo_url"`
	AuthDate  string `json:"auth_date"`
	Hash      string `json:"hash"`
}

// TelegramAuth структура для проверки аутентификации
type TelegramAuth struct {
	Token string
}

// NewTelegramAuth создает новый экземпляр TelegramAuth
func NewTelegramAuth(token string) *TelegramAuth {
	return &TelegramAuth{
		Token: token,
	}
}

// ValidateInitData проверяет initData строку из Telegram WebApp
func (ta *TelegramAuth) ValidateInitData(initData string) (bool, *TelegramWebAppData, error) {
	// Парсим query строку
	parsedData, err := url.ParseQuery(initData)
	if err != nil {
		return false, nil, fmt.Errorf("failed to parse init data: %v", err)
	}

	// Получаем хеш из параметров
	hash := parsedData.Get("hash")
	if hash == "" {
		return false, nil, fmt.Errorf("hash not found in init data")
	}

	// Проверяем время аутентификации (не старше 1 дня)
	authDateStr := parsedData.Get("auth_date")
	if authDateStr == "" {
		return false, nil, fmt.Errorf("auth_date not found in init data")
	}

	authDateInt, err := strconv.ParseInt(authDateStr, 10, 64)
	if err != nil {
		return false, nil, fmt.Errorf("invalid auth_date format: %v", err)
	}

	if time.Now().Unix()-authDateInt > 86400 {
		return false, nil, fmt.Errorf("auth data is too old")
	}

	// Удаляем хеш из данных и сортируем ключи
	var keys []string
	for key := range parsedData {
		if key != "hash" {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)

	// Создаем строку для проверки в формате key=value
	var items []string
	for _, key := range keys {
		values := parsedData[key]
		if len(values) > 0 {
			items = append(items, key+"="+values[0])
		}
	}
	dataCheckString := strings.Join(items, "\n")

	// Генерируем секретный ключ
	secretKey := hmacSHA256([]byte(ta.Token), []byte("WebAppData"))

	// Генерируем хеш для проверки
	generatedHash := hex.EncodeToString(hmacSHA256([]byte(dataCheckString), secretKey))

	// Сравниваем хеши
	if !hmac.Equal([]byte(generatedHash), []byte(hash)) {
		return false, nil, fmt.Errorf("invalid hash")
	}

	// Парсим данные пользователя
	userData, err := extractUserData(parsedData)
	if err != nil {
		return false, nil, fmt.Errorf("failed to extract user data: %v", err)
	}

	return true, userData, nil
}

// extractUserData извлекает данные пользователя из parsedData
func extractUserData(parsedData url.Values) (*TelegramWebAppData, error) {
	userData := &TelegramWebAppData{
		Hash: parsedData.Get("hash"),
	}

	// Парсим auth_date как string
	if authDate := parsedData.Get("auth_date"); authDate != "" {
		userData.AuthDate = authDate
	}

	// Парсим данные пользователя из JSON строки user
	if userStr := parsedData.Get("user"); userStr != "" {
		var user struct {
			ID        int64  `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			PhotoURL  string `json:"photo_url"`
		}

		if err := json.Unmarshal([]byte(userStr), &user); err != nil {
			return nil, fmt.Errorf("failed to parse user JSON: %v", err)
		}

		userData.ID = user.ID
		userData.FirstName = user.FirstName
		userData.LastName = user.LastName
		userData.Username = user.Username
		userData.PhotoURL = user.PhotoURL
	}

	return userData, nil
}

// hmacSHA256 вычисляет HMAC-SHA256
func hmacSHA256(data, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}
