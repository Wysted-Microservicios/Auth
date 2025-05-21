package model

import "time"

type Role string

// Roles
const (
	USER_ROLE  Role = "user"
	ADMIN_ROLE Role = "admin"
)

type User struct {
	ID        int64     `json:"id,omitempty"`
	Email     string    `json:"email,omitempty"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Roles     []Role    `json:"roles,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
