package database

import (
	"database/sql"
	"fmt"
	"log"
	"prophecy/backend/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB инициализирует подключение к базе данных
func InitDB() {
	cfg := config.GetConfig()

	// Формирование строки подключения
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Открытие подключения к базе данных
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	// Проверка подключения
	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	fmt.Println("Successfully connected to database")

	// Если указан ID администратора, обновляем пользователя в базе данных
	if cfg.AdminTelegramID != "" {
		updateAdminUser(cfg.AdminTelegramID)
	}
}

// updateAdminUser обновляет пользователя с указанным Telegram ID как администратора
func updateAdminUser(telegramID string) {
	query := `UPDATE telegram_users SET is_admin = TRUE WHERE telegram_id = $1`
	result, err := DB.Exec(query, telegramID)
	if err != nil {
		log.Printf("Failed to update admin user: %v", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Failed to get rows affected: %v", err)
		return
	}

	if rowsAffected > 0 {
		fmt.Printf("Successfully updated user with Telegram ID %s as admin\n", telegramID)
	} else {
		fmt.Printf("No user found with Telegram ID %s\n", telegramID)
	}
}
