package models

import (
	"errors"
	"strings"
)

// LoginDTO ...
type LoginDTO struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	AuthToken string `json:"auth_token"`
}

// ValidateFields ...
func (l LoginDTO) ValidateFields() error {
	if l.Email == "" {
		return errors.New("email is a required field")
	}
	if l.Password == "" {
		return errors.New("password is a required field")
	}
	if len(l.Email) > 255 {
		return errors.New("email must be less than 255 characters")
	}
	if len(l.Password) > 255 {
		return errors.New("passwod must be less than 255 characters")
	}

	return nil
}

// FormatFields ....
func (l *LoginDTO) FormatFields() {
	l.Email = strings.TrimSpace(l.Email)
	l.Email = strings.ToLower(l.Email)
	l.Password = strings.TrimSpace(l.Password)
}
