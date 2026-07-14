package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Farukcoder/eCommerce-go/backend/internal/domain"
	"gorm.io/gorm"
)

type StockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) *StockRepository {
	return &StockRepository{db: db}
}

func (r *StockRepository) List(ctx context.Context) ([]domain.Stock, error) {
	var stocks []domain.Stock
	err := r.db.WithContext(ctx).
		Model(&domain.Stock{}).
		Order("updated_at desc").
		Find(&stocks).Error
	return stocks, err
}

func (r *StockRepository) GetByID(ctx context.Context, id int64) (*domain.Stock, error) {
	var stock domain.Stock
	err := r.db.WithContext(ctx).
		Model(&domain.Stock{}).
		Where("id = ?", id).
		Take(&stock).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("stock not found")
		}
		return nil, err
	}
	return &stock, nil
}

func (r *StockRepository) GetByProductID(ctx context.Context, productID int64) (*domain.Stock, error) {
	var stock domain.Stock
	err := r.db.WithContext(ctx).
		Model(&domain.Stock{}).
		Where("product_id = ?", productID).
		Take(&stock).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("stock not found")
		}
		return nil, err
	}
	return &stock, nil
}

func (r *StockRepository) Update(ctx context.Context, id int64, quantity int, lowStockThreshold int, sku *string, location *string) (*domain.Stock, error) {
	now := time.Now().UTC()
	result := r.db.WithContext(ctx).Table("stock").Where("id = ?", id).Updates(map[string]any{
		"quantity":            quantity,
		"low_stock_threshold": lowStockThreshold,
		"sku":                 sku,
		"location":            location,
		"updated_at":          now,
	})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("stock not found")
	}
	return r.GetByID(ctx, id)
}

func (r *StockRepository) CreateIfMissing(ctx context.Context, productID int64) (*domain.Stock, error) {
	var stock domain.Stock
	err := r.db.WithContext(ctx).Model(&domain.Stock{}).Where("product_id = ?", productID).Take(&stock).Error
	if err == nil {
		return &stock, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	stock = domain.Stock{
		ProductID:         productID,
		Quantity:          0,
		LowStockThreshold: 5,
		UpdatedAt:         time.Now().UTC(),
	}
	if err := r.db.WithContext(ctx).Table("stock").Create(&stock).Error; err != nil {
		return nil, err
	}
	return &stock, nil
}

func (r *StockRepository) ApplyMovement(ctx context.Context, productID int64, change int) (*domain.Stock, error) {
	var stock domain.Stock
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&domain.Stock{}).Where("product_id = ?", productID).Take(&stock).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				stock = domain.Stock{
					ProductID:         productID,
					Quantity:          0,
					LowStockThreshold: 5,
					UpdatedAt:         time.Now().UTC(),
				}
				if err := tx.Table("stock").Create(&stock).Error; err != nil {
					return err
				}
				return nil
			}
			return err
		}

		newQuantity := stock.Quantity + change
		if newQuantity < 0 {
			return errors.New("insufficient stock for product")
		}

		stock.Quantity = newQuantity
		if change > 0 {
			restockedAt := time.Now().UTC()
			stock.LastRestockedAt = &restockedAt
		}
		stock.UpdatedAt = time.Now().UTC()
		result := tx.Table("stock").Where("id = ?", stock.ID).Updates(map[string]any{
			"quantity":            stock.Quantity,
			"last_restocked_at":  stock.LastRestockedAt,
			"updated_at":          stock.UpdatedAt,
		})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("stock update failed")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &stock, nil
}

func (r *StockRepository) CreateMovement(ctx context.Context, movement domain.StockMovement) (*domain.StockMovement, error) {
	movement.CreatedAt = time.Now().UTC()
	if err := r.db.WithContext(ctx).Table("stock_movements").Create(&movement).Error; err != nil {
		return nil, err
	}
	return &movement, nil
}

func (r *StockRepository) ListMovements(ctx context.Context) ([]domain.StockMovement, error) {
	var movements []domain.StockMovement
	err := r.db.WithContext(ctx).
		Model(&domain.StockMovement{}).
		Order("created_at desc").
		Find(&movements).Error
	return movements, err
}
