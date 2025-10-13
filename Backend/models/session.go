package models

import (
	"time"
)

// Session представляет сессию/комнату, созданную архитектором
type Session struct {
	ID           int       `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Description  string    `json:"description" db:"description"`
	ArchitectID  int       `json:"architect_id" db:"architect_id"`
	ReferralLink string    `json:"referral_link" db:"referral_link"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// SessionWithArchitect включает информацию об архитекторе
type SessionWithArchitect struct {
	Session
	ArchitectName string `json:"architect_name" db:"architect_name"`
}
