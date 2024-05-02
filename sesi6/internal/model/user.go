package model

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"primarykey"`
	Name         string    `json:"name" gorm:"not null"`
	Username     string    `json:"username" gorm:"unique;not null"`
	Password     []byte    `json:"-" gorm:"not null"`
	Email        string    `json:"email" gorm:"unique;not null" validate:"email"`
	Role         string    `json:"role" gorm:"not null"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"-" gorm:"autoCreateTime;not null"`
	UpdatedAt    time.Time `json:"-" gorm:"autoUpdateTime;not null"`
}

func (User) TableName() string {
	return "users"
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type UpdateUserRequest struct {
	Name     *string `json:"name"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	Email    *string `json:"email"`
	Role     *string `json:"role"`
}
