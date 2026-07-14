package handler

import (
	"strings"

	"github.com/Farukcoder/eCommerce-go/backend/internal/domain"
	"github.com/Farukcoder/eCommerce-go/backend/internal/dto"
	"github.com/Farukcoder/eCommerce-go/backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

type SubcategoryHandler struct {
	service *service.SubcategoryService
}

func NewSubcategoryHandler(s *service.SubcategoryService) *SubcategoryHandler {
	return &SubcategoryHandler{service: s}
}

func (h *SubcategoryHandler) ListSubcategories(c *fiber.Ctx) error {
	subcategories, err := h.service.ListSubcategories(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse(fiber.StatusInternalServerError, "Failed to retrieve subcategories", nil))
	}
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(fiber.StatusOK, "Subcategories retrieved successfully", subcategories))
}

func (h *SubcategoryHandler) GetSubcategory(c *fiber.Ctx) error {
	subcategoryID := c.Params("id")
	subcategory, err := h.service.GetSubcategory(c.Context(), subcategoryID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse(fiber.StatusNotFound, err.Error(), nil))
	}
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(fiber.StatusOK, "Subcategory retrieved successfully", subcategory))
}

func (h *SubcategoryHandler) CreateSubcategory(c *fiber.Ctx) error {
	var input dto.CreateSubcategoryInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Invalid request body", nil))
	}
	if strings.TrimSpace(input.Name) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Subcategory name is required", nil))
	}
	if input.CategoryID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Category ID is required", nil))
	}

	fileURL, err := saveUploadedFile(c, "image", "subcategories")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse(fiber.StatusInternalServerError, "Failed to upload image", nil))
	}
	if fileURL != "" {
		input.ImageURL = &fileURL
	}

	subcategory, err := h.service.CreateSubcategory(c.Context(), domain.Subcategory{
		CategoryID:  input.CategoryID,
		Name:        input.Name,
		Slug:        input.Slug,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		SortOrder:   input.SortOrder,
		IsActive:    input.IsActive,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, err.Error(), nil))
	}
	return c.Status(fiber.StatusCreated).JSON(dto.SuccessResponse(fiber.StatusCreated, "Subcategory created successfully", subcategory))
}

func (h *SubcategoryHandler) UpdateSubcategory(c *fiber.Ctx) error {
	subcategoryID := c.Params("id")
	var input dto.UpdateSubcategoryInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Invalid request body", nil))
	}
	if strings.TrimSpace(input.Name) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Subcategory name is required", nil))
	}
	if input.CategoryID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Category ID is required", nil))
	}

	fileURL, err := saveUploadedFile(c, "image", "subcategories")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse(fiber.StatusInternalServerError, "Failed to upload image", nil))
	}
	if fileURL != "" {
		input.ImageURL = &fileURL
	}

	subcategory, err := h.service.UpdateSubcategory(c.Context(), subcategoryID, domain.Subcategory{
		CategoryID:  input.CategoryID,
		Name:        input.Name,
		Slug:        input.Slug,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		SortOrder:   input.SortOrder,
		IsActive:    input.IsActive,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, err.Error(), nil))
	}
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(fiber.StatusOK, "Subcategory updated successfully", subcategory))
}

func (h *SubcategoryHandler) DeleteSubcategory(c *fiber.Ctx) error {
	subcategoryID := c.Params("id")
	if err := h.service.DeleteSubcategory(c.Context(), subcategoryID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, err.Error(), nil))
	}
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(fiber.StatusOK, "Subcategory deleted successfully", nil))
}
