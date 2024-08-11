package main

import (
	"log"

	"github.com/besyuzkirk/feature-flag-management/cmd/server"
	"github.com/besyuzkirk/feature-flag-management/internal/domain/entities"
	"github.com/besyuzkirk/feature-flag-management/internal/infrastructure"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {

	cont, err := infrastructure.NewContainer()
	if err != nil {
		log.Fatalf("Could not initialize container: %v", err)
	}

	// Veritabanı migration işlemi
	infrastructure.Migrate(cont.DB,
		&entities.FeatureFlag{},
		&entities.FeatureFlagHistory{},
		&entities.RolloutStrategy{},
		&entities.Segment{},
	)

	go func() {
		log.Println("Starting gRPC server...")
		server.StartGRPCServer(cont)
		log.Println("gRPC server has started.")
	}()

	server.StartHTTPServer(cont)

}
