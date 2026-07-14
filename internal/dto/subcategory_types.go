package dto

// CreateSubcategoryInput is the request body for creating a subcategory.
type CreateSubcategoryInput struct {
	Name        string  `json:"name" form:"name"`
	Slug        string  `json:"slug,omitempty" form:"slug,omitempty"`
	CategoryID  int64   `json:"category_id" form:"category_id"`
	Description *string `json:"description,omitempty" form:"description,omitempty"`
	ImageURL    *string `json:"image_url,omitempty" form:"image_url,omitempty"`
	SortOrder   int     `json:"sort_order" form:"sort_order"`
	IsActive    bool    `json:"is_active" form:"is_active"`
}

// UpdateSubcategoryInput is the request body for updating a subcategory.
type UpdateSubcategoryInput struct {
	Name        string  `json:"name" form:"name"`
	Slug        string  `json:"slug,omitempty" form:"slug,omitempty"`
	CategoryID  int64   `json:"category_id" form:"category_id"`
	Description *string `json:"description,omitempty" form:"description,omitempty"`
	ImageURL    *string `json:"image_url,omitempty" form:"image_url,omitempty"`
	SortOrder   int     `json:"sort_order" form:"sort_order"`
	IsActive    bool    `json:"is_active" form:"is_active"`
}
