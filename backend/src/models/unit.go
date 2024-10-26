package models

// Unit ...
type Unit struct {
	ID          int    `json:"id,omitempty"`
	Level       int    `json:"level"`
	Order       int    `json:"order"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Course      int    `json:"course"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	DeletedAt   string `json:"deleted_at,omitempty"`
}

// UnitDTO ...
type UnitDTO struct {
	ID          int       `json:"id"`
	Level       int       `json:"level"`
	Order       int       `json:"order"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Course      CourseDTO `json:"course"`
}
