package routes

import (
	"prophecy/backend/auth"
	"prophecy/backend/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes регистрирует маршруты для аутентификации
func RegisterAuthRoutes(router *gin.Engine) {
	// Маршруты для аутентификации Telegram WebApp
	router.POST("/auth/telegram", handlers.ValidateTelegramToken)
	router.GET("/auth/telegram/token", handlers.GetTelegramBotToken)

	// Маршруты для работы с JWT
	router.GET("/auth/verify", auth.JWTAuthMiddleware(), handlers.VerifyJWT)

	// Маршрут для проверки статуса администратора
	router.GET("/auth/admin", auth.JWTAuthMiddleware(), handlers.CheckAdminStatus)
}
