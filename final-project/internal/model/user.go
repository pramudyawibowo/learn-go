package model

import (
	"final-project/internal/helper"
	"log"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID           uint      `json:"id" gorm:"primarykey"`
	Name         string    `json:"name" gorm:"not null" valid:"required~Nama tidak boleh kosong"`
	Username     string    `json:"username" gorm:"unique;not null" valid:"required~Username tidak boleh kosong,alphanum~Username hanya boleh berisi huruf dan angka,runelength(5|20)~Username harus memiliki panjang 5-20 karakter"`
	Password     string    `json:"password" gorm:"not null" valid:"required~Password tidak boleh kosong"`
	Email        string    `json:"email" gorm:"unique;not null" valid:"required~Email tidak boleh kosong,email~Email tidak valid"`
	Role         string    `json:"role" gorm:"not null"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"createdAt" gorm:"autoCreateTime;not null"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"autoUpdateTime;not null"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {

	// validate user
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		log.Println(err)
		return err
	}

	// hash password
	hashedPassword, err := helper.HashPassword(u.Password)
	if err != nil {
		log.Println(err)
		return err
	}

	u.Password = hashedPassword
	u.Role = "user"
	return nil
}

type LoginRequest struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required~Password tidak boleh kosong"`
}
