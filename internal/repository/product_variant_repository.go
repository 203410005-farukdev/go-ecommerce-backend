package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Farukcoder/eCommerce-go/backend/internal/domain"
	"gorm.io/gorm"
)

type ProductVariantRepository struct {
	db *gorm.DB
}

func NewProductVariantRepository(db *gorm.DB) *ProductVariantRepository {
	return &ProductVariantRepository{db: db}
}

func (r *ProductVariantRepository) ListByProductID(ctx context.Context, productID int64) ([]domain.ProductVariant, error) {
	var variants []domain.ProductVariant
	err := r.db.WithContext(ctx).
		Model(&domain.ProductVariant{}).
		Where("product_id = ?", productID).
		Order("sort_order asc, created_at desc").
		Find(&variants).Error
	return variants, err
}

func (r *ProductVariantRepository) GetByID(ctx context.Context, id int64) (*domain.ProductVariant, error) {
	var variant domain.ProductVariant
	err := r.db.WithContext(ctx).Model(&domain.ProductVariant{}).Where("id = ?", id).Take(&variant).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product variant not found")
		}
		return nil, err
	}
	return &variant, nil
}

func (r *ProductVariantRepository) Create(ctx context.Context, input domain.ProductVariant) (*domain.ProductVariant, error) {
	input.CreatedAt = time.Now().UTC()
	input.UpdatedAt = input.CreatedAt
	if err := r.db.WithContext(ctx).Table("product_variants").Create(&input).Error; err != nil {
		return nil, err
	}
	return &input, nil
}

func (r *ProductVariantRepository) Update(ctx context.Context, id int64, input domain.ProductVariant) (*domain.ProductVariant, error) {
	result := r.db.WithContext(ctx).Table("product_variants").Where("id = ?", id).Updates(map[string]any{
		"name":                input.Name,
		"sku":                 input.SKU,
		"price":               input.Price,
		"compare_at_price":    input.CompareAtPrice,
		"quantity":            input.Quantity,
		"low_stock_threshold": input.LowStockThreshold,
		"is_active":           input.IsActive,
		"sort_order":          input.SortOrder,
		"updated_at":          time.Now().UTC(),
	})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("product variant not found")
	}
	return r.GetByID(ctx, id)
}

func (r *ProductVariantRepository) Delete(ctx context.Context, id int64) error {
	result := r.db.WithContext(ctx).Table("product_variants").Where("id = ?", id).Delete(&domain.ProductVariant{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("product variant not found")
	}
	return nil
}
