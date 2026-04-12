package domain

import "time"

type User struct {
	Id           int        `json:"id"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"password_hash"`
	IsDeleted    bool       `json:"is_deleted"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	CreatedAt    *time.Time `json:"created_at"`
}

type UserForList struct {
	Id        int        `json:"id"`
	Email     string     `json:"email"`
	IsDeleted bool       `json:"is_deleted"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	CreatedAt *time.Time `json:"created_at"`
	RoleName  string     `json:"role_name"`
	RoleTitle string     `json:"role_title"`
}
