package router

import (
	"github.com/Muhfikri12/shipping/infra"
	"github.com/gin-gonic/gin"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	r := gin.Default()

	r.GET("/shipping/:id", ctx.Ctl.Shipping.GetShippingDetailByID)
	r.POST("/shipping", ctx.Ctl.Shipping.CreateShipping)
	r.PUT("/shipping/:id", ctx.Ctl.Shipping.UpdateStatus)
	return r
}
