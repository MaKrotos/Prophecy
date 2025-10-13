package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strconv"

	"prophecy/backend/models"

	"github.com/gin-gonic/gin"
)

// generateReferralLink генерирует случайную строку для реферальной ссылки
func generateReferralLink() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// CreateSession создает новую сессию
func CreateSession(c *gin.Context) {
	// Получаем информацию о пользователе из контекста
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Получаем информацию о пользователе из базы данных
	user, err := models.GetTelegramUserByID(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user information"})
		return
	}

	// Парсим данные из запроса
	var requestData struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Генерируем реферальную ссылку
	referralLink, err := generateReferralLink()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate referral link"})
		return
	}

	// Создаем новую сессию
	session := &models.Session{
		Name:         requestData.Name,
		Description:  requestData.Description,
		ArchitectID:  user.ID,
		ReferralLink: referralLink,
	}

	if err := models.CreateSession(session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session"})
		return
	}

	c.JSON(http.StatusCreated, session)
}

// GetSessions получает список сессий
func GetSessions(c *gin.Context) {
	// Получаем информацию о пользователе из контекста
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Получаем информацию о пользователе из базы данных
	user, err := models.GetTelegramUserByID(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user information"})
		return
	}

	var sessions interface{}
	var sessionsErr error

	// Если пользователь админ, получаем все сессии
	if user.IsAdmin {
		sessions, sessionsErr = models.GetAllSessions()
	} else if user.Role == "Архитектор" {
		// Если пользователь архитектор, получаем только его сессии
		sessions, sessionsErr = models.GetSessionsByArchitectID(user.ID)
	} else {
		// Для обычных пользователей возвращаем пустой список
		sessions = []models.Session{}
	}

	if sessionsErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get sessions"})
		return
	}

	c.JSON(http.StatusOK, sessions)
}

// GetSession получает информацию о конкретной сессии
func GetSession(c *gin.Context) {
	// Получаем ID сессии из параметров URL
	sessionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
		return
	}

	// Получаем информацию о пользователе из контекста
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Получаем информацию о пользователе из базы данных
	user, err := models.GetTelegramUserByID(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user information"})
		return
	}

	// Получаем сессию из базы данных
	session, err := models.GetSessionByID(sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session"})
		return
	}

	if session == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}

	// Проверяем права доступа к сессии
	// Админы и архитектор, создавший сессию, могут получить к ней доступ
	if !user.IsAdmin && user.Role != "Архитектор" && session.ArchitectID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	c.JSON(http.StatusOK, session)
}

// UpdateSession обновляет информацию о сессии
func UpdateSession(c *gin.Context) {
	// Получаем ID сессии из параметров URL
	sessionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
		return
	}

	// Получаем информацию о пользователе из контекста
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Получаем информацию о пользователе из базы данных
	user, err := models.GetTelegramUserByID(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user information"})
		return
	}

	// Получаем сессию из базы данных
	session, err := models.GetSessionByID(sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session"})
		return
	}

	if session == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}

	// Только архитектор, создавший сессию, или админ может обновить её
	if !user.IsAdmin && session.ArchitectID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Парсим данные из запроса
	var requestData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Обновляем данные сессии
	if requestData.Name != "" {
		session.Name = requestData.Name
	}
	if requestData.Description != "" {
		session.Description = requestData.Description
	}

	if err := models.UpdateSession(session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update session"})
		return
	}

	c.JSON(http.StatusOK, session)
}

// DeleteSession удаляет сессию
func DeleteSession(c *gin.Context) {
	// Получаем ID сессии из параметров URL
	sessionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
		return
	}

	// Получаем информацию о пользователе из контекста
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Получаем информацию о пользователе из базы данных
	user, err := models.GetTelegramUserByID(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user information"})
		return
	}

	// Получаем сессию из базы данных
	session, err := models.GetSessionByID(sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session"})
		return
	}

	if session == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}

	// Только архитектор, создавший сессию, или админ может удалить её
	if !user.IsAdmin && session.ArchitectID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Удаляем сессию
	if err := models.DeleteSession(sessionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Session deleted successfully"})
}

// AddPlayerToSession добавляет игрока к сессии
func AddPlayerToSession(c *gin.Context) {
	// Получаем ID сессии из параметров URL
	sessionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
		return
	}

	// Получаем информацию о пользователе из контекста
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Получаем информацию о пользователе из базы данных
	user, err := models.GetTelegramUserByID(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user information"})
		return
	}

	// Получаем сессию из базы данных
	session, err := models.GetSessionByID(sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session"})
		return
	}

	if session == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}

	// Проверяем права доступа
	// Игроки могут присоединяться к сессиям, админы и архитекторы могут добавлять игроков
	var playerID int
	if user.IsAdmin || user.Role == "Архитектор" {
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
	sessionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
		return
	}

	// Получаем информацию о пользователе из контекста
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Получаем информацию о пользователе из базы данных
	user, err := models.GetTelegramUserByID(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user information"})
		return
	}

	// Получаем сессию из базы данных
	session, err := models.GetSessionByID(sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session"})
		return
	}

	if session == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}

	// Определяем, какого игрока нужно удалить
	var playerID int
	if user.IsAdmin || (user.Role == "Архитектор" && session.ArchitectID == user.ID) {
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
	sessionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
		return
	}

	// Получаем информацию о пользователе из контекста
	_, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Получаем сессию из базы данных
	session, err := models.GetSessionByID(sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session"})
		return
	}

	if session == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
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
	// Получаем информацию о пользователе из контекста
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Получаем информацию о пользователе из базы данных
	user, err := models.GetTelegramUserByID(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user information"})
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
		// Получаем информацию о пользователе из контекста
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
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
		if session.ArchitectID == userID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Architect cannot be added as player"})
			return
		}

		// Проверяем, не участвует ли пользователь уже в сессии
		isPlayer, err := models.IsPlayerInSession(userID.(int), session.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check player status"})
			return
		}

		if isPlayer {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User is already in session"})
			return
		}

		// Добавляем игрока к сессии
		if err := models.AddPlayerToSession(userID.(int), session.ID); err != nil {
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
