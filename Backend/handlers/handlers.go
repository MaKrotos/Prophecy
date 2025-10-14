package handlers

import (
	"net/http"
	"strconv"

	"prophecy/backend/models"

	"github.com/gin-gonic/gin"
)

// HealthCheck возвращает статус работоспособности приложения
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

// Welcome возвращает приветственное сообщение
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Prophecy Backend API",
		"status":  "success",
	})
}

// GetUser возвращает информацию о пользователе
func GetUser(c *gin.Context) {
	// Параметры из URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Получение пользователя из базы данных
	user, err := models.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser создает нового пользователя
func CreateUser(c *gin.Context) {
	var user models.User

	// Проверка и парсинг JSON
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Создание пользователя в базе данных
	err := models.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Создание ответа
	response := models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    response,
	})
}

// GetAllUsers возвращает список всех пользователей Telegram с пагинацией
func GetAllUsers(c *gin.Context) {
	// Получение параметров пагинации из запроса
	limit := 10 // значение по умолчанию
	offset := 0 // значение по умолчанию

	if limitParam := c.Query("limit"); limitParam != "" {
		if parsedLimit, err := strconv.Atoi(limitParam); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	if offsetParam := c.Query("offset"); offsetParam != "" {
		if parsedOffset, err := strconv.Atoi(offsetParam); err == nil && parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	// Ограничение максимального значения limit для предотвращения перегрузки
	if limit > 100 {
		limit = 100
	}

	// Получение пользователей из базы данных с пагинацией
	users, err := models.GetAllTelegramUsers(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserStats возвращает статистику пользователей
func GetUserStats(c *gin.Context) {
	// Получение статистики из базы данных
	stats, err := models.GetUserStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user stats"})
		return
	}

	// Преобразуем статистику в формат, ожидаемый фронтендом
	response := gin.H{
		"total_users": stats["total_users"],
		"admin_users": stats["admin_users"],
		"role_stats":  stats["role_stats"],
	}

	c.JSON(http.StatusOK, response)
}
