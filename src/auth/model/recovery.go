package model

import "time"

type Recovery struct {
	ID         int64     `json:"id"`
	Code       string    `json:"code"`
	IDUser     int64     `json:"idUser"`
	IsActive   bool      `json:"isActive"`
	Expires_at time.Time `json:"expires_at,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}
