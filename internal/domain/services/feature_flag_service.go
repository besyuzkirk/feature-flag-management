package services

import (
	"github.com/besyuzkirk/feature-flag-management/internal/domain/entities"
	"github.com/besyuzkirk/feature-flag-management/internal/domain/repositories"
)

type FeatureFlagService interface {
	CreateFeatureFlag(name, description string) (*entities.FeatureFlag, error)
	UpdateFeatureFlag(id uint, name, description string, isActive bool) (*entities.FeatureFlag, error)
	DeleteFeatureFlag(id uint) error
	GetFeatureFlag(id uint) (*entities.FeatureFlag, error)
	ListFeatureFlags() ([]*entities.FeatureFlag, error)
	RollbackFeatureFlag(id uint) (*entities.FeatureFlag, error)
}

type featureFlagService struct {
	repo repositories.FeatureFlagRepository
}

func NewFeatureFlagService(repo repositories.FeatureFlagRepository) FeatureFlagService {
	return &featureFlagService{repo}
}

func (s *featureFlagService) CreateFeatureFlag(name, description string) (*entities.FeatureFlag, error) {
	flag := &entities.FeatureFlag{
		Name:        name,
		Description: description,
		IsActive:    true,
	}

	if err := s.repo.CreateFeatureFlag(flag); err != nil {
		return nil, err
	}
	return flag, nil
}

func (s *featureFlagService) UpdateFeatureFlag(id uint, name, description string, isActive bool) (*entities.FeatureFlag, error) {
	flag, err := s.repo.GetFeatureFlag(id)
	if err != nil {
		return nil, err
	}

	flag.Name = name
	flag.Description = description
	flag.IsActive = isActive

	if err := s.repo.UpdateFeatureFlag(flag); err != nil {
		return nil, err
	}
	return flag, nil
}

func (s *featureFlagService) DeleteFeatureFlag(id uint) error {
	return s.repo.DeleteFeatureFlag(id)
}

func (s *featureFlagService) GetFeatureFlag(id uint) (*entities.FeatureFlag, error) {
	return s.repo.GetFeatureFlag(id)
}

func (s *featureFlagService) ListFeatureFlags() ([]*entities.FeatureFlag, error) {
	return s.repo.ListFeatureFlags()
}

func (s *featureFlagService) RollbackFeatureFlag(id uint) (*entities.FeatureFlag, error) {
	flag, err := s.repo.GetFeatureFlag(id)
	if err != nil {
		return nil, err
	}

	lastHistory, err := s.repo.GetLastFeatureFlagHistory(id)
	if err != nil {
		return nil, err
	}

	flag.Name = lastHistory.Name
	flag.Description = lastHistory.Description
	flag.IsActive = lastHistory.IsActive

	if err := s.repo.UpdateFeatureFlag(flag); err != nil {
		return nil, err
	}

	return flag, nil
}
