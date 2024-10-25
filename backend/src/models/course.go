package models

import "errors"

// Course ...
type Course struct {
	ID          int    `json:"id,omitempty"`
	Owner       int    `json:"owner"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AccessType  int8   `json:"access_type"`
	AccessCode  string `json:"access_code,omitempty`
	CreatedAt   string `json:"created_at",omitempty`
	UpdatedAt   string `json:"updated_at",omitempty`
	DeleteddAt  string `json:"deleted_at",omitempty`
}

func (c Course) validateFields() error {
	if c.Owner == 0 {
		return errors.New("owner is a required field")
	}
	if c.Name == "" {
		return errors.New("name is a required field")
	}
	if c.Description == "" {
		return errors.New("description is a required field")
	}
	if c.AccessType == 0 {
		return errors.New("access type is a required field")
	}
	if c.AccessCode == "" && c.AccessType == 1 {
		return errors.New("access code is a required field")
	}

	return nil
}