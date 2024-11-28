package service

import (
	"github.com/Muhfikri12/shipping/models"
	"github.com/Muhfikri12/shipping/repository"
)

type ShippingService struct {
	repo repository.ShippingRepository
}

func NewShippingService(repo repository.ShippingRepository) *ShippingService {
	return &ShippingService{repo}
}

func (s *ShippingService) GetShippingDetailByID(id uint) (*models.Shipping, error) {
	shipping, err := s.repo.GetShippingDetailByID(id)
	if err != nil {
		return nil, err
	}

	return shipping, err
}

func (s *ShippingService) CreateShipping(shipping *models.Shipping) error {
	if err := s.repo.CreateShipping(shipping); err != nil {
		return err
	}

	return nil
}

func (s *ShippingService) UpdateStatus(id uint, shipping *models.ShippingTracking) error {
	if err := s.repo.UpdateStatus(id, shipping); err != nil {
		return err
	}

	return nil
}
