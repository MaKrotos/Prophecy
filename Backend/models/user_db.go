package models

import (
	"database/sql"
	"prophecy/backend/database"
	"time"
)

// CreateUser создает нового пользователя в базе данных
func CreateUser(user *User) error {
	query := `
		INSERT INTO users (name, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	now := time.Now()
	err := database.DB.QueryRow(query, user.Name, user.Email, now, now).Scan(&user.ID)
	if err != nil {
		return err
	}

	user.CreatedAt = now
	user.UpdatedAt = now

	return nil
}

// GetUserByID получает пользователя по ID из базы данных
func GetUserByID(id int) (*User, error) {
	var user User
	query := `
		SELECT id, name, email, created_at, updated_at
		FROM users
		WHERE id = $1`

	err := database.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// CreateTelegramUser создает нового пользователя Telegram в базе данных
func CreateTelegramUser(telegramUser *TelegramUser) error {
	query := `
		INSERT INTO telegram_users (telegram_id, first_name, last_name, username, photo_url, auth_date, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`

	now := time.Now()
	err := database.DB.QueryRow(query,
		telegramUser.TelegramID,
		telegramUser.FirstName,
		telegramUser.LastName,
		telegramUser.Username,
		telegramUser.PhotoURL,
		telegramUser.AuthDate,
		now,
	).Scan(&telegramUser.ID)

	if err != nil {
		return err
	}

	telegramUser.CreatedAt = now

	return nil
}

// GetTelegramUserByTelegramID получает пользователя Telegram по его Telegram ID
func GetTelegramUserByTelegramID(telegramID int64) (*TelegramUser, error) {
	var telegramUser TelegramUser
	query := `
		SELECT id, telegram_id, first_name, last_name, username, photo_url, auth_date, created_at
		FROM telegram_users
		WHERE telegram_id = $1`

	err := database.DB.QueryRow(query, telegramID).Scan(
		&telegramUser.ID,
		&telegramUser.TelegramID,
		&telegramUser.FirstName,
		&telegramUser.LastName,
		&telegramUser.Username,
		&telegramUser.PhotoURL,
		&telegramUser.AuthDate,
		&telegramUser.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &telegramUser, nil
}
