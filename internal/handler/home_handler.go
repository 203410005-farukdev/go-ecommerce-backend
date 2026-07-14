package handler

import (
	"github.com/Farukcoder/eCommerce-go/backend/internal/dto"
	"github.com/Farukcoder/eCommerce-go/backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	service *service.HomeService
}

func NewHomeHandler(s *service.HomeService) *HomeHandler {
	return &HomeHandler{service: s}
}

func (h *HomeHandler) GetHomePage(c *fiber.Ctx) error {
	data, err := h.service.GetHomePageData(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse(
			fiber.StatusInternalServerError,
			"Failed to retrieve homepage data",
			nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(
		fiber.StatusOK,
		"Homepage data retrieved successfully",
		data,
	))
}
