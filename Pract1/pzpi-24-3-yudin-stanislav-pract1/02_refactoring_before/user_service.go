package service

import (
	"errors"
	"os"
)

// SaveUser містить надмірну вкладеність ("піраміда приреченості") та ігнорує перевірку ресурсів
func SaveUser(u *User) error {
	file, _ := os.Open("config.json") // Ігнорування помилки
	defer file.Close()                // Ризик виклику Close() на nil-вказівнику

	if u != nil {
		if u.Email != "" {
			err := db.Save(u)
			if err == nil {
				return nil
			} else {
				return err
			}
		} else {
			return errors.New("empty email")
		}
	} else {
		return errors.New("nil user")
	}
}
