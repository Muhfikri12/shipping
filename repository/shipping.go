package repository

import (
	"fmt"
	"time"

	"github.com/Muhfikri12/shipping/models"
	"gorm.io/gorm"
)

type ShippingRepository struct {
	db *gorm.DB
}

func NewShippingRepository(db *gorm.DB) *ShippingRepository {
	return &ShippingRepository{db}
}

func (s *ShippingRepository) GetShippingDetailByID(id uint) (*models.Shipping, error) {
	var shipping models.Shipping

	if err := s.db.Preload("ShippingTracking").Find(&shipping, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get shipping data for ID %d: %w", id, err)
	}

	return &shipping, nil
}

func (s *ShippingRepository) CreateShipping(shipping *models.Shipping) error {

	tx := s.db.Begin()

	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Create(&shipping).Error; err != nil {
		tx.Rollback()
		return err
	}

	shippingTracking := models.ShippingTracking{
		ShippingID: shipping.ID,
		Location:   "-",
		Status:     "Sedang Diproses",
		UpdatedAt:  time.Now(),
	}

	if err := tx.Create(&shippingTracking).Error; err != nil {

		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	if err := s.db.Preload("ShippingTracking").First(&shipping, shipping.ID).Error; err != nil {
		return err
	}

	return nil
}

func (s *ShippingRepository) UpdateStatus(id uint, shippingTracking *models.ShippingTracking) error {
	// Lakukan pembaruan dan tangkap error
	result := s.db.Model(&models.ShippingTracking{}).
		Where("id = ?", id).
		Updates(models.ShippingTracking{
			Location:  shippingTracking.Location,
			Status:    shippingTracking.Status,
			UpdatedAt: time.Now(),
		})

	// Cek jika ada error
	if result.Error != nil {
		return result.Error
	}

	// Cek jika tidak ada record yang diperbarui
	if result.RowsAffected == 0 {
		return fmt.Errorf("no record found with shipping_id %d", id)
	}

	s.db.First(&shippingTracking, id)

	return nil
}
