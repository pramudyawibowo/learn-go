package model

import "time"

type Shipping struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	RecipientID uint      `json:"recipient_id" gorm:"not null"`
	Price       float64   `json:"price" gorm:"not null"`
	Weight      float64   `json:"weight" gorm:"not null"`
	Status      string    `json:"status" gorm:"not null"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime;not null"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime;not null"`

	Reviews []Review `json:"reviews" gorm:"foreignKey:ShippingID"`
}

func (Shipping) TableName() string {
	return "shippings"
}
