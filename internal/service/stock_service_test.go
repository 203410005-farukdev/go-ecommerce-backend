package service

import (
	"context"
	"testing"

	"github.com/Farukcoder/eCommerce-go/backend/internal/dto"
)

func TestStockService_CreateStockMovement_ValidatesChange(t *testing.T) {
	svc := &StockService{}

	_, err := svc.CreateStockMovement(context.Background(), dto.CreateStockMovementInput{
		ProductID: 1,
		Change:    0,
		Reason:    "restock",
	})
	if err == nil {
		t.Fatal("expected validation error when stock change is zero")
	}
}
