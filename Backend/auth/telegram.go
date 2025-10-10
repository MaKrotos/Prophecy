package auth

import (
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
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
	ID         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Username   string `json:"username"`
	PhotoURL   string `json:"photo_url"`
	AuthDate   string `json:"auth_date"`
	Hash       string `json:"hash"`
	Signature  string `json:"signature"`
	QueryID    string `json:"query_id"`
	Chat       string `json:"chat"`
	StartParam string `json:"start_param"`
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
	// Преобразование строки auth_date в int64
	authDateInt, err := strconv.ParseInt(data.AuthDate, 10, 64)
	if err != nil {
		return false, fmt.Errorf("invalid auth_date format: %v", err)
	}

	// Проверка времени аутентификации (не старше 1 дня)
	if time.Now().Unix()-authDateInt > 86400 {
		return false, fmt.Errorf("auth data is too old")
	}

	// Если есть подпись, используем новый метод проверки Ed25519
	if data.Signature != "" {
		return ta.ValidateTokenEd25519(data)
	}

	// Иначе используем старый метод проверки HMAC-SHA-256
	return ta.ValidateTokenHMAC(data)
}

// ValidateTokenHMAC проверяет токен Telegram WebApp с использованием HMAC-SHA-256
func (ta *TelegramAuth) ValidateTokenHMAC(data *TelegramWebAppData) (bool, error) {
	// Создание данных для хэширования
	params := url.Values{}
	params.Add("auth_date", data.AuthDate) // Используем строковое значение напрямую
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
	if data.QueryID != "" {
		params.Add("query_id", data.QueryID)
	}
	if data.Chat != "" {
		params.Add("chat", data.Chat)
	}
	if data.StartParam != "" {
		params.Add("start_param", data.StartParam)
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
	secretKey := hmac.New(sha256.New, []byte("WebAppData"))
	secretKey.Write([]byte(ta.Token))
	h := hmac.New(sha256.New, secretKey.Sum(nil))
	h.Write([]byte(dataCheckString))
	calculatedHash := hex.EncodeToString(h.Sum(nil))

	// Сравнение хэшей
	if calculatedHash != data.Hash {
		return false, fmt.Errorf("invalid hash")
	}

	return true, nil
}

// ValidateTokenEd25519 проверяет токен Telegram WebApp с использованием Ed25519
func (ta *TelegramAuth) ValidateTokenEd25519(data *TelegramWebAppData) (bool, error) {
	// Для Ed25519 проверки нам нужно получить bot_id из токена
	// Формат токена: 123456789:ABCDEFabcdef1234567890ABCDEFabcd
	tokenParts := strings.Split(ta.Token, ":")
	if len(tokenParts) != 2 {
		return false, fmt.Errorf("invalid bot token format")
	}

	botID := tokenParts[0]

	// Создание данных для проверки
	params := url.Values{}
	params.Add("auth_date", data.AuthDate)
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
	if data.QueryID != "" {
		params.Add("query_id", data.QueryID)
	}
	if data.Chat != "" {
		params.Add("chat", data.Chat)
	}
	if data.StartParam != "" {
		params.Add("start_param", data.StartParam)
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

	// Формирование строки для проверки согласно документации Telegram
	checkString := botID + ":WebAppData\n" + dataCheckString

	// Декодирование подписи из base64url
	signature, err := base64.URLEncoding.DecodeString(data.Signature)
	if err != nil {
		return false, fmt.Errorf("invalid signature format: %v", err)
	}

	// Получение публичного ключа Telegram (в production используйте production ключ)
	// Тестовая среда: 40055058a4ee38156a06562e52eece92a771bcd8346a8c4615cb7376eddf72ec
	publicKeyHex := "40055058a4ee38156a06562e52eece92a771bcd8346a8c4615cb7376eddf72ec"
	publicKey, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return false, fmt.Errorf("invalid public key format: %v", err)
	}

	// Проверка длины публичного ключа
	if len(publicKey) != ed25519.PublicKeySize {
		return false, fmt.Errorf("invalid public key size")
	}

	// Проверка подписи с использованием Ed25519
	if !ed25519.Verify(publicKey, []byte(checkString), signature) {
		return false, fmt.Errorf("invalid signature")
	}

	return true, nil
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

	// Если есть подпись, используем Ed25519 проверку
	signature := parsedData.Get("signature")
	if signature != "" {
		valid, userData, err := ta.validateInitDataEd25519(parsedData)
		return valid, userData, err
	}

	// Иначе используем HMAC-SHA-256 проверку
	valid, userData, err := ta.validateInitDataHMAC(parsedData)
	return valid, userData, err
}

// validateInitDataHMAC проверяет initData с использованием HMAC-SHA-256
func (ta *TelegramAuth) validateInitDataHMAC(parsedData url.Values) (bool, *TelegramWebAppData, error) {
	// Получаем хеш из параметров
	hash := parsedData.Get("hash")
	if hash == "" {
		return false, nil, fmt.Errorf("hash not found in init data")
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

	// Создаем структуру с данными пользователя
	userData := &TelegramWebAppData{
		Hash:       hash,
		AuthDate:   parsedData.Get("auth_date"),
		QueryID:    parsedData.Get("query_id"),
		Chat:       parsedData.Get("chat"),
		StartParam: parsedData.Get("start_param"),
		Signature:  parsedData.Get("signature"),
	}

	// Парсим данные пользователя из JSON строки
	if userStr := parsedData.Get("user"); userStr != "" {
		// Для простоты предполагаем, что user - это JSON строка
		// В реальной реализации нужно бы парсить JSON, но для этой задачи
		// мы просто извлекаем ID из строки
		userData.ID = extractUserID(userStr)
		userData.FirstName = extractUserFirstName(userStr)
		userData.LastName = extractUserLastName(userStr)
		userData.Username = extractUserUsername(userStr)
		userData.PhotoURL = extractUserPhotoURL(userStr)
	}

	return true, userData, nil
}

// validateInitDataEd25519 проверяет initData с использованием Ed25519
func (ta *TelegramAuth) validateInitDataEd25519(parsedData url.Values) (bool, *TelegramWebAppData, error) {
	// Получаем подпись из параметров
	signature := parsedData.Get("signature")
	if signature == "" {
		return false, nil, fmt.Errorf("signature not found in init data")
	}

	// Удаляем подпись из данных и сортируем ключи
	var keys []string
	for key := range parsedData {
		if key != "signature" {
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

	// Для Ed25519 проверки нам нужно получить bot_id из токена
	tokenParts := strings.Split(ta.Token, ":")
	if len(tokenParts) != 2 {
		return false, nil, fmt.Errorf("invalid bot token format")
	}

	botID := tokenParts[0]

	// Формирование строки для проверки согласно документации Telegram
	checkString := botID + ":WebAppData\n" + dataCheckString

	// Декодирование подписи из base64url
	signatureBytes, err := base64.URLEncoding.DecodeString(signature)
	if err != nil {
		return false, nil, fmt.Errorf("invalid signature format: %v", err)
	}

	// Получение публичного ключа Telegram (в production используйте production ключ)
	// Тестовая среда: 40055058a4ee38156a06562e52eece92a771bcd8346a8c4615cb7376eddf72ec
	publicKeyHex := "40055058a4ee38156a06562e52eece92a771bcd8346a8c4615cb7376eddf72ec"
	publicKey, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return false, nil, fmt.Errorf("invalid public key format: %v", err)
	}

	// Проверка длины публичного ключа
	if len(publicKey) != ed25519.PublicKeySize {
		return false, nil, fmt.Errorf("invalid public key size")
	}

	// Проверка подписи с использованием Ed25519
	if !ed25519.Verify(publicKey, []byte(checkString), signatureBytes) {
		return false, nil, fmt.Errorf("invalid signature")
	}

	// Создаем структуру с данными пользователя
	userData := &TelegramWebAppData{
		Signature:  signature,
		AuthDate:   parsedData.Get("auth_date"),
		QueryID:    parsedData.Get("query_id"),
		Chat:       parsedData.Get("chat"),
		StartParam: parsedData.Get("start_param"),
	}

	// Парсим данные пользователя из JSON строки
	if userStr := parsedData.Get("user"); userStr != "" {
		// Для простоты предполагаем, что user - это JSON строка
		// В реальной реализации нужно бы парсить JSON, но для этой задачи
		// мы просто извлекаем ID из строки
		userData.ID = extractUserID(userStr)
		userData.FirstName = extractUserFirstName(userStr)
		userData.LastName = extractUserLastName(userStr)
		userData.Username = extractUserUsername(userStr)
		userData.PhotoURL = extractUserPhotoURL(userStr)
	}

	return true, userData, nil
}

// hmacSHA256 вычисляет HMAC-SHA256
func hmacSHA256(data, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

// extractUserID извлекает ID пользователя из строки user
func extractUserID(userStr string) int64 {
	// Простая реализация для извлечения ID из строки вида:
	// {"id":123456789,"first_name":"John","last_name":"Doe","username":"johndoe","language_code":"en"}
	// В реальной реализации нужно бы парсить JSON
	start := strings.Index(userStr, "\"id\"")
	if start == -1 {
		return 0
	}
	start = strings.Index(userStr[start:], ":")
	if start == -1 {
		return 0
	}
	start += 1 // Пропускаем ":"
	end := strings.Index(userStr[start:], ",")
	if end == -1 {
		end = strings.Index(userStr[start:], "}")
	}
	if end == -1 {
		return 0
	}
	idStr := strings.TrimSpace(userStr[start : start+end])
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0
	}
	return id
}

// extractUserFirstName извлекает имя пользователя из строки user
func extractUserFirstName(userStr string) string {
	// Простая реализация для извлечения имени из строки вида:
	// {"id":123456789,"first_name":"John","last_name":"Doe","username":"johndoe","language_code":"en"}
	start := strings.Index(userStr, "\"first_name\"")
	if start == -1 {
		return ""
	}
	start = strings.Index(userStr[start:], ":")
	if start == -1 {
		return ""
	}
	start += 1 // Пропускаем ":"
	if start >= len(userStr) {
		return ""
	}
	// Проверяем, является ли следующий символ кавычкой
	if userStr[start] != '"' {
		return ""
	}
	start += 1 // Пропускаем открывающую кавычку
	end := strings.Index(userStr[start:], "\"")
	if end == -1 {
		return ""
	}
	return userStr[start : start+end]
}

// extractUserLastName извлекает фамилию пользователя из строки user
func extractUserLastName(userStr string) string {
	// Простая реализация для извлечения фамилии из строки вида:
	// {"id":123456789,"first_name":"John","last_name":"Doe","username":"johndoe","language_code":"en"}
	start := strings.Index(userStr, "\"last_name\"")
	if start == -1 {
		return ""
	}
	start = strings.Index(userStr[start:], ":")
	if start == -1 {
		return ""
	}
	start += 1 // Пропускаем ":"
	if start >= len(userStr) {
		return ""
	}
	// Проверяем, является ли следующий символ кавычкой
	if userStr[start] != '"' {
		return ""
	}
	start += 1 // Пропускаем открывающую кавычку
	end := strings.Index(userStr[start:], "\"")
	if end == -1 {
		return ""
	}
	return userStr[start : start+end]
}

// extractUserUsername извлекает имя пользователя (username) из строки user
func extractUserUsername(userStr string) string {
	// Простая реализация для извлечения username из строки вида:
	// {"id":123456789,"first_name":"John","last_name":"Doe","username":"johndoe","language_code":"en"}
	start := strings.Index(userStr, "\"username\"")
	if start == -1 {
		return ""
	}
	start = strings.Index(userStr[start:], ":")
	if start == -1 {
		return ""
	}
	start += 1 // Пропускаем ":"
	if start >= len(userStr) {
		return ""
	}
	// Проверяем, является ли следующий символ кавычкой
	if userStr[start] != '"' {
		return ""
	}
	start += 1 // Пропускаем открывающую кавычку
	end := strings.Index(userStr[start:], "\"")
	if end == -1 {
		return ""
	}
	return userStr[start : start+end]
}

// extractUserPhotoURL извлекает URL фото пользователя из строки user
func extractUserPhotoURL(userStr string) string {
	// Простая реализация для извлечения photo_url из строки вида:
	// {"id":123456789,"first_name":"John","last_name":"Doe","username":"johndoe","language_code":"en","photo_url":"https://..."}
	start := strings.Index(userStr, "\"photo_url\"")
	if start == -1 {
		return ""
	}
	start = strings.Index(userStr[start:], ":")
	if start == -1 {
		return ""
	}
	start += 1 // Пропускаем ":"
	if start >= len(userStr) {
		return ""
	}
	// Проверяем, является ли следующий символ кавычкой
	if userStr[start] != '"' {
		return ""
	}
	start += 1 // Пропускаем открывающую кавычку
	end := strings.Index(userStr[start:], "\"")
	if end == -1 {
		return ""
	}
	return userStr[start : start+end]
}
