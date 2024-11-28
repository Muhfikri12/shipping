package database

import (
	"fmt"

	"github.com/Muhfikri12/shipping/database/seeder"
	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) error {
	if err := seeder.SeedShipping(db); err != nil {
		return fmt.Errorf("failed to seed Shipping data: %w", err)
	}

	return nil
}
