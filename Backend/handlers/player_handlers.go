package handlers

import (
	"net/http"
	"strconv"

	"prophecy/backend/models"

	"github.com/gin-gonic/gin"
)

// AddPlayerToSession добавляет игрока к сессии
func AddPlayerToSession(c *gin.Context) {
	// Получаем ID сессии из параметров URL
	sessionID, err := getSessionIDFromParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получаем информацию о пользователе из контекста и базы данных
	user, err := getUserFromContext(c)
	if err != nil {
		// Определяем тип ошибки и возвращаем соответствующий HTTP статус
		if err.Error() == "user not authenticated" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Получаем сессию из базы данных
	session, err := getSessionByID(sessionID)
	if err != nil {
		// Определяем тип ошибки и возвращаем соответствующий HTTP статус
		if err.Error() == "session not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Проверяем права доступа
	// Игроки могут присоединяться к сессиям, админы и архитекторы могут добавлять игроков
	var playerID int
	if user.IsAdmin || user.Role == models.RoleArchitect {
		// Админы и архитекторы могут добавлять любого игрока
		playerIDParam := c.Query("player_id")
		if playerIDParam == "" {
			// Если не указан игрок, добавляем текущего пользователя
			playerID = user.ID
		} else {
			playerID, err = strconv.Atoi(playerIDParam)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
				return
			}
		}
	} else {
		// Обычные пользователи могут добавить только себя
		playerID = user.ID
	}

	// Проверяем, не является ли пользователь архитектором этой сессии
	if session.ArchitectID == playerID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Architect cannot be added as player"})
		return
	}

	// Добавляем игрока к сессии
	if err := models.AddPlayerToSession(playerID, sessionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add player to session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Player added to session successfully"})
}

// RemovePlayerFromSession удаляет игрока из сессии
func RemovePlayerFromSession(c *gin.Context) {
	// Получаем ID сессии из параметров URL
	sessionID, err := getSessionIDFromParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получаем информацию о пользователе из контекста и базы данных
	user, err := getUserFromContext(c)
	if err != nil {
		// Определяем тип ошибки и возвращаем соответствующий HTTP статус
		if err.Error() == "user not authenticated" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Получаем сессию из базы данных
	session, err := getSessionByID(sessionID)
	if err != nil {
		// Определяем тип ошибки и возвращаем соответствующий HTTP статус
		if err.Error() == "session not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Определяем, какого игрока нужно удалить
	var playerID int
	if user.IsAdmin || (user.Role == models.RoleArchitect && session.ArchitectID == user.ID) {
		// Админы и архитекторы своей сессии могут удалить любого игрока
		playerIDParam := c.Query("player_id")
		if playerIDParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Player ID is required"})
			return
		}

		playerID, err = strconv.Atoi(playerIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
			return
		}
	} else {
		// Обычные пользователи могут удалить только себя
		playerID = user.ID
	}

	// Удаляем игрока из сессии
	if err := models.RemovePlayerFromSession(playerID, sessionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove player from session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Player removed from session successfully"})
}

// GetSessionPlayers получает всех игроков в сессии
func GetSessionPlayers(c *gin.Context) {
	// Получаем ID сессии из параметров URL
	sessionID, err := getSessionIDFromParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверяем аутентификацию пользователя
	_, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Получаем сессию из базы данных
	_, err = getSessionByID(sessionID)
	if err != nil {
		// Определяем тип ошибки и возвращаем соответствующий HTTP статус
		if err.Error() == "session not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Проверяем права доступа
	// Все пользователи могут видеть игроков в сессии
	players, err := models.GetSessionPlayers(sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session players"})
		return
	}

	c.JSON(http.StatusOK, players)
}

// GetPlayerSessions получает все сессии, в которых участвует игрок
func GetPlayerSessions(c *gin.Context) {
	// Получаем информацию о пользователе из контекста и базы данных
	user, err := getUserFromContext(c)
	if err != nil {
		// Определяем тип ошибки и возвращаем соответствующий HTTP статус
		if err.Error() == "user not authenticated" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Определяем, для какого игрока получаем сессии
	var playerID int
	if user.IsAdmin {
		// Админы могут получить сессии любого игрока
		playerIDParam := c.Query("player_id")
		if playerIDParam == "" {
			// Если не указан игрок, получаем сессии текущего пользователя
			playerID = user.ID
		} else {
			playerID, err = strconv.Atoi(playerIDParam)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
				return
			}
		}
	} else {
		// Обычные пользователи могут получить только свои сессии
		playerID = user.ID
	}

	// Получаем сессии игрока
	sessions, err := models.GetPlayerSessions(playerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get player sessions"})
		return
	}

	c.JSON(http.StatusOK, sessions)
}

// JoinSessionByReferral присоединяет пользователя к сессии по реферальной ссылке
func JoinSessionByReferral(c *gin.Context) {
	// Проверяем метод запроса
	if c.Request.Method == "POST" {
		// Получаем информацию о пользователе из контекста и базы данных
		user, err := getUserFromContext(c)
		if err != nil {
			// Определяем тип ошибки и возвращаем соответствующий HTTP статус
			if err.Error() == "user not authenticated" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		// Получаем сессию из базы данных по реферальной ссылке
		session, err := models.GetSessionByReferralLink(c.Param("referral_link"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session"})
			return
		}

		if session == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
			return
		}

		// Проверяем, не является ли пользователь архитектором этой сессии
		if session.ArchitectID == user.ID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Architect cannot be added as player"})
			return
		}

		// Проверяем, не участвует ли пользователь уже в сессии
		isPlayer, err := models.IsPlayerInSession(user.ID, session.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check player status"})
			return
		}

		if isPlayer {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User is already in session"})
			return
		}

		// Добавляем игрока к сессии
		if err := models.AddPlayerToSession(user.ID, session.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add player to session"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Successfully joined session"})
	} else {
		// Для GET запроса возвращаем информацию о сессии
		session, err := models.GetSessionByReferralLink(c.Param("referral_link"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session"})
			return
		}

		if session == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
			return
		}

		c.JSON(http.StatusOK, session)
	}
}