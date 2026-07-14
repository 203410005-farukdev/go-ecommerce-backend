package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Farukcoder/eCommerce-go/backend/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) List(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.WithContext(ctx).
		Model(&domain.Product{}).
		Preload("Categories").
		Preload("Subcategories").
		Preload("Stock").
		Preload("ProductVariants").
		Preload("ProductImages").
		Order("created_at desc").
		Find(&products).Error
	return products, err
}

func (r *ProductRepository) GetByID(ctx context.Context, id int64) (*domain.Product, error) {
	var product domain.Product
	err := r.db.WithContext(ctx).
		Model(&domain.Product{}).
		Preload("Categories").
		Preload("Subcategories").
		Preload("Stock").
		Preload("ProductVariants").
		Preload("ProductImages").
		Where("id = ?", id).
		Take(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Create(ctx context.Context, input domain.Product) (*domain.Product, error) {
	input.CreatedAt = time.Now().UTC()
	input.UpdatedAt = input.CreatedAt
	if err := r.db.WithContext(ctx).Table("products").Create(&input).Error; err != nil {
		return nil, err
	}
	return &input, nil
}

func (r *ProductRepository) Update(ctx context.Context, id int64, input domain.Product) (*domain.Product, error) {
	updates := map[string]any{
		"name":             input.Name,
		"slug":             input.Slug,
		"sku":              input.SKU,
		"description":      input.Description,
		"price":            input.Price,
		"compare_at_price": input.CompareAtPrice,
		"category_id":      input.CategoryID,
		"subcategory_id":   input.SubcategoryID,
		"is_active":        input.IsActive,
		"is_featured":      input.IsFeatured,
		"updated_at":       time.Now().UTC(),
	}
	if input.ImageURL != nil {
		updates["image_url"] = input.ImageURL
	}
	result := r.db.WithContext(ctx).Table("products").Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("product not found")
	}
	return r.GetByID(ctx, id)
}

func (r *ProductRepository) Delete(ctx context.Context, id int64) error {
	result := r.db.WithContext(ctx).Table("products").Where("id = ?", id).Delete(&domain.Product{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}
	return nil
}
