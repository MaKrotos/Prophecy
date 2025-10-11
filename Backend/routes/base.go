package routes

import (
	"prophecy/backend/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterBaseRoutes регистрирует базовые маршруты приложения
func RegisterBaseRoutes(router *gin.Engine) {
	// Базовые маршруты
	router.GET("/", handlers.Welcome)
	router.GET("/health", handlers.HealthCheck)
}
