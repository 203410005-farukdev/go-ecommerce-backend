package dto

type UpdateStockInput struct {
	Quantity          int     `json:"quantity"`
	LowStockThreshold int     `json:"low_stock_threshold"`
	SKU               *string `json:"sku,omitempty"`
	Location          *string `json:"location,omitempty"`
}

type CreateStockMovementInput struct {
	ProductID int64   `json:"product_id"`
	Change    int     `json:"change"`
	Reason    string  `json:"reason"`
	Note      *string `json:"note,omitempty"`
}
