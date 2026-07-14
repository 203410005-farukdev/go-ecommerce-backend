package service

import (
	"testing"

	"github.com/Farukcoder/eCommerce-go/backend/internal/domain"
)

func TestNormalizeCategory_GeneratesSlugAndTrimsValues(t *testing.T) {
	description := "  Summer Collection  "
	imageURL := "  https://example.com/image.png  "
	category := normalizeCategory(domain.Category{
		Name:        "  Men Clothing  ",
		Slug:        "",
		Description: &description,
		ImageURL:    &imageURL,
		SortOrder:   3,
		IsActive:    false,
	})

	if category.Name != "Men Clothing" {
		t.Fatalf("expected trimmed name, got %q", category.Name)
	}
	if category.Slug != "men-clothing" {
		t.Fatalf("expected generated slug, got %q", category.Slug)
	}
	if category.Description == nil || *category.Description != "Summer Collection" {
		t.Fatalf("expected trimmed description, got %#v", category.Description)
	}
	if category.ImageURL == nil || *category.ImageURL != "https://example.com/image.png" {
		t.Fatalf("expected trimmed image url, got %#v", category.ImageURL)
	}
	if category.SortOrder != 3 {
		t.Fatalf("expected sort order to be preserved")
	}
	if category.IsActive {
		t.Fatalf("expected is_active to be preserved as false")
	}
}
