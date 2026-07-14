package service

import (
	"context"

	"github.com/Farukcoder/eCommerce-go/backend/internal/repository"
)

type SettingService struct {
	repo *repository.SettingRepository
}

func NewSettingService(repo *repository.SettingRepository) *SettingService {
	return &SettingService{repo: repo}
}

func (s *SettingService) GetSettings(ctx context.Context) (map[string]string, error) {
	settings, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	res := make(map[string]string)
	for _, setting := range settings {
		res[setting.Key] = setting.Value
	}
	return res, nil
}

func (s *SettingService) UpdateSettings(ctx context.Context, updates map[string]string) error {
	for k, v := range updates {
		if err := s.repo.Save(ctx, k, v); err != nil {
			return err
		}
	}
	return nil
}
