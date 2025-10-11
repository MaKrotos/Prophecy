package routes

import (
	"prophecy/backend/auth"
	"prophecy/backend/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes регистрирует маршруты для работы с пользователями
func RegisterUserRoutes(router *gin.Engine) {
	// Маршруты для работы с пользователями
	router.GET("/users/:id", handlers.GetUser)
	router.POST("/users", handlers.CreateUser)

	// Маршрут для получения списка всех пользователей (только для администраторов)
	router.GET("/users", auth.JWTAuthMiddleware(), auth.AdminAuthMiddleware(), handlers.GetAllUsers)

	// Маршрут для получения статистики пользователей (только для администраторов)
	router.GET("/users/stats", auth.JWTAuthMiddleware(), auth.AdminAuthMiddleware(), handlers.GetUserStats)
}
