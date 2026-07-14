package service

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Farukcoder/eCommerce-go/backend/internal/domain"
	"github.com/Farukcoder/eCommerce-go/backend/internal/dto"
	"github.com/Farukcoder/eCommerce-go/backend/internal/repository"
)

type StockService struct {
	repo *repository.StockRepository
}

func NewStockService(repo *repository.StockRepository) *StockService {
	return &StockService{repo: repo}
}

func (s *StockService) ListStock(ctx context.Context) ([]domain.Stock, error) {
	return s.repo.List(ctx)
}

func (s *StockService) GetStock(ctx context.Context, stockID string) (*domain.Stock, error) {
	id, err := strconv.ParseInt(stockID, 10, 64)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, id)
}

func (s *StockService) UpdateStock(ctx context.Context, stockID string, input dto.UpdateStockInput) (*domain.Stock, error) {
	id, err := strconv.ParseInt(stockID, 10, 64)
	if err != nil {
		return nil, err
	}
	if input.Quantity < 0 {
		return nil, errors.New("stock quantity cannot be negative")
	}
	if input.LowStockThreshold < 0 {
		return nil, errors.New("low stock threshold cannot be negative")
	}
	var sku *string
	if input.SKU != nil {
		s := strings.TrimSpace(*input.SKU)
		sku = &s
	}
	var location *string
	if input.Location != nil {
		loc := strings.TrimSpace(*input.Location)
		location = &loc
	}
	return s.repo.Update(ctx, id, input.Quantity, input.LowStockThreshold, sku, location)
}

func (s *StockService) ListStockMovements(ctx context.Context) ([]domain.StockMovement, error) {
	return s.repo.ListMovements(ctx)
}

func (s *StockService) CreateStockMovement(ctx context.Context, input dto.CreateStockMovementInput) (*domain.StockMovement, error) {
	if input.ProductID <= 0 {
		return nil, errors.New("product id is required")
	}
	if input.Change == 0 {
		return nil, errors.New("stock change must be non-zero")
	}
	if strings.TrimSpace(input.Reason) == "" {
		return nil, errors.New("stock movement reason is required")
	}

	if _, err := s.repo.GetByProductID(ctx, input.ProductID); err != nil {
		if _, createErr := s.repo.CreateIfMissing(ctx, input.ProductID); createErr != nil {
			return nil, createErr
		}
	}

	if _, err := s.repo.ApplyMovement(ctx, input.ProductID, input.Change); err != nil {
		return nil, err
	}

	movement := domain.StockMovement{
		ProductID: input.ProductID,
		Change:    input.Change,
		Reason:    strings.TrimSpace(input.Reason),
		Note:      input.Note,
		CreatedAt: time.Now().UTC(),
	}

	return s.repo.CreateMovement(ctx, movement)
}
