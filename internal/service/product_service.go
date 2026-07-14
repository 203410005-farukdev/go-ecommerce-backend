package service

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/Farukcoder/eCommerce-go/backend/internal/domain"
	"github.com/Farukcoder/eCommerce-go/backend/internal/dto"
	"github.com/Farukcoder/eCommerce-go/backend/internal/repository"
)

type ProductService struct {
	repo        *repository.ProductRepository
	variantRepo *repository.ProductVariantRepository
	imageRepo   *repository.ProductImageRepository
}

func NewProductService(repo *repository.ProductRepository, variantRepo *repository.ProductVariantRepository, imageRepo *repository.ProductImageRepository) *ProductService {
	return &ProductService{repo: repo, variantRepo: variantRepo, imageRepo: imageRepo}
}

func (s *ProductService) ListProducts(ctx context.Context) ([]domain.Product, error) {
	return s.repo.List(ctx)
}

func (s *ProductService) GetProduct(ctx context.Context, id string) (*domain.Product, error) {
	productID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, productID)
}

func (s *ProductService) CreateProduct(ctx context.Context, input dto.CreateProductInput) (*domain.Product, error) {
	normalized, err := normalizeProduct(input)
	if err != nil {
		return nil, err
	}
	product, err := s.repo.Create(ctx, normalized)
	if err != nil {
		return nil, err
	}

	if len(input.GalleryURLs) > 0 {
		var images []domain.ProductImage
		for index, rawURL := range input.GalleryURLs {
			url := strings.TrimSpace(rawURL)
			if url == "" {
				continue
			}
			images = append(images, domain.ProductImage{
				ProductID: product.ID,
				URL:       url,
				SortOrder: index,
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			})
		}
		if len(images) > 0 {
			if _, err := s.imageRepo.CreateBatch(ctx, images); err != nil {
				return nil, err
			}
			product.ProductImages = images
		}
	}

	return product, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, id string, input dto.UpdateProductInput) (*domain.Product, error) {
	productID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	normalized, err := normalizeProduct(dto.CreateProductInput(input))
	if err != nil {
		return nil, err
	}
	product, err := s.repo.Update(ctx, productID, normalized)
	if err != nil {
		return nil, err
	}

	if len(input.RemovedGalleryURLs) > 0 {
		if err := s.imageRepo.DeleteByURLs(ctx, productID, input.RemovedGalleryURLs); err != nil {
			return nil, err
		}
	}

	if len(input.GalleryURLs) > 0 {
		var images []domain.ProductImage
		for index, rawURL := range input.GalleryURLs {
			url := strings.TrimSpace(rawURL)
			if url == "" {
				continue
			}
			images = append(images, domain.ProductImage{
				ProductID: product.ID,
				URL:       url,
				SortOrder: index,
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			})
		}
		if len(images) > 0 {
			if _, err := s.imageRepo.CreateBatch(ctx, images); err != nil {
				return nil, err
			}
			product.ProductImages = append(product.ProductImages, images...)
		}
	}

	return product, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id string) error {
	productID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, productID)
}

func (s *ProductService) ListProductVariants(ctx context.Context, productID string) ([]domain.ProductVariant, error) {
	pid, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		return nil, err
	}
	return s.variantRepo.ListByProductID(ctx, pid)
}

func (s *ProductService) CreateProductVariant(ctx context.Context, productID string, input dto.CreateProductVariantInput) (*domain.ProductVariant, error) {
	pid, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		return nil, err
	}
	normalized, err := normalizeProductVariant(input)
	if err != nil {
		return nil, err
	}
	normalized.ProductID = pid
	return s.variantRepo.Create(ctx, normalized)
}

func (s *ProductService) UpdateProductVariant(ctx context.Context, variantID string, input dto.UpdateProductVariantInput) (*domain.ProductVariant, error) {
	vid, err := strconv.ParseInt(variantID, 10, 64)
	if err != nil {
		return nil, err
	}
	normalized, err := normalizeProductVariant(dto.CreateProductVariantInput(input))
	if err != nil {
		return nil, err
	}
	return s.variantRepo.Update(ctx, vid, normalized)
}

func (s *ProductService) DeleteProductVariant(ctx context.Context, variantID string) error {
	vid, err := strconv.ParseInt(variantID, 10, 64)
	if err != nil {
		return err
	}
	return s.variantRepo.Delete(ctx, vid)
}

func normalizeProduct(input dto.CreateProductInput) (domain.Product, error) {
	name := strings.TrimSpace(input.Name)
	if name == "" {
		return domain.Product{}, errors.New("product name is required")
	}
	slug := strings.TrimSpace(input.Slug)
	if slug == "" {
		slug = slugify(name)
	}

	product := domain.Product{
		Name:          name,
		Slug:          slug,
		Price:         input.Price,
		CategoryID:    input.CategoryID,
		SubcategoryID: input.SubcategoryID,
		IsActive:      input.IsActive,
		IsFeatured:    input.IsFeatured,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	}

	if input.SKU != nil {
		sku := strings.TrimSpace(*input.SKU)
		if sku != "" {
			product.SKU = &sku
		}
	}

	if input.Description != nil {
		desc := strings.TrimSpace(*input.Description)
		if desc != "" {
			product.Description = &desc
		}
	}

	if input.CompareAtPrice != nil {
		product.CompareAtPrice = input.CompareAtPrice
	}

	if input.ImageURL != nil {
		image := strings.TrimSpace(*input.ImageURL)
		if image != "" {
			product.ImageURL = &image
		}
	}

	return product, nil
}

func normalizeProductVariant(input dto.CreateProductVariantInput) (domain.ProductVariant, error) {
	name := strings.TrimSpace(input.Name)
	if name == "" {
		return domain.ProductVariant{}, errors.New("variant name is required")
	}

	variant := domain.ProductVariant{
		Name:              name,
		Price:             input.Price,
		Quantity:          input.Quantity,
		LowStockThreshold: input.LowStockThreshold,
		IsActive:          input.IsActive,
		SortOrder:         input.SortOrder,
		CreatedAt:         time.Now().UTC(),
		UpdatedAt:         time.Now().UTC(),
	}

	if input.SKU != nil {
		sku := strings.TrimSpace(*input.SKU)
		if sku != "" {
			variant.SKU = &sku
		}
	}

	if input.CompareAtPrice != nil {
		variant.CompareAtPrice = input.CompareAtPrice
	}

	return variant, nil
}

func slugify(input string) string {
	s := strings.ToLower(strings.TrimSpace(input))
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, "_", "-")

	var b strings.Builder
	prevHyphen := false
	for _, r := range s {
		if r == '-' {
			if !prevHyphen {
				b.WriteRune('-')
				prevHyphen = true
			}
			continue
		}
		if strings.ContainsRune("abcdefghijklmnopqrstuvwxyz0123456789", r) || unicode.IsLetter(r) || unicode.IsNumber(r) {
			b.WriteRune(r)
			prevHyphen = false
			continue
		}
	}

	out := strings.Trim(b.String(), "-")
	return out
}
