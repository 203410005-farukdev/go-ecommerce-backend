package handler

import (
	"strings"

	"github.com/Farukcoder/eCommerce-go/backend/internal/domain"
	"github.com/Farukcoder/eCommerce-go/backend/internal/dto"
	"github.com/Farukcoder/eCommerce-go/backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s}
}

func (h *CategoryHandler) ListCategories(c *fiber.Ctx) error {
	categories, err := h.service.ListCategories(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse(fiber.StatusInternalServerError, "Failed to retrieve categories", nil))
	}
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(fiber.StatusOK, "Categories retrieved successfully", categories))
}

func (h *CategoryHandler) GetCategory(c *fiber.Ctx) error {
	categoryID := c.Params("id")
	category, err := h.service.GetCategory(c.Context(), categoryID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse(fiber.StatusNotFound, err.Error(), nil))
	}
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(fiber.StatusOK, "Category retrieved successfully", category))
}

func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var input dto.CreateCategoryInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Invalid request body", nil))
	}
	if strings.TrimSpace(input.Name) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Category name is required", nil))
	}

	fileURL, err := saveUploadedFile(c, "image", "categories")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse(fiber.StatusInternalServerError, "Failed to upload image", nil))
	}
	if fileURL != "" {
		input.ImageURL = &fileURL
	}

	category, err := h.service.CreateCategory(c.Context(), domain.Category{
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
	return c.Status(fiber.StatusCreated).JSON(dto.SuccessResponse(fiber.StatusCreated, "Category created successfully", category))
}

func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	categoryID := c.Params("id")
	var input dto.UpdateCategoryInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Invalid request body", nil))
	}
	if strings.TrimSpace(input.Name) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Category name is required", nil))
	}

	fileURL, err := saveUploadedFile(c, "image", "categories")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse(fiber.StatusInternalServerError, "Failed to upload image", nil))
	}
	if fileURL != "" {
		input.ImageURL = &fileURL
	}

	category, err := h.service.UpdateCategory(c.Context(), categoryID, domain.Category{
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
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(fiber.StatusOK, "Category updated successfully", category))
}

func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	categoryID := c.Params("id")
	if err := h.service.DeleteCategory(c.Context(), categoryID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, err.Error(), nil))
	}
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(fiber.StatusOK, "Category deleted successfully", nil))
}
