package models

import (
	"time"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Hidden from JSON output

	Name     string  `gorm:"type:varchar(100);not null" json:"name"`
	Barcode  string  `gorm:"type:varchar(50);not null" json:"barcode"`
	Price    float64 `gorm:"type:numeric(10,2);not null" json:"price"`
	Stock    int     `gorm:"not null;default:0" json:"stock"`
	Category string  `gorm:"type:varchar(50)" json:"category"`
}

