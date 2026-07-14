package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Farukcoder/eCommerce-go/backend/internal/domain"
	"gorm.io/gorm"
)

type SubcategoryRepository struct {
	db *gorm.DB
}

func NewSubcategoryRepository(db *gorm.DB) *SubcategoryRepository {
	return &SubcategoryRepository{db: db}
}

func (r *SubcategoryRepository) List(ctx context.Context) ([]domain.Subcategory, error) {
	var subcategories []domain.Subcategory
	err := r.db.WithContext(ctx).Table("subcategories").Order("sort_order ASC, name ASC").Find(&subcategories).Error
	return subcategories, err
}

func (r *SubcategoryRepository) GetByID(ctx context.Context, id int64) (*domain.Subcategory, error) {
	var subcategory domain.Subcategory
	err := r.db.WithContext(ctx).Table("subcategories").Where("id = ?", id).Take(&subcategory).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("subcategory not found")
		}
		return nil, err
	}
	return &subcategory, nil
}

func (r *SubcategoryRepository) Create(ctx context.Context, input domain.Subcategory) (*domain.Subcategory, error) {
	input.CreatedAt = time.Now().UTC()
	input.UpdatedAt = input.CreatedAt
	if err := r.db.WithContext(ctx).Table("subcategories").Create(&input).Error; err != nil {
		return nil, err
	}
	return &input, nil
}

func (r *SubcategoryRepository) Update(ctx context.Context, id int64, input domain.Subcategory) (*domain.Subcategory, error) {
	updates := map[string]any{
		"category_id": input.CategoryID,
		"name":        input.Name,
		"slug":        input.Slug,
		"description": input.Description,
		"sort_order":  input.SortOrder,
		"is_active":   input.IsActive,
		"updated_at":  time.Now().UTC(),
	}
	if input.ImageURL != nil {
		updates["image_url"] = input.ImageURL
	}
	result := r.db.WithContext(ctx).Table("subcategories").Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("subcategory not found")
	}
	return r.GetByID(ctx, id)
}

func (r *SubcategoryRepository) Delete(ctx context.Context, id int64) error {
	result := r.db.WithContext(ctx).Table("subcategories").Where("id = ?", id).Delete(&domain.Subcategory{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("subcategory not found")
	}
	return nil
}
