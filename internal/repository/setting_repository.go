package repository

import (
	"context"

	"github.com/Farukcoder/eCommerce-go/backend/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SettingRepository struct {
	db *gorm.DB
}

func NewSettingRepository(db *gorm.DB) *SettingRepository {
	return &SettingRepository{db: db}
}

func (r *SettingRepository) FindAll(ctx context.Context) ([]domain.Setting, error) {
	var settings []domain.Setting
	err := r.db.WithContext(ctx).Table("settings").Find(&settings).Error
	return settings, err
}

func (r *SettingRepository) Save(ctx context.Context, key, value string) error {
	setting := domain.Setting{
		Key:   key,
		Value: value,
	}
	return r.db.WithContext(ctx).Table("settings").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"value": value, "updated_at": gorm.Expr("NOW()")}),
	}).Create(&setting).Error
}
