package services

import (
	"errors"

	"github.com/besyuzkirk/feature-flag-management/internal/domain/entities"
	"github.com/besyuzkirk/feature-flag-management/internal/domain/repositories"
)

type RolloutStrategyService interface {
	CreateRolloutStrategy(featureFlagID uint, percentage int, description string) (*entities.RolloutStrategy, error)
	UpdateRolloutStrategy(id uint, percentage int, description string) (*entities.RolloutStrategy, error)
	DeleteRolloutStrategy(id uint) error
	GetRolloutStrategy(id uint) (*entities.RolloutStrategy, error)
	ListRolloutStrategiesByFlag(flagID uint) ([]*entities.RolloutStrategy, error)
	TrackRolloutProgress(flagID uint) (int, error) // Rollout ilerlemesini izler
}

type rolloutStrategyService struct {
	repo repositories.RolloutStrategyRepository
}

func NewRolloutStrategyService(repo repositories.RolloutStrategyRepository) RolloutStrategyService {
	return &rolloutStrategyService{repo}
}

func (s *rolloutStrategyService) CreateRolloutStrategy(featureFlagID uint, percentage int, description string) (*entities.RolloutStrategy, error) {
	if percentage < 0 || percentage > 100 {
		return nil, errors.New("percentage must be between 0 and 100")
	}

	strategy := &entities.RolloutStrategy{
		FeatureFlagID: featureFlagID,
		Percentage:    percentage,
		Description:   description,
	}

	if err := s.repo.CreateRolloutStrategy(strategy); err != nil {
		return nil, err
	}
	return strategy, nil
}

func (s *rolloutStrategyService) UpdateRolloutStrategy(id uint, percentage int, description string) (*entities.RolloutStrategy, error) {
	if percentage < 0 || percentage > 100 {
		return nil, errors.New("percentage must be between 0 and 100")
	}

	strategy, err := s.repo.GetRolloutStrategy(id)
	if err != nil {
		return nil, err
	}

	strategy.Percentage = percentage
	strategy.Description = description

	if err := s.repo.UpdateRolloutStrategy(strategy); err != nil {
		return nil, err
	}
	return strategy, nil
}

func (s *rolloutStrategyService) DeleteRolloutStrategy(id uint) error {
	return s.repo.DeleteRolloutStrategy(id)
}

func (s *rolloutStrategyService) GetRolloutStrategy(id uint) (*entities.RolloutStrategy, error) {
	return s.repo.GetRolloutStrategy(id)
}

func (s *rolloutStrategyService) ListRolloutStrategiesByFlag(flagID uint) ([]*entities.RolloutStrategy, error) {
	return s.repo.ListRolloutStrategiesByFlag(flagID)
}

func (s *rolloutStrategyService) TrackRolloutProgress(flagID uint) (int, error) {
	strategies, err := s.repo.ListRolloutStrategiesByFlag(flagID)
	if err != nil {
		return 0, err
	}

	if len(strategies) == 0 {
		return 0, errors.New("no rollout strategies found for this flag")
	}

	latestStrategy := strategies[len(strategies)-1]
	return latestStrategy.Percentage, nil
}
