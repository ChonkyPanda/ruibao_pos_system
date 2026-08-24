package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Hidden from JSON output

	UserID uint 
	TransactionID uint
	Items []OrderItem
}