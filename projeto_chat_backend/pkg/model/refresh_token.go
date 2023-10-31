package model

import "time"

type RefreshToken struct {
	Token     string    `json:"token"`
	UserID    int       `json:"user_id"`
	Expiry    time.Time `json:"expiry"`
	CreatedAt time.Time `json:"created_at"`
}
