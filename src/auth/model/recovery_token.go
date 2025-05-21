package model

import "time"

type RecoveryToken struct {
	ID         int64     `json:"id"`
	Token      string    `json:"token"`
	IDUser     int64     `json:"idUser"`
	IsUsed     bool      `json:"isUsed"`
	Expires_at time.Time `json:"expires_at,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}
