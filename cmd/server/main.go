package main

import (
	"log"

	"github.com/besyuzkirk/feature-flag-management/config"
	"github.com/besyuzkirk/feature-flag-management/internal/domain/repositories"
	"github.com/besyuzkirk/feature-flag-management/internal/domain/services"
	"github.com/besyuzkirk/feature-flag-management/internal/handlers"
	"github.com/besyuzkirk/feature-flag-management/internal/infrastructure"
	"github.com/besyuzkirk/feature-flag-management/internal/routes"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	// Config and DB setup
	cfg := config.LoadDBConfig()
	err := infrastructure.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Repository, Service, and Handler setup for Feature Flags
	flagRepo := repositories.NewFeatureFlagRepository(infrastructure.DB)
	flagService := services.NewFeatureFlagService(flagRepo)
	flagHandler := handlers.NewFeatureFlagHandler(flagService)

	// Repository, Service, and Handler setup for Segments
	segmentRepo := repositories.NewSegmentRepository(infrastructure.DB)
	segmentService := services.NewSegmentService(segmentRepo)
	segmentHandler := handlers.NewSegmentHandler(segmentService)

	// Repository, Service, and Handler setup for Rollout Strategies
	rolloutRepo := repositories.NewRolloutStrategyRepository(infrastructure.DB)
	rolloutService := services.NewRolloutStrategyService(rolloutRepo)
	rolloutHandler := handlers.NewRolloutStrategyHandler(rolloutService)

	// Router setup
	router := routes.SetupRouter(flagHandler, segmentHandler, rolloutHandler)

	// Start server
	log.Fatal(router.Run(":8080"))
}
