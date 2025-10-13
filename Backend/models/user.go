package models

import (
	"time"
)

// User представляет собой модель пользователя
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserResponse представляет собой ответ при создании пользователя
type UserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// TelegramUser представляет собой модель пользователя Telegram
type TelegramUser struct {
	ID            int       `json:"id"`
	TelegramID    int64     `json:"telegram_id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Username      string    `json:"username"`
	PhotoURL      string    `json:"photo_url"`
	AuthDate      int64     `json:"auth_date"`
	GeneratedName string    `json:"generated_name"`
	IsAdmin       bool      `json:"is_admin"`
	Role          string    `json:"role"`
	CreatedAt     time.Time `json:"created_at"`
}
