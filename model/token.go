package model

import "time"

type TokenPair struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type RefreshToken struct {
	ID        int
	UserID    int
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
	Revoked   bool
}
