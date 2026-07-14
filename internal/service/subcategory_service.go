package service

import (
	"context"
	"strconv"
	"strings"
	"unicode"

	"github.com/Farukcoder/eCommerce-go/backend/internal/domain"
	"github.com/Farukcoder/eCommerce-go/backend/internal/repository"
)

type SubcategoryService struct {
	repo *repository.SubcategoryRepository
}

func NewSubcategoryService(repo *repository.SubcategoryRepository) *SubcategoryService {
	return &SubcategoryService{repo: repo}
}

func (s *SubcategoryService) ListSubcategories(ctx context.Context) ([]domain.Subcategory, error) {
	return s.repo.List(ctx)
}

func (s *SubcategoryService) GetSubcategory(ctx context.Context, id string) (*domain.Subcategory, error) {
	subcategoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, subcategoryID)
}

func (s *SubcategoryService) CreateSubcategory(ctx context.Context, input domain.Subcategory) (*domain.Subcategory, error) {
	normalized := normalizeSubcategory(input)
	return s.repo.Create(ctx, normalized)
}

func (s *SubcategoryService) UpdateSubcategory(ctx context.Context, id string, input domain.Subcategory) (*domain.Subcategory, error) {
	subcategoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	normalized := normalizeSubcategory(input)
	return s.repo.Update(ctx, subcategoryID, normalized)
}

func (s *SubcategoryService) DeleteSubcategory(ctx context.Context, id string) error {
	subcategoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, subcategoryID)
}

func normalizeSubcategory(input domain.Subcategory) domain.Subcategory {
	name := strings.TrimSpace(input.Name)
	slug := strings.TrimSpace(input.Slug)
	if slug == "" {
		slug = slugifySubcategory(name)
	}

	subcategory := domain.Subcategory{
		CategoryID: input.CategoryID,
		Name:       name,
		Slug:       slug,
		SortOrder:  input.SortOrder,
		IsActive:   input.IsActive,
	}

	if input.Description != nil {
		description := strings.TrimSpace(*input.Description)
		if description != "" {
			subcategory.Description = &description
		} else {
			subcategory.Description = nil
		}
	}

	if input.ImageURL != nil {
		imageURL := strings.TrimSpace(*input.ImageURL)
		if imageURL != "" {
			subcategory.ImageURL = &imageURL
		} else {
			subcategory.ImageURL = nil
		}
	}

	return subcategory
}

func slugifySubcategory(input string) string {
	s := strings.ToLower(strings.TrimSpace(input))
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
		if strings.ContainsRune("abcdefghijklmnopqrstuvwxyz0123456789", r) || unicode.IsLetter(r) || unicode.IsNumber(r) {
			b.WriteRune(r)
			prevHyphen = false
			continue
		}
	}

	out := b.String()
	out = strings.Trim(out, "-")
	return out
}
