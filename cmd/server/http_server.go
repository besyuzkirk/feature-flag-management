package server

import (
	"log"

	"github.com/besyuzkirk/feature-flag-management/internal/domain/repositories"
	"github.com/besyuzkirk/feature-flag-management/internal/domain/services"
	"github.com/besyuzkirk/feature-flag-management/internal/handlers"
	"github.com/besyuzkirk/feature-flag-management/internal/infrastructure"
	"github.com/besyuzkirk/feature-flag-management/internal/routes"
)

func StartHTTPServer() {

	flagRepo := repositories.NewFeatureFlagRepository(infrastructure.DB)
	flagService := services.NewFeatureFlagService(flagRepo)
	flagHandler := handlers.NewFeatureFlagHandler(flagService)

	segmentRepo := repositories.NewSegmentRepository(infrastructure.DB)
	segmentService := services.NewSegmentService(segmentRepo)
	segmentHandler := handlers.NewSegmentHandler(segmentService)

	rolloutRepo := repositories.NewRolloutStrategyRepository(infrastructure.DB)
	rolloutService := services.NewRolloutStrategyService(rolloutRepo)
	rolloutHandler := handlers.NewRolloutStrategyHandler(rolloutService)

	router := routes.SetupRouter(flagHandler, segmentHandler, rolloutHandler)

	log.Println("Starting HTTP server on port 8080...")
	log.Fatal(router.Run(":8080"))
}
