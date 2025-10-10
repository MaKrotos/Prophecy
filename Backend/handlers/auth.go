package handlers

import (
	"net/http"
	"os"
	"strconv"

	"prophecy/backend/auth"
	"prophecy/backend/models"
	"prophecy/backend/names"

	"github.com/gin-gonic/gin"
)

// TelegramAuthRequest структура для запроса аутентификации Telegram
type TelegramAuthRequest struct {
	InitData string `json:"initData" binding:"required"`
}

// ValidateTelegramToken проверяет токен Telegram WebApp
func ValidateTelegramToken(c *gin.Context) {
	var req TelegramAuthRequest

	// Проверка и парсинг JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	// Проверка токена с использованием initData
	valid, userData, err := telegramAuth.ValidateInitData(req.InitData)
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

	// Преобразование строки auth_date в int64
	authDateInt, err := strconv.ParseInt(userData.AuthDate, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid auth_date format"})
		return
	}

	// Генерация случайного имени
	generatedName := names.GenerateRandomName()

	// Проверка, существует ли пользователь Telegram в базе данных
	telegramUser, err := models.GetTelegramUserByTelegramID(userData.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve user",
			"message": err.Error(),
		})
		return
	}

	// Если пользователь не существует, создаем нового
	if telegramUser == nil {
		telegramUser = &models.TelegramUser{
			TelegramID:    userData.ID,
			FirstName:     userData.FirstName,
			LastName:      userData.LastName,
			Username:      userData.Username,
			PhotoURL:      userData.PhotoURL,
			AuthDate:      authDateInt,
			GeneratedName: generatedName,
		}

		err = models.CreateTelegramUser(telegramUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to save user",
				"message": err.Error(),
			})
			return
		}
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
			"id":             telegramUser.ID,
			"telegram_id":    telegramUser.TelegramID,
			"first_name":     telegramUser.FirstName,
			"last_name":      telegramUser.LastName,
			"generated_name": telegramUser.GeneratedName,
			"is_admin":       telegramUser.IsAdmin,
			"photo_url":      telegramUser.PhotoURL,
			"auth_date":      telegramUser.AuthDate,
			"created_at":     telegramUser.CreatedAt,
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
			"id":             telegramUser.ID,
			"telegram_id":    telegramUser.TelegramID,
			"first_name":     telegramUser.FirstName,
			"last_name":      telegramUser.LastName,
			"generated_name": telegramUser.GeneratedName,
			"is_admin":       telegramUser.IsAdmin,
			"photo_url":      telegramUser.PhotoURL,
			"auth_date":      telegramUser.AuthDate,
			"created_at":     telegramUser.CreatedAt,
		},
	})
}

// CheckAdminStatus проверяет статус администратора пользователя
func CheckAdminStatus(c *gin.Context) {
	// Получение is_admin из контекста (установлен в middleware)
	isAdmin, exists := c.Get("is_admin")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Admin status not found in context"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"is_admin": isAdmin,
	})
}
