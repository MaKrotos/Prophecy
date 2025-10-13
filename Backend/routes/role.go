package routes

import (
	"prophecy/backend/auth"
	"prophecy/backend/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterRoleRoutes регистрирует маршруты для работы с ролями пользователей
func RegisterRoleRoutes(router *gin.Engine) {
	// Маршрут для установки роли пользователю (только для администраторов)
	router.PUT("/users/:id/role", auth.JWTAuthMiddleware(), auth.AdminAuthMiddleware(), handlers.SetUserRole)
}
