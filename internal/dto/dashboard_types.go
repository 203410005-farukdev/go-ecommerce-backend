package dto

import (
	"time"

	"github.com/Farukcoder/eCommerce-go/backend/internal/domain"
)

type CategoryCount struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type LowStockItem struct {
	P domain.Product `json:"p"`
	S domain.Stock   `json:"s"`
}

type TopProduct struct {
	P     domain.Product `json:"p"`
	Qty   int            `json:"qty"`
	Value float64        `json:"value"`
}

type ActivityProduct struct {
	Name string `json:"name"`
}

type ActivityItem struct {
	ID        int64            `json:"id"`
	ProductID int64            `json:"product_id"`
	Change    int              `json:"change"`
	Reason    string           `json:"reason"`
	Note      *string          `json:"note,omitempty"`
	CreatedAt time.Time        `json:"created_at"`
	Products  *ActivityProduct `json:"products,omitempty"`
}

type DashboardStats struct {
	TotalProducts         int             `json:"products"`
	TotalCategories       int             `json:"categories"`
	TotalVariants         int             `json:"variants"`
	TotalStockUnits       int             `json:"totalStockUnits"`
	LowStockCount         int             `json:"lowStock"`
	OutOfStockCount       int             `json:"outOfStock"`
	ActiveProductsCount   int             `json:"activeProducts"`
	FeaturedProductsCount int             `json:"featured"`
	InventoryValue        float64         `json:"inventoryValue"`
	TotalStockMovements   int             `json:"totalStockMovements"`
	ProductTrend          string          `json:"productTrend"`
	MovementTrend         string          `json:"movementTrend"`
	CategoryDistribution  []CategoryCount `json:"categoryDistribution"`
	LowStockItems         []LowStockItem  `json:"lowStockItems"`
	TopProducts           []TopProduct    `json:"topProducts"`
	RecentActivities      []ActivityItem  `json:"recentActivities"`
}
