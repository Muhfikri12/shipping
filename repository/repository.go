package repository

import "gorm.io/gorm"

type Repository struct {
	// User UserRepository
	Shipping ShippingRepository
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		// User: *NewUserRepository(db),
		Shipping: *NewShippingRepository(db),
	}
}
