package models

import (
	"prophecy/backend/database"
	"time"
)

// SessionFriend представляет собой связь между пользователем и его другом в конкретной сессии
type SessionFriend struct {
	ID        int       `json:"id"`
	SessionID int       `json:"session_id"`
	UserID    int       `json:"user_id"`
	FriendID  int       `json:"friend_id"`
	CreatedAt time.Time `json:"created_at"`
}

// SessionFriendWithDetails включает детальную информацию о друге
type SessionFriendWithDetails struct {
	SessionFriend
	FriendName string `json:"friend_name"`
	Online     bool   `json:"online"`
}

// AddFriendToSession добавляет друга к пользователю в конкретной сессии
func AddFriendToSession(sessionID, userID, friendID int) error {
	query := `
		INSERT INTO session_friends (session_id, user_id, friend_id)
		VALUES ($1, $2, $3)
		ON CONFLICT (session_id, user_id, friend_id) DO NOTHING`

	_, err := database.DB.Exec(query, sessionID, userID, friendID)
	return err
}

// RemoveFriendFromSession удаляет друга пользователя из конкретной сессии
func RemoveFriendFromSession(sessionID, userID, friendID int) error {
	query := `DELETE FROM session_friends WHERE session_id = $1 AND user_id = $2 AND friend_id = $3`
	_, err := database.DB.Exec(query, sessionID, userID, friendID)
	return err
}

// GetSessionFriends получает список друзей пользователя в конкретной сессии
func GetSessionFriends(sessionID, userID int) ([]SessionFriendWithDetails, error) {
	query := `
		SELECT sf.id, sf.session_id, sf.user_id, sf.friend_id, sf.created_at, tu.generated_name as friend_name
		FROM session_friends sf
		JOIN telegram_users tu ON sf.friend_id = tu.id
		WHERE sf.session_id = $1 AND sf.user_id = $2
		ORDER BY sf.created_at DESC`

	rows, err := database.DB.Query(query, sessionID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var friends []SessionFriendWithDetails
	for rows.Next() {
		var friend SessionFriendWithDetails
		err := rows.Scan(
			&friend.ID,
			&friend.SessionID,
			&friend.UserID,
			&friend.FriendID,
			&friend.CreatedAt,
			&friend.FriendName,
		)
		if err != nil {
			return nil, err
		}
		// Пока все друзья онлайн
		friend.Online = true
		friends = append(friends, friend)
	}

	return friends, rows.Err()
}

// IsFriendInSession проверяет, является ли пользователь другом в конкретной сессии
func IsFriendInSession(sessionID, userID, friendID int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM session_friends WHERE session_id = $1 AND user_id = $2 AND friend_id = $3)`

	var exists bool
	err := database.DB.QueryRow(query, sessionID, userID, friendID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}