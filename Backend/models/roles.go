package models

import (
	"encoding/json"
)

// UserRole представляет роль пользователя в системе
type UserRole int

const (
	// RoleUser обычная роль пользователя
	RoleUser UserRole = iota
	// RoleArchitect роль архитектора
	RoleArchitect
)

// MarshalJSON кастомная сериализация роли в JSON (возвращаем числовое значение)
func (r UserRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(r))
}

// UnmarshalJSON кастомная десериализация роли из JSON (принимаем числовое значение)
func (r *UserRole) UnmarshalJSON(data []byte) error {
	var roleInt int
	if err := json.Unmarshal(data, &roleInt); err != nil {
		return err
	}
	
	*r = ParseRoleFromInt(roleInt)
	return nil
}

// ParseRoleFromInt преобразует число в роль
func ParseRoleFromInt(roleInt int) UserRole {
	switch roleInt {
	case 1:
		return RoleArchitect
	default:
		return RoleUser
	}
}

// Int возвращает числовое представление роли
func (r UserRole) Int() int {
	return int(r)
}