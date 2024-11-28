package service

import "github.com/Muhfikri12/shipping/repository"

type Service struct {
	// User UserService
	Shipping ShippingService
}

func NewService(repo repository.Repository) Service {
	return Service{
		// User: NewUserService(repo.User),
		Shipping: *NewShippingService(repo.Shipping),
	}
}
