package controller

import (
	"net/http"
	"strconv"

	"github.com/Muhfikri12/shipping/helper"
	"github.com/Muhfikri12/shipping/models"
	"github.com/Muhfikri12/shipping/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ShippingHandler struct {
	service service.ShippingService
	logger  *zap.Logger
}

func NewShippingHandler(service service.ShippingService, logger *zap.Logger) *ShippingHandler {
	return &ShippingHandler{service, logger}
}

func (h *ShippingHandler) GetShippingDetailByID(c *gin.Context) {
	// Parsing ID from URL params
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid ID parameter", zap.Error(err))
		helper.ResponseError(c, "INVALID_ID", "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	// Fetching shipping detail from service
	shipping, err := h.service.GetShippingDetailByID(uint(id))
	if err != nil {
		h.logger.Error("Shipping not found", zap.Error(err))
		helper.ResponseError(c, "NOT_FOUND", "Shipping not found", http.StatusNotFound)
		return
	}

	h.logger.Info("Shipping details retrieved successfully", zap.Any("shipping", nil))
	helper.ResponseOK(c, shipping, "Shipping details retrieved successfully")
}

func (h *ShippingHandler) CreateShipping(c *gin.Context) {

	shipping := models.Shipping{}

	if err := c.ShouldBindJSON(&shipping); err != nil {
		h.logger.Error("Invalid payload", zap.Error(err))
		helper.ResponseError(c, "INVALID", "Invalid payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.CreateShipping(&shipping)
	if err != nil {
		h.logger.Error("Failed to create", zap.Error(err))
		helper.ResponseError(c, "FAILED", "Failed to create Shipping", http.StatusInternalServerError)
		return
	}

	h.logger.Info("Create shipping successfully", zap.Any("shipping", nil))
	helper.ResponseOK(c, shipping, "Create Shipping successfully")
}

func (h *ShippingHandler) UpdateStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	shipping := models.ShippingTracking{}

	if err := c.ShouldBindJSON(&shipping); err != nil {
		h.logger.Error("Invalid payload", zap.Error(err))
		helper.ResponseError(c, "INVALID", "Invalid payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.UpdateStatus(uint(id), &shipping)
	if err != nil {
		h.logger.Error("Failed to Update Status Shipping", zap.Error(err))
		helper.ResponseError(c, "FAILED", "Failed to Update Status Shipping", http.StatusInternalServerError)
		return
	}

	h.logger.Info("Successfully Updated Status", zap.Any("shipping", nil))
	helper.ResponseOK(c, shipping, "Successfully Updated Status")
}
