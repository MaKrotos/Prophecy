package handlers

import (
	"net/http"
	"strconv"

	"prophecy/backend/models"

	"github.com/gin-gonic/gin"
)

// AddFriendToSession добавляет друга к пользователю в конкретной сессии
func AddFriendToSession(c *gin.Context) {
	// Получаем ID сессии из параметров URL
	sessionID, err := getSessionIDFromParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получаем информацию о пользователе из контекста (пользователь, который добавляет друга)
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

	// Получаем ID друга из параметров запроса
	friendID, err := strconv.Atoi(c.Query("friend_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid friend ID"})
		return
	}

	// Проверяем, что пользователь не пытается добавить себя в друзья
	if user.ID == friendID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot add yourself as friend"})
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

	// Проверяем, участвует ли друг в сессии
	isFriendPlayer, err := models.IsPlayerInSession(friendID, session.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check friend status"})
		return
	}

	if !isFriendPlayer {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Friend is not a player in this session"})
		return
	}

	// Добавляем друга в список друзей пользователя в сессии
	if err := models.AddFriendToSession(sessionID, user.ID, friendID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add friend to session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend added to session successfully"})
}

// RemoveFriendFromSession удаляет друга пользователя из конкретной сессии
func RemoveFriendFromSession(c *gin.Context) {
	// Получаем ID сессии из параметров URL
	sessionID, err := getSessionIDFromParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получаем информацию о пользователе из контекста (пользователь, который удаляет друга)
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

	// Получаем ID друга из параметров запроса
	friendID, err := strconv.Atoi(c.Query("friend_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid friend ID"})
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

	// Удаляем друга из списка друзей пользователя в сессии
	if err := models.RemoveFriendFromSession(sessionID, user.ID, friendID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove friend from session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend removed from session successfully"})
}

// GetSessionFriends получает список друзей пользователя в сессии
func GetSessionFriends(c *gin.Context) {
	// Получаем ID сессии из параметров URL
	sessionID, err := getSessionIDFromParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получаем информацию о пользователе из контекста
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

	// Получаем список друзей пользователя в сессии
	friends, err := models.GetSessionFriends(sessionID, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session friends"})
		return
	}

	c.JSON(http.StatusOK, friends)
}