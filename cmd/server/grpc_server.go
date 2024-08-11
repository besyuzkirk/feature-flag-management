package server

import (
	"log"
	"net"

	"github.com/besyuzkirk/feature-flag-management/internal/domain/repositories"
	"github.com/besyuzkirk/feature-flag-management/internal/domain/services"
	featureFlagPb "github.com/besyuzkirk/feature-flag-management/internal/grpc/generated/feature_flag"
	rolloutStrategyPb "github.com/besyuzkirk/feature-flag-management/internal/grpc/generated/rollout_strategy"
	segmentPb "github.com/besyuzkirk/feature-flag-management/internal/grpc/generated/segment"
	grpcService "github.com/besyuzkirk/feature-flag-management/internal/grpc/services"
	"github.com/besyuzkirk/feature-flag-management/internal/infrastructure"
	"google.golang.org/grpc"
)

func StartGRPCServer() {
	flagRepo := repositories.NewFeatureFlagRepository(infrastructure.DB)
	flagService := services.NewFeatureFlagService(flagRepo)
	flagGrpcServer := grpcService.NewFeatureFlagServiceServer(flagService)

	rolloutRepo := repositories.NewRolloutStrategyRepository(infrastructure.DB)
	rolloutService := services.NewRolloutStrategyService(rolloutRepo)
	rolloutGrpcServer := grpcService.NewRolloutStrategyServiceServer(rolloutService)

	segmentRepo := repositories.NewSegmentRepository(infrastructure.DB)
	segmentService := services.NewSegmentService(segmentRepo)
	segmentGrpcServer := grpcService.NewSegmentServiceServer(segmentService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	server := grpc.NewServer()

	featureFlagPb.RegisterFeatureFlagServiceServer(server, flagGrpcServer)
	rolloutStrategyPb.RegisterRolloutStrategyServiceServer(server, rolloutGrpcServer)
	segmentPb.RegisterSegmentServiceServer(server, segmentGrpcServer)

	log.Println("Starting gRPC server on port 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
