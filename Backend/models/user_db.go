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
		INSERT INTO telegram_users (telegram_id, first_name, last_name, username, photo_url, auth_date, generated_name, is_admin, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`

	now := time.Now()
	err := database.DB.QueryRow(query,
		telegramUser.TelegramID,
		telegramUser.FirstName,
		telegramUser.LastName,
		telegramUser.Username,
		telegramUser.PhotoURL,
		telegramUser.AuthDate,
		telegramUser.GeneratedName,
		telegramUser.IsAdmin,
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
		SELECT id, telegram_id, first_name, last_name, username, photo_url, auth_date, generated_name, is_admin, created_at
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
		&telegramUser.GeneratedName,
		&telegramUser.IsAdmin,
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

// GetAllTelegramUsers получает всех пользователей Telegram из базы данных с пагинацией
func GetAllTelegramUsers(limit, offset int) ([]TelegramUser, error) {
	var users []TelegramUser
	query := `
		SELECT id, telegram_id, first_name, last_name, username, photo_url, auth_date, generated_name, is_admin, created_at
		FROM telegram_users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`

	rows, err := database.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user TelegramUser
		err := rows.Scan(
			&user.ID,
			&user.TelegramID,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.PhotoURL,
			&user.AuthDate,
			&user.GeneratedName,
			&user.IsAdmin,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserStats получает статистику пользователей
func GetUserStats() (map[string]int, error) {
	stats := make(map[string]int)

	// Получаем общее количество пользователей
	var totalUsers int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM telegram_users").Scan(&totalUsers)
	if err != nil {
		return nil, err
	}
	stats["total_users"] = totalUsers

	// Получаем количество администраторов
	var adminUsers int
	err = database.DB.QueryRow("SELECT COUNT(*) FROM telegram_users WHERE is_admin = true").Scan(&adminUsers)
	if err != nil {
		return nil, err
	}
	stats["admin_users"] = adminUsers

	return stats, nil
}
