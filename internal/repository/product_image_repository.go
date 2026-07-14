package repository

import (
	"context"

	"github.com/Farukcoder/eCommerce-go/backend/internal/domain"
	"gorm.io/gorm"
)

type ProductImageRepository struct {
	db *gorm.DB
}

func NewProductImageRepository(db *gorm.DB) *ProductImageRepository {
	return &ProductImageRepository{db: db}
}

func (r *ProductImageRepository) CreateBatch(ctx context.Context, images []domain.ProductImage) ([]domain.ProductImage, error) {
	if len(images) == 0 {
		return images, nil
	}
	if err := r.db.WithContext(ctx).Table("product_images").Create(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}

func (r *ProductImageRepository) DeleteByURLs(ctx context.Context, productID int64, urls []string) error {
	if len(urls) == 0 {
		return nil
	}
	result := r.db.WithContext(ctx).
		Table("product_images").
		Where("product_id = ? AND url IN ?", productID, urls).
		Delete(&domain.ProductImage{})
	return result.Error
}
