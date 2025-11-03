package auth

import (
	"net/http"

	"prophecy/backend/models"

	"github.com/gin-gonic/gin"
)

// ArchitectOrAdminMiddleware middleware для проверки, что пользователь является архитектором или админом
func ArchitectOrAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получение роли пользователя из контекста
		role, roleExists := c.Get("role")
		if !roleExists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User role not found in context"})
			c.Abort()
			return
		}

		// Получение информации о правах администратора из контекста
		isAdmin, adminExists := c.Get("is_admin")
		if !adminExists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Admin privileges not found in context"})
			c.Abort()
			return
		}

		// Преобразуем роль в тип UserRole
		userRole, ok := role.(models.UserRole)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role type in context"})
			c.Abort()
			return
		}

		// Проверка, является ли пользователь администратором или архитектором
		if !isAdmin.(bool) && userRole != models.RoleArchitect {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. Only architects or admins can perform this action."})
			c.Abort()
			return
		}

		// Продолжение выполнения
		c.Next()
	}
}
