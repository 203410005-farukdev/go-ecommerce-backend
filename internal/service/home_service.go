package service

import (
	"context"

	"github.com/Farukcoder/eCommerce-go/backend/internal/domain"
	"github.com/Farukcoder/eCommerce-go/backend/internal/dto"
	"github.com/Farukcoder/eCommerce-go/backend/internal/repository"
)

type HomeService struct {
	settingRepo  *repository.SettingRepository
	categoryRepo *repository.CategoryRepository
	productRepo  *repository.ProductRepository
}

func NewHomeService(settingRepo *repository.SettingRepository, categoryRepo *repository.CategoryRepository, productRepo *repository.ProductRepository) *HomeService {
	return &HomeService{
		settingRepo:  settingRepo,
		categoryRepo: categoryRepo,
		productRepo:  productRepo,
	}
}

func (s *HomeService) GetHomePageData(ctx context.Context) (*dto.HomePageData, error) {
	settings, err := s.settingRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	settingsMap := make(map[string]string, len(settings))
	for _, setting := range settings {
		settingsMap[setting.Key] = setting.Value
	}

	categories, err := s.categoryRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	products, err := s.productRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	activeCategories := make([]domain.Category, 0, len(categories))
	for _, category := range categories {
		if category.IsActive {
			activeCategories = append(activeCategories, category)
		}
	}

	activeProducts := make([]domain.Product, 0, len(products))
	for _, product := range products {
		if product.IsActive {
			activeProducts = append(activeProducts, product)
		}
	}

	featuredProducts := make([]domain.Product, 0, 8)
	newArrivals := make([]domain.Product, 0, 8)
	for _, product := range activeProducts {
		if product.IsFeatured && len(featuredProducts) < 8 {
			featuredProducts = append(featuredProducts, product)
		}
		if len(newArrivals) < 8 {
			newArrivals = append(newArrivals, product)
		}
	}

	return &dto.HomePageData{
		Settings:         settingsMap,
		Categories:       activeCategories,
		Products:         activeProducts,
		FeaturedProducts: featuredProducts,
		NewArrivals:      newArrivals,
	}, nil
}
