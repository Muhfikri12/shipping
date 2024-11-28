package seeder

import (
	"time"

	"github.com/Muhfikri12/shipping/models"
	"gorm.io/gorm"
)

func SeedShipping(db *gorm.DB) error {

	shippings := []models.Shipping{
		{
			IDOrder:          "ORD001",
			Ekspedisi:        "JNE",
			SellerCoordinate: "123,456",
			BuyerCoordinate:  "789,012",
			TotalPayment:     50000.00,
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
			ShippingTracking: models.ShippingTracking{
				Location:  "Warehouse A",
				Status:    "In Transit",
				UpdatedAt: time.Now(),
			},
		},
		{
			IDOrder:          "ORD002",
			Ekspedisi:        "SiCepat",
			SellerCoordinate: "234,567",
			BuyerCoordinate:  "890,123",
			TotalPayment:     75000.00,
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
			ShippingTracking: models.ShippingTracking{
				Location:  "Hub B",
				Status:    "Delivered",
				UpdatedAt: time.Now(),
			},
		},
		{
			IDOrder:          "ORD003",
			Ekspedisi:        "TIKI",
			SellerCoordinate: "345,678",
			BuyerCoordinate:  "901,234",
			TotalPayment:     30000.00,
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
			ShippingTracking: models.ShippingTracking{
				Location:  "Sorting Center",
				Status:    "Pending",
				UpdatedAt: time.Now(),
			},
		},
	}

	// Insert data ke database
	for _, shipping := range shippings {
		if err := db.Create(&shipping).Error; err != nil {
			return err
		}
	}

	return nil
}
