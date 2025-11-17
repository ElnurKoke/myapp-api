package model

import "time"

type UserResponse struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	AvatarUrl *string `json:"avatar_url,omitempty"`
	Bio       string  `json:"bio,omitempty"`
	Phone     *string `json:"phone,omitempty"`
}

type RegisterRequest struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required,min=6"`
	RepeatPassword string `json:"repeat_password" binding:"required,eqfield=Password"`
}

type User struct {
	ID             int
	Name           string
	Email          string
	Password       string
	RepeatPassword string
	Role           string
	IsVerified     bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	LastLogin      *time.Time
	Status         string
	TotpSecret     *string
	AvatarUrl      *string
	PhoneNumber    *string
	Bio            string
	Auth           bool
}
