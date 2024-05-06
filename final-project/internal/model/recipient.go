package model

import (
	"time"

	"gorm.io/gorm"
)

type Recipient struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	Name        string    `json:"name" gorm:"not null" valid:"required~Nama tidak boleh kosong"`
	Phonenumber string    `json:"phonenumber" valid:"required~Nomor telepon tidak boleh kosong"`
	Address     string    `json:"address" gorm:"not null" valid:"required~Alamat tidak boleh kosong"`
	City        string    `json:"city" gorm:"not null" valid:"required~Kota tidak boleh kosong"`
	Province    string    `json:"province" gorm:"not null" valid:"required~Provinsi tidak boleh kosong"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime;not null"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime;not null"`

	Shippings []Shipping `json:"shippings" gorm:"foreignKey:RecipientID"`
}

type CreateRecipientRequest struct {
	UserID      uint   `json:"user_id" valid:"required~User ID tidak boleh kosong"`
	Name        string `json:"name" valid:"required~Nama tidak boleh kosong"`
	Phonenumber string `json:"phonenumber" valid:"required~Nomor telepon tidak boleh kosong"`
	Address     string `json:"address" valid:"required~Alamat tidak boleh kosong"`
	City        string `json:"city" valid:"required~Kota tidak boleh kosong"`
	Province    string `json:"province" valid:"required~Provinsi tidak boleh kosong"`
}

type UpdateRecipientRequest struct {
	Name        *string `json:"name"`
	Phonenumber *string `json:"phonenumber"`
	Address     *string `json:"address"`
	City        *string `json:"city"`
	Province    *string `json:"province"`
}

func (Recipient) TableName() string {
	return "recipients"
}

func (r *Recipient) BeforeCreate(*gorm.DB) (err error) {
	if r.Phonenumber[0] == '0' {
		r.Phonenumber = "62" + r.Phonenumber[1:]
	} else if r.Phonenumber[0] == '8' {
		r.Phonenumber = "62" + r.Phonenumber
	}
	return
}
