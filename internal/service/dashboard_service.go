package service

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/Farukcoder/eCommerce-go/backend/internal/domain"
	"github.com/Farukcoder/eCommerce-go/backend/internal/dto"
	"gorm.io/gorm"
)

type DashboardService struct {
	db *gorm.DB
}

func NewDashboardService(db *gorm.DB) *DashboardService {
	return &DashboardService{db: db}
}

type smWithProduct struct {
	domain.StockMovement
	ProductName string `gorm:"column:product_name"`
}

func (s *DashboardService) GetStats(ctx context.Context) (*dto.DashboardStats, error) {
	var dbProducts []domain.Product
	if err := s.db.WithContext(ctx).Preload("Stock").Preload("Categories").Find(&dbProducts).Error; err != nil {
		return nil, err
	}

	var dbCategories []domain.Category
	if err := s.db.WithContext(ctx).Find(&dbCategories).Error; err != nil {
		return nil, err
	}

	var dbVariants []domain.ProductVariant
	if err := s.db.WithContext(ctx).Find(&dbVariants).Error; err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	oneWeekAgo := now.AddDate(0, 0, -7)
	twoWeeksAgo := now.AddDate(0, 0, -14)

	var totalMovements int64
	if err := s.db.WithContext(ctx).Model(&domain.StockMovement{}).Count(&totalMovements).Error; err != nil {
		return nil, err
	}

	var thisWeekMovements int64
	if err := s.db.WithContext(ctx).Model(&domain.StockMovement{}).Where("created_at >= ?", oneWeekAgo).Count(&thisWeekMovements).Error; err != nil {
		return nil, err
	}

	var lastWeekMovements int64
	if err := s.db.WithContext(ctx).Model(&domain.StockMovement{}).Where("created_at >= ? AND created_at < ?", twoWeeksAgo, oneWeekAgo).Count(&lastWeekMovements).Error; err != nil {
		return nil, err
	}

	var dbMovements []smWithProduct
	err := s.db.WithContext(ctx).Table("stock_movements").
		Select("stock_movements.*, products.name as product_name").
		Joins("left join products on products.id = stock_movements.product_id").
		Order("created_at desc").
		Limit(10).
		Find(&dbMovements).Error
	if err != nil {
		return nil, err
	}

	// Computations
	totalStockUnits := 0
	lowStockCount := 0
	outOfStockCount := 0
	activeProducts := 0
	featuredProducts := 0
	var inventoryValue float64

	stockMap := make(map[int64]domain.Stock)
	for _, p := range dbProducts {
		if p.Stock != nil {
			stockMap[p.ID] = *p.Stock
			totalStockUnits += p.Stock.Quantity
			if p.Stock.Quantity <= 0 {
				outOfStockCount++
			} else if p.Stock.Quantity <= p.Stock.LowStockThreshold {
				lowStockCount++
			}
		} else {
			outOfStockCount++ // No stock record counts as out of stock
		}

		if p.IsActive {
			activeProducts++
		}
		if p.IsFeatured {
			featuredProducts++
		}
	}

	variantsByProduct := make(map[int64][]domain.ProductVariant)
	for _, v := range dbVariants {
		totalStockUnits += v.Quantity
		variantsByProduct[v.ProductID] = append(variantsByProduct[v.ProductID], v)
		if v.Quantity <= 0 {
			// Out of stock counts
		} else if v.Quantity <= v.LowStockThreshold {
			lowStockCount++
		}
	}

	for _, p := range dbProducts {
		pvs := variantsByProduct[p.ID]
		if len(pvs) > 0 {
			for _, v := range pvs {
				inventoryValue += v.Price * float64(v.Quantity)
			}
		} else if p.Stock != nil {
			inventoryValue += p.Price * float64(p.Stock.Quantity)
		}
	}

	// 1. Category Distribution
	catCounts := make(map[int64]int)
	for _, p := range dbProducts {
		catCounts[p.CategoryID]++
	}
	categoryDistribution := []dto.CategoryCount{}
	for _, c := range dbCategories {
		categoryDistribution = append(categoryDistribution, dto.CategoryCount{
			Name:  c.Name,
			Count: catCounts[c.ID],
		})
	}
	sort.Slice(categoryDistribution, func(i, j int) bool {
		return categoryDistribution[i].Count > categoryDistribution[j].Count
	})

	// 2. Low Stock Items
	lowStockItems := []dto.LowStockItem{}
	for _, p := range dbProducts {
		if p.Stock != nil && p.Stock.Quantity <= p.Stock.LowStockThreshold {
			lowStockItems = append(lowStockItems, dto.LowStockItem{
				P: p,
				S: *p.Stock,
			})
		}
	}
	sort.Slice(lowStockItems, func(i, j int) bool {
		return lowStockItems[i].S.Quantity < lowStockItems[j].S.Quantity
	})
	if len(lowStockItems) > 6 {
		lowStockItems = lowStockItems[:6]
	}

	// 3. Top Products by Value
	topProducts := []dto.TopProduct{}
	for _, p := range dbProducts {
		qty := 0
		if p.Stock != nil {
			qty = p.Stock.Quantity
		}
		value := p.Price * float64(qty)
		topProducts = append(topProducts, dto.TopProduct{
			P:     p,
			Qty:   qty,
			Value: value,
		})
	}
	sort.Slice(topProducts, func(i, j int) bool {
		return topProducts[i].Value > topProducts[j].Value
	})
	if len(topProducts) > 5 {
		topProducts = topProducts[:5]
	}

	// 4. Recent Activities
	recentActivities := []dto.ActivityItem{}
	for _, m := range dbMovements {
		note := m.Note
		recentActivities = append(recentActivities, dto.ActivityItem{
			ID:        m.ID,
			ProductID: m.ProductID,
			Change:    m.Change,
			Reason:    m.Reason,
			Note:      note,
			CreatedAt: m.CreatedAt,
			Products: &dto.ActivityProduct{
				Name: m.ProductName,
			},
		})
	}

	// 5. Product Growth Trend
	productsThisWeek := 0
	productsLastWeek := 0
	for _, p := range dbProducts {
		if p.CreatedAt.After(oneWeekAgo) {
			productsThisWeek++
		} else if p.CreatedAt.After(twoWeeksAgo) && p.CreatedAt.Before(oneWeekAgo) {
			productsLastWeek++
		}
	}
	productTrend := "0%"
	if productsLastWeek > 0 {
		change := (float64(productsThisWeek-productsLastWeek) / float64(productsLastWeek)) * 100
		productTrend = fmt.Sprintf("%+.0f%%", change)
	} else if productsThisWeek > 0 {
		productTrend = "+100%"
	}

	// 6. Movement Trend
	movementTrend := "0%"
	if lastWeekMovements > 0 {
		change := (float64(thisWeekMovements-lastWeekMovements) / float64(lastWeekMovements)) * 100
		movementTrend = fmt.Sprintf("%+.0f%%", change)
	} else if thisWeekMovements > 0 {
		movementTrend = "+100%"
	}

	return &dto.DashboardStats{
		TotalProducts:         len(dbProducts),
		TotalCategories:       len(dbCategories),
		TotalVariants:         len(dbVariants),
		TotalStockUnits:       totalStockUnits,
		LowStockCount:         lowStockCount,
		OutOfStockCount:       outOfStockCount,
		ActiveProductsCount:   activeProducts,
		FeaturedProductsCount: featuredProducts,
		InventoryValue:        inventoryValue,
		TotalStockMovements:   int(totalMovements),
		ProductTrend:          productTrend,
		MovementTrend:         movementTrend,
		CategoryDistribution:  categoryDistribution,
		LowStockItems:         lowStockItems,
		TopProducts:           topProducts,
		RecentActivities:      recentActivities,
	}, nil
}
