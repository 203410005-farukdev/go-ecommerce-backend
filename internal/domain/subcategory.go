package domain

import "time"

// Subcategory is the DB entity for product subcategories.
type Subcategory struct {
	ID          int64     `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	CategoryID  int64     `json:"category_id" gorm:"column:category_id;not null"`
	Name        string    `json:"name" gorm:"column:name;not null"`
	Slug        string    `json:"slug" gorm:"column:slug;not null"`
	Description *string   `json:"description,omitempty" gorm:"column:description"`
	ImageURL    *string   `json:"image_url,omitempty" gorm:"column:image_url"`
	SortOrder   int       `json:"sort_order" gorm:"column:sort_order;default:0"`
	IsActive    bool      `json:"is_active" gorm:"column:is_active;default:true"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}
