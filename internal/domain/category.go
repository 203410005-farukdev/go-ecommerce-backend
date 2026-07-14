package domain

import "time"

// Category is the DB entity for product categories.
type Category struct {
	ID               int64     `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Name             string    `json:"name" gorm:"column:name;not null"`
	Slug             string    `json:"slug" gorm:"column:slug;not null;uniqueIndex"`
	Description      *string   `json:"description,omitempty" gorm:"column:description"`
	ImageURL         *string   `json:"image_url,omitempty" gorm:"column:image_url"`
	SortOrder        int       `json:"sort_order" gorm:"column:sort_order;default:0"`
	IsActive         bool      `json:"is_active" gorm:"column:is_active;default:true"`
	ProductCount     int       `json:"product_count,omitempty" gorm:"-"`
	SubcategoryCount int       `json:"subcategory_count,omitempty" gorm:"-"`
	CreatedAt        time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"column:updated_at"`
}
