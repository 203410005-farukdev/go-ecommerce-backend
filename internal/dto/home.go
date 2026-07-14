package dto

import "github.com/Farukcoder/eCommerce-go/backend/internal/domain"

// HomePageData is the public payload for the storefront landing page.
type HomePageData struct {
	Settings         map[string]string `json:"settings"`
	Categories       []domain.Category `json:"categories"`
	Products         []domain.Product  `json:"products"`
	FeaturedProducts []domain.Product  `json:"featured_products"`
	NewArrivals      []domain.Product  `json:"new_arrivals"`
}
