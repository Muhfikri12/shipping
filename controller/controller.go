package controller

import (
	"github.com/Muhfikri12/shipping/service"
	"go.uber.org/zap"
)

type Controller struct {
	// User UserController
	Shipping ShippingHandler
}

func NewController(service service.Service, logger *zap.Logger) *Controller {
	return &Controller{
		// User: *NewUserController(service.User, logger),
		Shipping: *NewShippingHandler(service.Shipping, logger),
	}
}
