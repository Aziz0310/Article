package models

import "time"

// Author ...
type Author struct {
	ID        string     `json:"id"`
	Firstname string     `json:"firstname" binding:"required" minLength:"2" maxLength:"50" example:"John" db:"firstname"`
	Lastname  string     `json:"lastname" binding:"required" minLength:"2" maxLength:"50" example:"Doe" db:"lastname"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

// CreateAuthorModel ...
type CreateAuthorModel struct {
	Firstname string `json:"firstname" binding:"required" minLength:"2" maxLength:"50" example:"John" db:"firstname"`
	Lastname  string `json:"lastname" binding:"required" minLength:"2" maxLength:"50" example:"Doe" db:"lastname"`
}

// UpdateAuthorModel ...
type UpdateAuthorModel struct {
	ID        string     `json:"id"`
	Firstname string     `json:"firstname" binding:"required" minLength:"2" maxLength:"50" example:"John" db:"firstname"`
	Lastname  string     `json:"lastname" binding:"required" minLength:"2" maxLength:"50" example:"Doe" db:"lastname"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// GetAuthorByIdResp ...
type GetAuthorByIdResp struct {
	ID        string     `json:"id"`
	Firstname string     `json:"firstname" binding:"required" minLength:"2" maxLength:"50" example:"John" db:"firstname"`
	Lastname  string     `json:"lastname" binding:"required" minLength:"2" maxLength:"50" example:"Doe" db:"lastname"`
	CreatedAt time.Time  `json:"created_at"`
	Article []Article
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}
