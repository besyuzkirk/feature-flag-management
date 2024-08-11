package repositories

import (
	"errors"

	"github.com/besyuzkirk/feature-flag-management/internal/domain/entities"
	"gorm.io/gorm"
)

type RolloutStrategyRepository interface {
	CreateRolloutStrategy(strategy *entities.RolloutStrategy) error
	UpdateRolloutStrategy(strategy *entities.RolloutStrategy) error
	DeleteRolloutStrategy(id uint) error
	GetRolloutStrategy(id uint) (*entities.RolloutStrategy, error)
	ListRolloutStrategiesByFlag(flagID uint) ([]*entities.RolloutStrategy, error)
}

type rolloutStrategyRepository struct {
	db *gorm.DB
}

func NewRolloutStrategyRepository(db *gorm.DB) RolloutStrategyRepository {
	return &rolloutStrategyRepository{db}
}

func (r *rolloutStrategyRepository) CreateRolloutStrategy(strategy *entities.RolloutStrategy) error {
	return r.db.Create(strategy).Error
}

func (r *rolloutStrategyRepository) UpdateRolloutStrategy(strategy *entities.RolloutStrategy) error {
	if r.db.Model(strategy).Where("id = ?", strategy.ID).Updates(strategy).RowsAffected == 0 {
		return errors.New("no rows affected or rollout strategy not found")
	}
	return nil
}

func (r *rolloutStrategyRepository) DeleteRolloutStrategy(id uint) error {
	if r.db.Delete(&entities.RolloutStrategy{}, id).RowsAffected == 0 {
		return errors.New("no rows affected or rollout strategy not found")
	}
	return nil
}

func (r *rolloutStrategyRepository) GetRolloutStrategy(id uint) (*entities.RolloutStrategy, error) {
	var strategy entities.RolloutStrategy
	if err := r.db.First(&strategy, id).Error; err != nil {
		return nil, err
	}
	return &strategy, nil
}

func (r *rolloutStrategyRepository) ListRolloutStrategiesByFlag(flagID uint) ([]*entities.RolloutStrategy, error) {
	var strategies []*entities.RolloutStrategy
	if err := r.db.Where("feature_flag_id = ?", flagID).Find(&strategies).Error; err != nil {
		return nil, err
	}
	return strategies, nil
}
