package routes

import (
	"prophecy/backend/auth"
	"prophecy/backend/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterSessionRoutes регистрирует маршруты для работы с сессиями
func RegisterSessionRoutes(router gin.IRouter) {
	// Группа маршрутов для сессий с JWT аутентификацией
	sessionGroup := router.Group("/sessions")
	sessionGroup.Use(auth.JWTAuthMiddleware())
	{
		// Создание новой сессии (только для архитекторов и админов)
		sessionGroup.POST("/", auth.ArchitectOrAdminMiddleware(), handlers.CreateSession)

		// Получение списка сессий
		sessionGroup.GET("/", handlers.GetSessions)

		// Получение информации о конкретной сессии
		sessionGroup.GET("/:id", handlers.GetSession)

		// Обновление информации о сессии
		sessionGroup.PUT("/:id", handlers.UpdateSession)

		// Удаление сессии
		sessionGroup.DELETE("/:id", handlers.DeleteSession)

		// Добавление игрока к сессии
		sessionGroup.POST("/:id/players", handlers.AddPlayerToSession)

		// Удаление игрока из сессии
		sessionGroup.DELETE("/:id/players", handlers.RemovePlayerFromSession)

		// Получение всех игроков в сессии
		sessionGroup.GET("/:id/players", handlers.GetSessionPlayers)

		// Присоединение к сессии по реферальной ссылке
		sessionGroup.POST("/join/:referral_link", handlers.JoinSessionByReferral)
		sessionGroup.GET("/join/:referral_link", handlers.JoinSessionByReferral)
	}

	// Группа маршрутов для получения сессий игрока
	playerGroup := router.Group("/players")
	playerGroup.Use(auth.JWTAuthMiddleware())
	{
		// Получение всех сессий, в которых участвует игрок
		playerGroup.GET("/sessions", handlers.GetPlayerSessions)
	}
}
