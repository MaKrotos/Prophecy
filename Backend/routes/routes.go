package routes

import (
	"prophecy/backend/auth"
	"prophecy/backend/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes регистрирует все маршруты приложения
func RegisterRoutes(router *gin.Engine) {
	// Базовые маршруты
	router.GET("/", handlers.Welcome)
	router.GET("/health", handlers.HealthCheck)

	// Маршруты для работы с пользователями
	router.GET("/users/:id", handlers.GetUser)
	router.POST("/users", handlers.CreateUser)

	// Маршруты для аутентификации Telegram WebApp
	router.POST("/auth/telegram", handlers.ValidateTelegramToken)
	router.GET("/auth/telegram/token", handlers.GetTelegramBotToken)

	// Маршруты для работы с JWT
	router.GET("/auth/verify", auth.JWTAuthMiddleware(), handlers.VerifyJWT)

	// Маршрут для проверки статуса администратора
	router.GET("/auth/admin", auth.JWTAuthMiddleware(), handlers.CheckAdminStatus)
}
