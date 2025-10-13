package models

import (
	"database/sql"
	"prophecy/backend/database"
	"time"
)

// CreateSession создает новую сессию в базе данных
func CreateSession(session *Session) error {
	query := `
		INSERT INTO sessions (name, description, architect_id, referral_link)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at`

	err := database.DB.QueryRow(query, session.Name, session.Description, session.ArchitectID, session.ReferralLink).
		Scan(&session.ID, &session.CreatedAt, &session.UpdatedAt)

	return err
}

// GetSessionsByArchitectID получает все сессии, созданные определенным архитектором
func GetSessionsByArchitectID(architectID int) ([]Session, error) {
	query := `
		SELECT id, name, description, architect_id, referral_link, created_at, updated_at
		FROM sessions
		WHERE architect_id = $1
		ORDER BY created_at DESC`

	rows, err := database.DB.Query(query, architectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []Session
	for rows.Next() {
		var session Session
		err := rows.Scan(
			&session.ID,
			&session.Name,
			&session.Description,
			&session.ArchitectID,
			&session.ReferralLink,
			&session.CreatedAt,
			&session.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}

	return sessions, rows.Err()
}

// GetAllSessions получает все сессии (для админов)
func GetAllSessions() ([]SessionWithArchitect, error) {
	query := `
		SELECT s.id, s.name, s.description, s.architect_id, s.referral_link, s.created_at, s.updated_at, u.generated_name as architect_name
		FROM sessions s
		JOIN telegram_users u ON s.architect_id = u.id
		ORDER BY s.created_at DESC`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []SessionWithArchitect
	for rows.Next() {
		var session SessionWithArchitect
		err := rows.Scan(
			&session.ID,
			&session.Name,
			&session.Description,
			&session.ArchitectID,
			&session.ReferralLink,
			&session.CreatedAt,
			&session.UpdatedAt,
			&session.ArchitectName,
		)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}

	return sessions, rows.Err()
}

// GetSessionByID получает сессию по ID
func GetSessionByID(id int) (*Session, error) {
	query := `
		SELECT id, name, description, architect_id, referral_link, created_at, updated_at
		FROM sessions
		WHERE id = $1`

	var session Session
	err := database.DB.QueryRow(query, id).Scan(
		&session.ID,
		&session.Name,
		&session.Description,
		&session.ArchitectID,
		&session.ReferralLink,
		&session.CreatedAt,
		&session.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &session, nil
}

// UpdateSession обновляет информацию о сессии
func UpdateSession(session *Session) error {
	query := `
		UPDATE sessions
		SET name = $1, description = $2, updated_at = CURRENT_TIMESTAMP
		WHERE id = $3`

	_, err := database.DB.Exec(query, session.Name, session.Description, session.ID)
	if err != nil {
		return err
	}

	// Обновляем время обновления в объекте
	session.UpdatedAt = time.Now()

	return nil
}

// DeleteSession удаляет сессию по ID
func DeleteSession(id int) error {
	query := `DELETE FROM sessions WHERE id = $1`
	_, err := database.DB.Exec(query, id)
	return err
}

// AddPlayerToSession добавляет игрока к сессии
func AddPlayerToSession(playerID, sessionID int) error {
	query := `
		INSERT INTO player_sessions (player_id, session_id)
		VALUES ($1, $2)
		ON CONFLICT (player_id, session_id) DO NOTHING`

	_, err := database.DB.Exec(query, playerID, sessionID)
	return err
}

// RemovePlayerFromSession удаляет игрока из сессии
func RemovePlayerFromSession(playerID, sessionID int) error {
	query := `DELETE FROM player_sessions WHERE player_id = $1 AND session_id = $2`
	_, err := database.DB.Exec(query, playerID, sessionID)
	return err
}

// GetSessionPlayers получает всех игроков в сессии
func GetSessionPlayers(sessionID int) ([]TelegramUser, error) {
	query := `
		SELECT u.id, u.telegram_id, u.first_name, u.last_name, u.username, u.photo_url, u.auth_date, u.generated_name, u.is_admin, u.role, u.created_at
		FROM player_sessions ps
		JOIN telegram_users u ON ps.player_id = u.id
		WHERE ps.session_id = $1
		ORDER BY ps.joined_at DESC`

	rows, err := database.DB.Query(query, sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []TelegramUser
	for rows.Next() {
		var player TelegramUser
		err := rows.Scan(
			&player.ID,
			&player.TelegramID,
			&player.FirstName,
			&player.LastName,
			&player.Username,
			&player.PhotoURL,
			&player.AuthDate,
			&player.GeneratedName,
			&player.IsAdmin,
			&player.Role,
			&player.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}

	return players, rows.Err()
}

// GetPlayerSessions получает все сессии, в которых участвует игрок
func GetPlayerSessions(playerID int) ([]Session, error) {
	query := `
		SELECT s.id, s.name, s.description, s.architect_id, s.created_at, s.updated_at
		FROM player_sessions ps
		JOIN sessions s ON ps.session_id = s.id
		WHERE ps.player_id = $1
		ORDER BY s.created_at DESC`

	rows, err := database.DB.Query(query, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []Session
	for rows.Next() {
		var session Session
		err := rows.Scan(
			&session.ID,
			&session.Name,
			&session.Description,
			&session.ArchitectID,
			&session.CreatedAt,
			&session.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}

	return sessions, rows.Err()
}

// IsPlayerInSession проверяет, участвует ли игрок в сессии
func IsPlayerInSession(playerID, sessionID int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM player_sessions WHERE player_id = $1 AND session_id = $2)`

	var exists bool
	err := database.DB.QueryRow(query, playerID, sessionID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// GetSessionByReferralLink получает сессию по реферальной ссылке
func GetSessionByReferralLink(referralLink string) (*Session, error) {
	query := `
		SELECT id, name, description, architect_id, referral_link, created_at, updated_at
		FROM sessions
		WHERE referral_link = $1`

	var session Session
	err := database.DB.QueryRow(query, referralLink).Scan(
		&session.ID,
		&session.Name,
		&session.Description,
		&session.ArchitectID,
		&session.ReferralLink,
		&session.CreatedAt,
		&session.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &session, nil
}
