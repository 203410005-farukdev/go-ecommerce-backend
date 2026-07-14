package dto

type CreateProductInput struct {
	Name               string   `json:"name" form:"name"`
	Slug               string   `json:"slug,omitempty" form:"slug,omitempty"`
	SKU                *string  `json:"sku,omitempty" form:"sku,omitempty"`
	Description        *string  `json:"description,omitempty" form:"description,omitempty"`
	Price              float64  `json:"price" form:"price"`
	CompareAtPrice     *float64 `json:"compare_at_price,omitempty" form:"compare_at_price,omitempty"`
	CategoryID         int64    `json:"category_id" form:"category_id"`
	SubcategoryID      *int64   `json:"subcategory_id,omitempty" form:"subcategory_id,omitempty"`
	ImageURL           *string  `json:"image_url,omitempty" form:"image_url,omitempty"`
	GalleryURLs        []string `json:"gallery_urls,omitempty" form:"gallery_urls,omitempty"`
	RemovedGalleryURLs []string `json:"removed_gallery_urls,omitempty" form:"removed_gallery_urls,omitempty"`
	IsActive           bool     `json:"is_active" form:"is_active"`
	IsFeatured         bool     `json:"is_featured" form:"is_featured"`
}

type UpdateProductInput struct {
	Name               string   `json:"name" form:"name"`
	Slug               string   `json:"slug,omitempty" form:"slug,omitempty"`
	SKU                *string  `json:"sku,omitempty" form:"sku,omitempty"`
	Description        *string  `json:"description,omitempty" form:"description,omitempty"`
	Price              float64  `json:"price" form:"price"`
	CompareAtPrice     *float64 `json:"compare_at_price,omitempty" form:"compare_at_price,omitempty"`
	CategoryID         int64    `json:"category_id" form:"category_id"`
	SubcategoryID      *int64   `json:"subcategory_id,omitempty" form:"subcategory_id,omitempty"`
	ImageURL           *string  `json:"image_url,omitempty" form:"image_url,omitempty"`
	GalleryURLs        []string `json:"gallery_urls,omitempty" form:"gallery_urls,omitempty"`
	RemovedGalleryURLs []string `json:"removed_gallery_urls,omitempty" form:"removed_gallery_urls,omitempty"`
	IsActive           bool     `json:"is_active" form:"is_active"`
	IsFeatured         bool     `json:"is_featured" form:"is_featured"`
}

type CreateProductVariantInput struct {
	Name              string   `json:"name"`
	SKU               *string  `json:"sku,omitempty"`
	Price             float64  `json:"price"`
	CompareAtPrice    *float64 `json:"compare_at_price,omitempty"`
	Quantity          int      `json:"quantity"`
	LowStockThreshold int      `json:"low_stock_threshold"`
	IsActive          bool     `json:"is_active"`
	SortOrder         int      `json:"sort_order"`
}

type UpdateProductVariantInput struct {
	Name              string   `json:"name"`
	SKU               *string  `json:"sku,omitempty"`
	Price             float64  `json:"price"`
	CompareAtPrice    *float64 `json:"compare_at_price,omitempty"`
	Quantity          int      `json:"quantity"`
	LowStockThreshold int      `json:"low_stock_threshold"`
	IsActive          bool     `json:"is_active"`
	SortOrder         int      `json:"sort_order"`
}
