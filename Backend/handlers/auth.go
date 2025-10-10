package handlers

import (
	"net/http"
	"os"
	"strconv"

	"prophecy/backend/auth"
	"prophecy/backend/models"

	"github.com/gin-gonic/gin"
)

// TelegramAuthRequest структура для запроса аутентификации Telegram
type TelegramAuthRequest struct {
	ID        int64  `json:"id" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	PhotoURL  string `json:"photo_url"`
	AuthDate  string `json:"auth_date" binding:"required"`
	Hash      string `json:"hash" binding:"required"`
}

// ValidateTelegramToken проверяет токен Telegram WebApp
func ValidateTelegramToken(c *gin.Context) {
	var req TelegramAuthRequest

	// Проверка и парсинг JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Преобразование строки auth_date в int64
	authDateInt, err := strconv.ParseInt(req.AuthDate, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid auth_date format"})
		return
	}

	// Создание экземпляра TelegramAuth
	// В production лучше хранить токен в переменных окружения
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		// Для тестирования используем заглушку
		botToken = "123456789:ABCDEFabcdef1234567890ABCDEFabcd"
	}

	telegramAuth := auth.NewTelegramAuth(botToken)

	// Создание данных для проверки
	data := &auth.TelegramWebAppData{
		ID:        req.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		PhotoURL:  req.PhotoURL,
		AuthDate:  req.AuthDate,
		Hash:      req.Hash,
	}

	// Проверка токена
	valid, err := telegramAuth.ValidateToken(data)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Invalid token",
			"message": err.Error(),
		})
		return
	}

	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Invalid token",
			"message": "Token validation failed",
		})
		return
	}

	// Сохранение пользователя Telegram в базе данных
	telegramUser := &models.TelegramUser{
		TelegramID: data.ID,
		FirstName:  data.FirstName,
		LastName:   data.LastName,
		Username:   data.Username,
		PhotoURL:   data.PhotoURL,
		AuthDate:   authDateInt,
	}

	err = models.CreateTelegramUser(telegramUser)
	if err != nil {
		// Если пользователь уже существует, получаем его из базы
		existingUser, getErr := models.GetTelegramUserByTelegramID(data.ID)
		if getErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to save or retrieve user",
				"message": err.Error(),
			})
			return
		}
		telegramUser = existingUser
	}

	// Генерация JWT токена
	token, err := auth.GenerateJWT(telegramUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to generate token",
			"message": err.Error(),
		})
		return
	}

	// Если токен валиден, возвращаем успех с JWT токеном
	c.JSON(http.StatusOK, gin.H{
		"message": "Token is valid",
		"token":   token,
		"user": gin.H{
			"id":          telegramUser.ID,
			"telegram_id": telegramUser.TelegramID,
			"first_name":  telegramUser.FirstName,
			"last_name":   telegramUser.LastName,
			"username":    telegramUser.Username,
			"photo_url":   telegramUser.PhotoURL,
			"auth_date":   telegramUser.AuthDate,
			"created_at":  telegramUser.CreatedAt,
		},
	})
}

// GetTelegramBotToken возвращает токен бота (для тестирования)
func GetTelegramBotToken(c *gin.Context) {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		botToken = "TELEGRAM_BOT_TOKEN_NOT_SET"
	}

	c.JSON(http.StatusOK, gin.H{
		"bot_token": botToken,
	})
}

// VerifyJWT проверяет JWT токен и возвращает информацию о пользователе
func VerifyJWT(c *gin.Context) {
	// Получение claims из контекста (установлены в middleware)
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Claims not found in context"})
		return
	}

	// Преобразование claims к нужному типу
	jwtClaims, ok := claims.(*auth.JWTClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse claims"})
		return
	}

	// Получение информации о пользователе из базы данных
	telegramUser, err := models.GetTelegramUserByTelegramID(jwtClaims.TelegramID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user", "message": err.Error()})
		return
	}

	if telegramUser == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Token is valid",
		"user": gin.H{
			"id":          telegramUser.ID,
			"telegram_id": telegramUser.TelegramID,
			"first_name":  telegramUser.FirstName,
			"last_name":   telegramUser.LastName,
			"username":    telegramUser.Username,
			"photo_url":   telegramUser.PhotoURL,
			"auth_date":   telegramUser.AuthDate,
			"created_at":  telegramUser.CreatedAt,
		},
	})
}
