package handler

import (
	"strings"

	"github.com/Farukcoder/eCommerce-go/backend/internal/dto"
	"github.com/Farukcoder/eCommerce-go/backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) ListProducts(c *fiber.Ctx) error {
	products, err := h.service.ListProducts(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse(fiber.StatusInternalServerError, "Failed to retrieve products", nil))
	}
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(fiber.StatusOK, "Products retrieved successfully", products))
}

func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	productID := c.Params("id")
	product, err := h.service.GetProduct(c.Context(), productID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse(fiber.StatusNotFound, err.Error(), nil))
	}
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(fiber.StatusOK, "Product retrieved successfully", product))
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var input dto.CreateProductInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Invalid request body", nil))
	}
	if strings.TrimSpace(input.Name) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Product name is required", nil))
	}

	fileURL, err := saveUploadedFile(c, "image", "products")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse(fiber.StatusInternalServerError, "Failed to upload image", nil))
	}
	if fileURL != "" {
		input.ImageURL = &fileURL
	}

	galleryURLs, err := saveUploadedFiles(c, "images", "products")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse(fiber.StatusInternalServerError, "Failed to upload gallery images", nil))
	}
	if len(galleryURLs) > 0 {
		input.GalleryURLs = galleryURLs
	}

	product, err := h.service.CreateProduct(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, err.Error(), nil))
	}
	return c.Status(fiber.StatusCreated).JSON(dto.SuccessResponse(fiber.StatusCreated, "Product created successfully", product))
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	productID := c.Params("id")
	var input dto.UpdateProductInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Invalid request body", nil))
	}
	if strings.TrimSpace(input.Name) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Product name is required", nil))
	}

	fileURL, err := saveUploadedFile(c, "image", "products")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse(fiber.StatusInternalServerError, "Failed to upload image", nil))
	}
	if fileURL != "" {
		input.ImageURL = &fileURL
	}

	galleryURLs, err := saveUploadedFiles(c, "images", "products")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse(fiber.StatusInternalServerError, "Failed to upload gallery images", nil))
	}
	if len(galleryURLs) > 0 {
		input.GalleryURLs = galleryURLs
	}

	product, err := h.service.UpdateProduct(c.Context(), productID, input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, err.Error(), nil))
	}
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(fiber.StatusOK, "Product updated successfully", product))
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	productID := c.Params("id")
	if err := h.service.DeleteProduct(c.Context(), productID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, err.Error(), nil))
	}
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(fiber.StatusOK, "Product deleted successfully", nil))
}

func (h *ProductHandler) ListProductVariants(c *fiber.Ctx) error {
	productID := c.Params("id")
	variants, err := h.service.ListProductVariants(c.Context(), productID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse(fiber.StatusInternalServerError, "Failed to retrieve product variants", nil))
	}
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(fiber.StatusOK, "Product variants retrieved successfully", variants))
}

func (h *ProductHandler) CreateProductVariant(c *fiber.Ctx) error {
	productID := c.Params("id")
	var input dto.CreateProductVariantInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Invalid request body", nil))
	}
	if strings.TrimSpace(input.Name) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Product variant name is required", nil))
	}
	variant, err := h.service.CreateProductVariant(c.Context(), productID, input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, err.Error(), nil))
	}
	return c.Status(fiber.StatusCreated).JSON(dto.SuccessResponse(fiber.StatusCreated, "Product variant created successfully", variant))
}

func (h *ProductHandler) UpdateProductVariant(c *fiber.Ctx) error {
	variantID := c.Params("variant_id")
	var input dto.UpdateProductVariantInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Invalid request body", nil))
	}
	if strings.TrimSpace(input.Name) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, "Product variant name is required", nil))
	}
	variant, err := h.service.UpdateProductVariant(c.Context(), variantID, input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, err.Error(), nil))
	}
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(fiber.StatusOK, "Product variant updated successfully", variant))
}

func (h *ProductHandler) DeleteProductVariant(c *fiber.Ctx) error {
	variantID := c.Params("variant_id")
	if err := h.service.DeleteProductVariant(c.Context(), variantID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse(fiber.StatusBadRequest, err.Error(), nil))
	}
	return c.Status(fiber.StatusOK).JSON(dto.SuccessResponse(fiber.StatusOK, "Product variant deleted successfully", nil))
}
