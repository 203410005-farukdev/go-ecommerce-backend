package handler

import (
	"github.com/Farukcoder/eCommerce-go/backend/internal/dto"
	"github.com/Farukcoder/eCommerce-go/backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

type SettingHandler struct {
	service *service.SettingService
}

func NewSettingHandler(s *service.SettingService) *SettingHandler {
	return &SettingHandler{service: s}
}

func (h *SettingHandler) GetSettings(c *fiber.Ctx) error {
	settings, err := h.service.GetSettings(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse(
			fiber.StatusInternalServerError,
			"Failed to retrieve settings: "+err.Error(),
			nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(
		fiber.StatusOK,
		"Settings retrieved successfully",
		settings,
	))
}

func (h *SettingHandler) UpdateSettings(c *fiber.Ctx) error {
	var updates map[string]string
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(
			fiber.StatusBadRequest,
			"Invalid request body: "+err.Error(),
			nil,
		))
	}

	if err := h.service.UpdateSettings(c.Context(), updates); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse(
			fiber.StatusInternalServerError,
			"Failed to update settings: "+err.Error(),
			nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(
		fiber.StatusOK,
		"Settings updated successfully",
		nil,
	))
}
