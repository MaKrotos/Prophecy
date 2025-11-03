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
		Role int `json:"role"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверка допустимых значений роли
	if requestData.Role != 1 && requestData.Role != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role. Allowed values: 1 (Архитектор) or 0 (Пользователь)"})
		return
	}

	// Преобразуем числовое значение в тип UserRole
	role := models.ParseRoleFromInt(requestData.Role)

	// Установка роли пользователю
	err = models.SetUserRole(userID, role)
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
