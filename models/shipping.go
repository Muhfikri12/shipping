package models

import (
	"time"

	"gorm.io/gorm"
)

type Shipping struct {
	ID               uint             `gorm:"primaryKey"`
	IDOrder          string           `gorm:"type:varchar(50);uniqueIndex" binding:"required"`
	Ekspedisi        string           `gorm:"type:varchar(50)" binding:"required"`
	SellerCoordinate string           `gorm:"type:varchar(50)" binding:"required"`
	BuyerCoordinate  string           `gorm:"type:varchar(50)" binding:"required"`
	TotalPayment     float64          `binding:"required,gt=0"`
	ShippingTracking ShippingTracking `gorm:"foreignKey:ShippingID;references:ID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *gorm.DeletedAt `gorm:"index"`
}

type ShippingTracking struct {
	ID         uint   `gorm:"primaryKey"`
	ShippingID uint   `gorm:"uniqueIndex"`
	Location   string `gorm:"type:varchar(100)"`
	Status     string `gorm:"type:varchar(50)"`
	UpdatedAt  time.Time
}
