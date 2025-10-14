package handlers

import (
	"net/http"
	"strconv"

	"prophecy/backend/models"

	"github.com/gin-gonic/gin"
)

// SetUserRole устанавливает роль пользователю
func SetUserRole(c *gin.Context) {
	// Получение ID пользователя из URL
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Получение роли из тела запроса
	var requestData struct {
		Role string `json:"role"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверка допустимых значений роли
	if requestData.Role != "Архитектор" && requestData.Role != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role. Allowed values: 'Архитектор' or empty"})
		return
	}

	// Установка роли пользователю
	err = models.SetUserRole(userID, requestData.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set user role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User role updated successfully",
		"user_id": userID,
		"role":    requestData.Role,
	})
}
