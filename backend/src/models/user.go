package models

import "errors"

// User ...
type User struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	CreateadAt string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	DeletedAt  string `json:"deleted_at,omitempty"`
}

// ValidateFields ...
func (u User) ValidateFields() error {
	if u.Name == "" {
		return errors.New("name is a required field")
	}
	if u.Email == "" {
		return errors.New("email is a required field")
	}
	if u.Password == "" && u.ID != 0 {
		return errors.New("password is a required field")
	}

	if len(u.Name) > 255 {
		return errors.New("name cannot be bigger than 255 characters")
	}
	if len(u.Email) > 255 {
		return errors.New("email cannot be bigger than 255 characters")
	}
	if len(u.Password) > 255 && u.ID != 0 {
		return errors.New("password cannot be bigger than 255 characters")
	}
	return nil
}
