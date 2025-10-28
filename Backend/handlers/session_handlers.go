package handlers

import (
	"net/http"

	"prophecy/backend/models"

	"github.com/gin-gonic/gin"
)

// CreateSession создает новую сессию
func CreateSession(c *gin.Context) {
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

	// Проверяем права доступа к сессии
	if err := checkSessionAccess(user, session, "view"); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, session)
}

// UpdateSession обновляет информацию о сессии
func UpdateSession(c *gin.Context) {
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

	// Проверяем права доступа к сессии
	if err := checkSessionAccess(user, session, "modify"); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
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

	// Проверяем права доступа к сессии
	if err := checkSessionAccess(user, session, "modify"); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	// Удаляем сессию
	if err := models.DeleteSession(sessionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Session deleted successfully"})
}