package services

import (
	"context"

	"github.com/besyuzkirk/feature-flag-management/internal/domain/services"
	pb "github.com/besyuzkirk/feature-flag-management/internal/grpc/protos"
)

type RolloutStrategyServiceServer struct {
	pb.UnimplementedRolloutStrategyServiceServer
	service services.RolloutStrategyService
}

func NewRolloutStrategyServiceServer(service services.RolloutStrategyService) *RolloutStrategyServiceServer {
	return &RolloutStrategyServiceServer{service: service}
}

func (s *RolloutStrategyServiceServer) GetRolloutStrategy(ctx context.Context, req *pb.GetRolloutStrategyRequest) (*pb.RolloutStrategyResponse, error) {
	strategy, err := s.service.GetRolloutStrategy(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.RolloutStrategyResponse{
		RolloutStrategy: &pb.RolloutStrategy{
			Id:            uint32(strategy.ID),
			FeatureFlagId: uint32(strategy.FeatureFlagID),
			Percentage:    int32(strategy.Percentage),
			Description:   strategy.Description,
			CreatedAt:     strategy.CreatedAt.String(),
			UpdatedAt:     strategy.UpdatedAt.String(),
		},
	}, nil
}

func (s *RolloutStrategyServiceServer) ListRolloutStrategiesByFlag(ctx context.Context, req *pb.GetRolloutStrategiesByFlagRequest) (*pb.ListRolloutStrategiesResponse, error) {
	strategies, err := s.service.ListRolloutStrategiesByFlag(uint(req.FeatureFlagId))
	if err != nil {
		return nil, err
	}

	var strategyResponses []*pb.RolloutStrategy
	for _, strategy := range strategies {
		strategyResponses = append(strategyResponses, &pb.RolloutStrategy{
			Id:            uint32(strategy.ID),
			FeatureFlagId: uint32(strategy.FeatureFlagID),
			Percentage:    int32(strategy.Percentage),
			Description:   strategy.Description,
			CreatedAt:     strategy.CreatedAt.String(),
			UpdatedAt:     strategy.UpdatedAt.String(),
		})
	}

	return &pb.ListRolloutStrategiesResponse{RolloutStrategies: strategyResponses}, nil
}

func (s *RolloutStrategyServiceServer) CreateRolloutStrategy(ctx context.Context, req *pb.CreateRolloutStrategyRequest) (*pb.RolloutStrategyResponse, error) {
	strategy, err := s.service.CreateRolloutStrategy(uint(req.FeatureFlagId), int(req.Percentage), req.Description)
	if err != nil {
		return nil, err
	}

	return &pb.RolloutStrategyResponse{
		RolloutStrategy: &pb.RolloutStrategy{
			Id:            uint32(strategy.ID),
			FeatureFlagId: uint32(strategy.FeatureFlagID),
			Percentage:    int32(strategy.Percentage),
			Description:   strategy.Description,
			CreatedAt:     strategy.CreatedAt.String(),
			UpdatedAt:     strategy.UpdatedAt.String(),
		},
	}, nil
}

func (s *RolloutStrategyServiceServer) UpdateRolloutStrategy(ctx context.Context, req *pb.UpdateRolloutStrategyRequest) (*pb.RolloutStrategyResponse, error) {
	strategy, err := s.service.UpdateRolloutStrategy(uint(req.Id), int(req.Percentage), req.Description)
	if err != nil {
		return nil, err
	}

	return &pb.RolloutStrategyResponse{
		RolloutStrategy: &pb.RolloutStrategy{
			Id:            uint32(strategy.ID),
			FeatureFlagId: uint32(strategy.FeatureFlagID),
			Percentage:    int32(strategy.Percentage),
			Description:   strategy.Description,
			CreatedAt:     strategy.CreatedAt.String(),
			UpdatedAt:     strategy.UpdatedAt.String(),
		},
	}, nil
}

func (s *RolloutStrategyServiceServer) DeleteRolloutStrategy(ctx context.Context, req *pb.DeleteRolloutStrategyRequest) (*pb.Empty, error) {
	err := s.service.DeleteRolloutStrategy(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *RolloutStrategyServiceServer) TrackRolloutProgress(ctx context.Context, req *pb.TrackRolloutProgressRequest) (*pb.TrackRolloutProgressResponse, error) {
	percentage, err := s.service.TrackRolloutProgress(uint(req.FeatureFlagId))
	if err != nil {
		return nil, err
	}

	return &pb.TrackRolloutProgressResponse{Percentage: int32(percentage)}, nil
}
