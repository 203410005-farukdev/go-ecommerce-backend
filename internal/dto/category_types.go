package dto

// CreateCategoryInput is the request body for creating a category.
type CreateCategoryInput struct {
	Name        string  `json:"name" form:"name"`
	Slug        string  `json:"slug,omitempty" form:"slug,omitempty"`
	Description *string `json:"description,omitempty" form:"description,omitempty"`
	ImageURL    *string `json:"image_url,omitempty" form:"image_url,omitempty"`
	SortOrder   int     `json:"sort_order" form:"sort_order"`
	IsActive    bool    `json:"is_active" form:"is_active"`
}

// UpdateCategoryInput is the request body for updating a category.
type UpdateCategoryInput struct {
	Name        string  `json:"name" form:"name"`
	Slug        string  `json:"slug,omitempty" form:"slug,omitempty"`
	Description *string `json:"description,omitempty" form:"description,omitempty"`
	ImageURL    *string `json:"image_url,omitempty" form:"image_url,omitempty"`
	SortOrder   int     `json:"sort_order" form:"sort_order"`
	IsActive    bool    `json:"is_active" form:"is_active"`
}
