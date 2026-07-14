package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Farukcoder/eCommerce-go/backend/internal/domain"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) List(ctx context.Context) ([]domain.Category, error) {
	var categories []domain.Category
	err := r.db.WithContext(ctx).Table("categories").Order("sort_order ASC, name ASC").Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) GetByID(ctx context.Context, id int64) (*domain.Category, error) {
	var category domain.Category
	err := r.db.WithContext(ctx).Table("categories").Where("id = ?", id).Take(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) Create(ctx context.Context, input domain.Category) (*domain.Category, error) {
	input.CreatedAt = time.Now().UTC()
	input.UpdatedAt = input.CreatedAt
	if err := r.db.WithContext(ctx).Table("categories").Create(&input).Error; err != nil {
		return nil, err
	}
	return &input, nil
}

func (r *CategoryRepository) Update(ctx context.Context, id int64, input domain.Category) (*domain.Category, error) {
	updates := map[string]any{
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
	result := r.db.WithContext(ctx).Table("categories").Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("category not found")
	}
	return r.GetByID(ctx, id)
}

func (r *CategoryRepository) Delete(ctx context.Context, id int64) error {
	result := r.db.WithContext(ctx).Table("categories").Where("id = ?", id).Delete(&domain.Category{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("category not found")
	}
	return nil
}

func (r *CategoryRepository) CountProductsByCategory(ctx context.Context, categoryID int64) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Table("products").Where("category_id = ?", categoryID).Count(&count).Error
	return count, err
}

func (r *CategoryRepository) CountSubcategoriesByCategory(ctx context.Context, categoryID int64) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Table("subcategories").Where("category_id = ?", categoryID).Count(&count).Error
	return count, err
}
