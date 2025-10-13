package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes регистрирует все маршруты приложения
func RegisterRoutes(router *gin.Engine) {
	// Регистрация маршрутов из разных категорий
	RegisterBaseRoutes(router)
	RegisterUserRoutes(router)
	RegisterAuthRoutes(router)
	RegisterRoleRoutes(router)
}
