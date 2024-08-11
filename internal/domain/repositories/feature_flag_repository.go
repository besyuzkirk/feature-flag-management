package repositories

import (
	"errors"

	"github.com/besyuzkirk/feature-flag-management/internal/domain/entities"
	"gorm.io/gorm"
)

type FeatureFlagRepository interface {
	CreateFeatureFlag(flag *entities.FeatureFlag) error
	UpdateFeatureFlag(flag *entities.FeatureFlag) error
	DeleteFeatureFlag(id uint) error
	GetFeatureFlag(id uint) (*entities.FeatureFlag, error)
	ListFeatureFlags() ([]*entities.FeatureFlag, error)
	SaveFeatureFlagHistory(flag *entities.FeatureFlag) error
	GetLastFeatureFlagHistory(featureFlagID uint) (*entities.FeatureFlagHistory, error)
}

type featureFlagRepository struct {
	db *gorm.DB
}

func NewFeatureFlagRepository(db *gorm.DB) FeatureFlagRepository {
	return &featureFlagRepository{db}
}

func (r *featureFlagRepository) CreateFeatureFlag(flag *entities.FeatureFlag) error {
	return r.db.Create(flag).Error
}

func (r *featureFlagRepository) UpdateFeatureFlag(flag *entities.FeatureFlag) error {
	if r.db.Model(flag).Where("id = ?", flag.ID).Updates(flag).RowsAffected == 0 {
		return errors.New("no rows affected or feature flag not found")
	}
	return nil
}

func (r *featureFlagRepository) DeleteFeatureFlag(id uint) error {
	if r.db.Delete(&entities.FeatureFlag{}, id).RowsAffected == 0 {
		return errors.New("no rows affected or feature flag not found")
	}
	return nil
}

func (r *featureFlagRepository) GetFeatureFlag(id uint) (*entities.FeatureFlag, error) {
	var flag entities.FeatureFlag
	if err := r.db.First(&flag, id).Error; err != nil {
		return nil, err
	}
	return &flag, nil
}

func (r *featureFlagRepository) ListFeatureFlags() ([]*entities.FeatureFlag, error) {
	var flags []*entities.FeatureFlag
	if err := r.db.Find(&flags).Error; err != nil {
		return nil, err
	}
	return flags, nil
}

func (r *featureFlagRepository) SaveFeatureFlagHistory(flag *entities.FeatureFlag) error {
	history := &entities.FeatureFlagHistory{
		FeatureFlagID: flag.ID,
		Name:          flag.Name,
		Description:   flag.Description,
		IsActive:      flag.IsActive,
	}
	return r.db.Create(history).Error
}

func (r *featureFlagRepository) GetLastFeatureFlagHistory(featureFlagID uint) (*entities.FeatureFlagHistory, error) {
	var history entities.FeatureFlagHistory
	if err := r.db.Where("feature_flag_id = ?", featureFlagID).Order("created_at desc").First(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}
