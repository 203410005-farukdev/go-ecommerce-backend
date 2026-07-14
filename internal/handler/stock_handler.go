package handler

import (
	"strings"

	"github.com/Farukcoder/eCommerce-go/backend/internal/dto"
	"github.com/Farukcoder/eCommerce-go/backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

type StockHandler struct {
	service *service.StockService
}

func NewStockHandler(s *service.StockService) *StockHandler {
	return &StockHandler{service: s}
}

func (h *StockHandler) ListStock(c *fiber.Ctx) error {
	stocks, err := h.service.ListStock(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse(fiber.StatusInternalServerError, "Failed to retrieve stock", nil))
	}
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(fiber.StatusOK, "Stock retrieved successfully", stocks))
}

func (h *StockHandler) GetStock(c *fiber.Ctx) error {
	stockID := c.Params("id")
	stock, err := h.service.GetStock(c.Context(), stockID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse(fiber.StatusNotFound, err.Error(), nil))
	}
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(fiber.StatusOK, "Stock retrieved successfully", stock))
}

func (h *StockHandler) UpdateStock(c *fiber.Ctx) error {
	stockID := c.Params("id")
	var input dto.UpdateStockInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Invalid request body", nil))
	}
	if input.Quantity < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Stock quantity cannot be negative", nil))
	}
	if input.LowStockThreshold < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Low stock threshold cannot be negative", nil))
	}
	if input.SKU != nil && strings.TrimSpace(*input.SKU) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "SKU cannot be empty", nil))
	}
	if input.Location != nil && strings.TrimSpace(*input.Location) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Location cannot be empty", nil))
	}
	stock, err := h.service.UpdateStock(c.Context(), stockID, input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, err.Error(), nil))
	}
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(fiber.StatusOK, "Stock updated successfully", stock))
}

func (h *StockHandler) ListStockMovements(c *fiber.Ctx) error {
	movements, err := h.service.ListStockMovements(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse(fiber.StatusInternalServerError, "Failed to retrieve stock movements", nil))
	}
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(fiber.StatusOK, "Stock movements retrieved successfully", movements))
}

func (h *StockHandler) CreateStockMovement(c *fiber.Ctx) error {
	var input dto.CreateStockMovementInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Invalid request body", nil))
	}
	movement, err := h.service.CreateStockMovement(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, err.Error(), nil))
	}
	return c.Status(fiber.StatusCreated).JSON(dto.SuccessResponse(fiber.StatusCreated, "Stock movement created successfully", movement))
}
