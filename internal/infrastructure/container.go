package infrastructure

import (
	"github.com/besyuzkirk/feature-flag-management/config"
	"github.com/besyuzkirk/feature-flag-management/internal/domain/repositories"
	"github.com/besyuzkirk/feature-flag-management/internal/domain/services"
	"gorm.io/gorm"
)

type Container struct {
	DB                     *gorm.DB
	FeatureFlagService     services.FeatureFlagService
	SegmentService         services.SegmentService
	RolloutStrategyService services.RolloutStrategyService
}

func NewContainer() (*Container, error) {
	cfg := config.LoadDBConfig()
	db, err := ConnectDB(cfg)
	if err != nil {
		return nil, err
	}

	flagRepo := repositories.NewFeatureFlagRepository(db)
	flagService := services.NewFeatureFlagService(flagRepo)

	segmentRepo := repositories.NewSegmentRepository(db)
	segmentService := services.NewSegmentService(segmentRepo)

	rolloutRepo := repositories.NewRolloutStrategyRepository(db)
	rolloutService := services.NewRolloutStrategyService(rolloutRepo)

	return &Container{
		DB:                     db, // Veritabanı bağlantısını ekliyoruz
		FeatureFlagService:     flagService,
		SegmentService:         segmentService,
		RolloutStrategyService: rolloutService,
	}, nil
}
