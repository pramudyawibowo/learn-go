package model

type Water struct {
	ID     uint   `json:"id" gorm:"primarykey"`
	Value  int    `json:"value" gorm:"not null"`
	Status string `json:"status" gorm:"not null"`
}
