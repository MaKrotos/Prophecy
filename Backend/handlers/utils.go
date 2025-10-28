package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
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

// getUserFromContext получает информацию о пользователе из контекста и базы данных
func getUserFromContext(c *gin.Context) (*models.TelegramUser, error) {
	userID, exists := c.Get("user_id")
	if !exists {
		return nil, fmt.Errorf("user not authenticated")
	}

	user, err := models.GetTelegramUserByID(userID.(int))
	if err != nil {
		return nil, fmt.Errorf("failed to get user information: %v", err)
	}

	return user, nil
}

// getSessionByID получает сессию по ID
func getSessionByID(sessionID int) (*models.Session, error) {
	session, err := models.GetSessionByID(sessionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get session: %v", err)
	}

	if session == nil {
		return nil, fmt.Errorf("session not found")
	}

	return session, nil
}

// checkSessionAccess проверяет права доступа к сессии
func checkSessionAccess(user *models.TelegramUser, session *models.Session, accessType string) error {
	switch accessType {
	case "view":
		// Админы, архитектор, создавший сессию, и участники сессии могут получить к ней доступ
		isPlayer, err := models.IsPlayerInSession(user.ID, session.ID)
		if err != nil {
			return fmt.Errorf("failed to check player status: %v", err)
		}

		if !user.IsAdmin && user.Role != "Архитектор" && session.ArchitectID != user.ID && !isPlayer {
			return fmt.Errorf("access denied")
		}
	case "modify":
		// Только архитектор, создавший сессию, или админ может обновить/удалить её
		if !user.IsAdmin && session.ArchitectID != user.ID {
			return fmt.Errorf("access denied")
		}
	case "join":
		// Игроки могут присоединяться к сессиям, админы и архитекторы могут добавлять игроков
		// Проверка будет выполнена в соответствующих обработчиках
		// Здесь просто проверим, что пользователь аутентифицирован
		if user == nil {
			return fmt.Errorf("user not authenticated")
		}
	default:
		return fmt.Errorf("unknown access type: %s", accessType)
	}

	return nil
}

// getSessionIDFromParam получает ID сессии из параметров URL
func getSessionIDFromParam(c *gin.Context) (int, error) {
	sessionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, fmt.Errorf("invalid session ID")
	}
	return sessionID, nil
}

// GetQRCodeData получает данные для QR-кода
func GetQRCodeData(c *gin.Context) {
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

	// Проверяем, участвует ли пользователь в сессии
	isPlayer, err := models.IsPlayerInSession(user.ID, session.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check player status"})
		return
	}

	// Если пользователь не участник сессии и не админ, возвращаем ошибку
	if !isPlayer && !user.IsAdmin && user.Role != "Архитектор" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Формируем данные для QR-кода
	qrData := map[string]interface{}{
		"player_id":   user.ID,
		"session_id":  session.ID,
		"session_name": session.Name,
	}

	c.JSON(http.StatusOK, qrData)
}