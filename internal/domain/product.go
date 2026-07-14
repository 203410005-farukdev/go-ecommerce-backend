package domain

import "time"

type Product struct {
	ID              int64            `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Name            string           `json:"name" gorm:"column:name;not null"`
	Slug            string           `json:"slug" gorm:"column:slug;not null;uniqueIndex"`
	SKU             *string          `json:"sku,omitempty" gorm:"column:sku;uniqueIndex"`
	Description     *string          `json:"description,omitempty" gorm:"column:description"`
	Price           float64          `json:"price" gorm:"column:price;type:numeric(10,2);not null;default:0"`
	CompareAtPrice  *float64         `json:"compare_at_price,omitempty" gorm:"column:compare_at_price;type:numeric(10,2)"`
	CategoryID      int64            `json:"category_id" gorm:"column:category_id;not null"`
	SubcategoryID   *int64           `json:"subcategory_id,omitempty" gorm:"column:subcategory_id"`
	ImageURL        *string          `json:"image_url,omitempty" gorm:"column:image_url"`
	IsActive        bool             `json:"is_active" gorm:"column:is_active;default:true"`
	IsFeatured      bool             `json:"is_featured" gorm:"column:is_featured;default:false"`
	Categories      *Category        `json:"categories,omitempty" gorm:"foreignKey:CategoryID;references:ID"`
	Subcategories   *Subcategory     `json:"subcategories,omitempty" gorm:"foreignKey:SubcategoryID;references:ID"`
	Stock           *Stock           `json:"stock,omitempty" gorm:"foreignKey:ProductID;references:ID"`
	ProductVariants []ProductVariant `json:"product_variants,omitempty" gorm:"foreignKey:ProductID;references:ID"`
	ProductImages   []ProductImage   `json:"product_images,omitempty" gorm:"foreignKey:ProductID;references:ID"`
	CreatedAt       time.Time        `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       time.Time        `json:"updated_at" gorm:"column:updated_at"`
}

type ProductVariant struct {
	ID                int64     `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	ProductID         int64     `json:"product_id" gorm:"column:product_id;not null"`
	Name              string    `json:"name" gorm:"column:name;not null"`
	SKU               *string   `json:"sku,omitempty" gorm:"column:sku"`
	Price             float64   `json:"price" gorm:"column:price;type:numeric(10,2);not null;default:0"`
	CompareAtPrice    *float64  `json:"compare_at_price,omitempty" gorm:"column:compare_at_price;type:numeric(10,2)"`
	Quantity          int       `json:"quantity" gorm:"column:quantity;not null;default:0"`
	LowStockThreshold int       `json:"low_stock_threshold" gorm:"column:low_stock_threshold;not null;default:5"`
	IsActive          bool      `json:"is_active" gorm:"column:is_active;default:true"`
	SortOrder         int       `json:"sort_order" gorm:"column:sort_order;not null;default:0"`
	CreatedAt         time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type ProductImage struct {
	ID        int64     `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	ProductID int64     `json:"product_id" gorm:"column:product_id;not null;index"`
	URL       string    `json:"url" gorm:"column:url;not null"`
	SortOrder int       `json:"sort_order" gorm:"column:sort_order;default:0;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type Stock struct {
	ID                int64      `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	ProductID         int64      `json:"product_id" gorm:"column:product_id;not null;uniqueIndex"`
	Quantity          int        `json:"quantity" gorm:"column:quantity;not null;default:0"`
	LowStockThreshold int        `json:"low_stock_threshold" gorm:"column:low_stock_threshold;not null;default:5"`
	SKU               *string    `json:"sku,omitempty" gorm:"column:sku"`
	Location          *string    `json:"location,omitempty" gorm:"column:location"`
	LastRestockedAt   *time.Time `json:"last_restocked_at,omitempty" gorm:"column:last_restocked_at"`
	UpdatedAt         time.Time  `json:"updated_at" gorm:"column:updated_at"`
}

func (Stock) TableName() string {
	return "stock"
}

type StockMovement struct {
	ID        int64     `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	ProductID int64     `json:"product_id" gorm:"column:product_id;not null"`
	Change    int       `json:"change" gorm:"column:change;not null"`
	Reason    string    `json:"reason" gorm:"column:reason;not null"`
	Note      *string   `json:"note,omitempty" gorm:"column:note"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}
