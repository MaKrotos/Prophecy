package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware middleware для проверки JWT токена
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получение заголовка Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Проверка формата заголовка
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header must start with Bearer"})
			c.Abort()
			return
		}

		// Извлечение токена
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Валидация токена
		claims, err := ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "message": err.Error()})
			c.Abort()
			return
		}

		// Сохранение claims в контексте
		c.Set("user_id", claims.UserID)
		c.Set("telegram_id", claims.TelegramID)
		c.Set("generated_name", claims.GeneratedName)
		c.Set("is_admin", claims.IsAdmin)
		c.Set("role", claims.Role)
		c.Set("claims", claims)

		// Продолжение выполнения
		c.Next()
	}
}

// AdminAuthMiddleware middleware для проверки прав администратора
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получение информации о правах администратора из контекста
		isAdmin, exists := c.Get("is_admin")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Admin privileges required"})
			c.Abort()
			return
		}

		// Проверка, является ли пользователь администратором
		if !isAdmin.(bool) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin privileges required"})
			c.Abort()
			return
		}

		// Продолжение выполнения
		c.Next()
	}
}
