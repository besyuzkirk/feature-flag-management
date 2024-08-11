package server

import (
	"log"

	"github.com/besyuzkirk/feature-flag-management/internal/handlers"
	"github.com/besyuzkirk/feature-flag-management/internal/infrastructure"
	"github.com/besyuzkirk/feature-flag-management/internal/routes"
)

func StartHTTPServer(cont *infrastructure.Container) {

	flagHandler := handlers.NewFeatureFlagHandler(cont.FeatureFlagService)
	segmentHandler := handlers.NewSegmentHandler(cont.SegmentService)
	rolloutHandler := handlers.NewRolloutStrategyHandler(cont.RolloutStrategyService)

	router := routes.SetupRouter(flagHandler, segmentHandler, rolloutHandler)

	log.Println("Starting HTTP server on port 8080...")
	log.Fatal(router.Run(":8080"))
}
