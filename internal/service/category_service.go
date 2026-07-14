package service

import (
	"context"
	"strconv"
	"strings"
	"unicode"

	"github.com/Farukcoder/eCommerce-go/backend/internal/domain"
	"github.com/Farukcoder/eCommerce-go/backend/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) ListCategories(ctx context.Context) ([]domain.Category, error) {
	categories, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return s.populateCounts(ctx, categories)
}

func (s *CategoryService) GetCategory(ctx context.Context, id string) (*domain.Category, error) {
	categoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	category, err := s.repo.GetByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}
	return s.populateCount(ctx, category)
}

func (s *CategoryService) CreateCategory(ctx context.Context, input domain.Category) (*domain.Category, error) {
	normalized := normalizeCategory(input)
	return s.repo.Create(ctx, normalized)
}

func (s *CategoryService) UpdateCategory(ctx context.Context, id string, input domain.Category) (*domain.Category, error) {
	categoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	normalized := normalizeCategory(input)
	return s.repo.Update(ctx, categoryID, normalized)
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id string) error {
	categoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, categoryID)
}

func (s *CategoryService) populateCounts(ctx context.Context, categories []domain.Category) ([]domain.Category, error) {
	for i := range categories {
		if populated, err := s.populateCount(ctx, &categories[i]); err != nil {
			return nil, err
		} else {
			categories[i] = *populated
		}
	}
	return categories, nil
}

func (s *CategoryService) populateCount(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	productCount, err := s.repo.CountProductsByCategory(ctx, category.ID)
	if err != nil {
		return nil, err
	}
	subcategoryCount, err := s.repo.CountSubcategoriesByCategory(ctx, category.ID)
	if err != nil {
		return nil, err
	}
	category.ProductCount = int(productCount)
	category.SubcategoryCount = int(subcategoryCount)
	return category, nil
}

func normalizeCategory(input domain.Category) domain.Category {
	name := strings.TrimSpace(input.Name)
	slug := strings.TrimSpace(input.Slug)
	if slug == "" {
		slug = slugifyCategory(name)
	}

	category := domain.Category{
		Name:      name,
		Slug:      slug,
		SortOrder: input.SortOrder,
		IsActive:  input.IsActive,
	}

	if input.Description != nil {
		description := strings.TrimSpace(*input.Description)
		if description != "" {
			category.Description = &description
		} else {
			category.Description = nil
		}
	}

	if input.ImageURL != nil {
		imageURL := strings.TrimSpace(*input.ImageURL)
		if imageURL != "" {
			category.ImageURL = &imageURL
		} else {
			category.ImageURL = nil
		}
	}

	return category
}

func slugifyCategory(input string) string {
	// Preserve Unicode letters and numbers (e.g., Bangla) while producing
	// a URL-friendly slug. Replace spaces/underscores with hyphens,
	// remove other punctuation, collapse duplicate hyphens and trim.
	s := strings.ToLower(strings.TrimSpace(input))

	// replace spaces and underscores with hyphen
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, "_", "-")

	var b strings.Builder
	prevHyphen := false
	for _, r := range s {
		if r == '-' {
			if !prevHyphen {
				b.WriteRune('-')
				prevHyphen = true
			}
			continue
		}
		// Keep letters and numbers from any script
		if strings.ContainsRune("abcdefghijklmnopqrstuvwxyz0123456789", r) ||
			unicode.IsLetter(r) || unicode.IsNumber(r) {
			b.WriteRune(r)
			prevHyphen = false
			continue
		}
		// drop other characters
		// (this will remove punctuation like commas, slashes, etc.)
	}

	out := b.String()
	out = strings.Trim(out, "-")
	return out
}
