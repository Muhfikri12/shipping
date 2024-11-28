package database

import (
	"github.com/Muhfikri12/shipping/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Shipping{},
		&models.ShippingTracking{})

	return err
}
