package service

import (
	"errors"
	"fmt"
	"os"
)

// SaveUser реалізує лінійний потік (Guard Clauses) та безпечну роботу з ресурсами
func SaveUser(u *User) error {
	file, err := os.Open("config.json")
	if err != nil {
		return fmt.Errorf("failed to open config: %w", err)
	}
	defer file.Close() // Безпечний виклик після перевірки помилки

	if u == nil {
		return errors.New("nil user")
	}
	if u.Email == "" {
		return errors.New("empty email")
	}

	if err := db.Save(u); err != nil {
		return fmt.Errorf("database error: %w", err)
	}

	return nil
}
