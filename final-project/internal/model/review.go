package model

import "time"

type Review struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	ShippingID uint      `json:"shipping_id" gorm:"not null"`
	Rating     uint      `json:"rating" gorm:"not null"`
	Comment    string    `json:"comment"`
	PhotoUrl   string    `json:"photoUrl"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime;not null"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime;not null"`
}

func (Review) TableName() string {
	return "reviews"
}
