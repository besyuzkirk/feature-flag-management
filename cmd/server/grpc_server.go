package server

import (
	"log"
	"net"

	"github.com/besyuzkirk/feature-flag-management/internal/domain/repositories"
	"github.com/besyuzkirk/feature-flag-management/internal/domain/services"
	pb "github.com/besyuzkirk/feature-flag-management/internal/grpc/generated/feature_flag"
	grpcService "github.com/besyuzkirk/feature-flag-management/internal/grpc/services"
	"github.com/besyuzkirk/feature-flag-management/internal/infrastructure"
	"google.golang.org/grpc"
)

func StartGRPCServer() {
	flagRepo := repositories.NewFeatureFlagRepository(infrastructure.DB)
	flagService := services.NewFeatureFlagService(flagRepo)
	grpcServer := grpcService.NewFeatureFlagServiceServer(flagService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterFeatureFlagServiceServer(server, grpcServer)

	log.Println("Starting gRPC server on port 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
