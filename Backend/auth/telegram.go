package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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
	AuthDate  int64  `json:"auth_date"`
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

// ValidateToken проверяет токен Telegram WebApp
func (ta *TelegramAuth) ValidateToken(data *TelegramWebAppData) (bool, error) {
	// Проверка времени аутентификации (не старше 1 дня)
	if time.Now().Unix()-data.AuthDate > 86400 {
		return false, fmt.Errorf("auth data is too old")
	}

	// Создание данных для хэширования
	params := url.Values{}
	params.Add("auth_date", strconv.FormatInt(data.AuthDate, 10))
	params.Add("first_name", data.FirstName)
	params.Add("id", strconv.FormatInt(data.ID, 10))
	if data.LastName != "" {
		params.Add("last_name", data.LastName)
	}
	if data.PhotoURL != "" {
		params.Add("photo_url", data.PhotoURL)
	}
	if data.Username != "" {
		params.Add("username", data.Username)
	}

	// Сортировка параметров по ключу
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Создание строки данных
	var dataCheckArray []string
	for _, k := range keys {
		dataCheckArray = append(dataCheckArray, k+"="+params.Get(k))
	}
	dataCheckString := strings.Join(dataCheckArray, "\n")

	// Создание хэша
	secretKey := sha256.Sum256([]byte(ta.Token))
	h := hmac.New(sha256.New, secretKey[:])
	h.Write([]byte(dataCheckString))
	calculatedHash := hex.EncodeToString(h.Sum(nil))

	// Сравнение хэшей
	if calculatedHash != data.Hash {
		return false, fmt.Errorf("invalid hash")
	}

	return true, nil
}
