package main

import (
	"log"

	"github.com/besyuzkirk/feature-flag-management/cmd/server"
	"github.com/besyuzkirk/feature-flag-management/config"
	"github.com/besyuzkirk/feature-flag-management/internal/domain/entities"
	"github.com/besyuzkirk/feature-flag-management/internal/infrastructure"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
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

	infrastructure.Migrate(
		&entities.FeatureFlag{},
		&entities.FeatureFlagHistory{},
		&entities.RolloutStrategy{},
		&entities.Segment{},
	)

	go func() {
		log.Println("Starting gRPC server...")
		server.StartGRPCServer()
		log.Println("gRPC server has started.")
	}()

	server.StartHTTPServer()

}
