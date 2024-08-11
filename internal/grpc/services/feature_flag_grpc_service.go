package services

import (
	"context"

	"github.com/besyuzkirk/feature-flag-management/internal/domain/services"
	pb "github.com/besyuzkirk/feature-flag-management/internal/grpc/generated/feature_flag"
)

type FeatureFlagServiceServer struct {
	pb.UnimplementedFeatureFlagServiceServer
	service services.FeatureFlagService
}

func NewFeatureFlagServiceServer(service services.FeatureFlagService) *FeatureFlagServiceServer {
	return &FeatureFlagServiceServer{service: service}
}

func (s *FeatureFlagServiceServer) GetFeatureFlag(ctx context.Context, req *pb.GetFeatureFlagRequest) (*pb.GetFeatureFlagResponse, error) {
	flag, err := s.service.GetFeatureFlag(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.GetFeatureFlagResponse{
		FeatureFlag: &pb.FeatureFlag{
			Id:          uint32(flag.ID),
			Name:        flag.Name,
			Description: flag.Description,
			IsActive:    flag.IsActive,
			CreatedAt:   flag.CreatedAt.String(),
			UpdatedAt:   flag.UpdatedAt.String(),
		},
	}, nil
}

func (s *FeatureFlagServiceServer) ListFeatureFlags(ctx context.Context, req *pb.Empty) (*pb.ListFeatureFlagsResponse, error) {
	flags, err := s.service.ListFeatureFlags()
	if err != nil {
		return nil, err
	}

	var flagResponses []*pb.FeatureFlag
	for _, flag := range flags {
		flagResponses = append(flagResponses, &pb.FeatureFlag{
			Id:          uint32(flag.ID),
			Name:        flag.Name,
			Description: flag.Description,
			IsActive:    flag.IsActive,
			CreatedAt:   flag.CreatedAt.String(),
			UpdatedAt:   flag.UpdatedAt.String(),
		})
	}

	return &pb.ListFeatureFlagsResponse{FeatureFlags: flagResponses}, nil
}

func (s *FeatureFlagServiceServer) CreateFeatureFlag(ctx context.Context, req *pb.CreateFeatureFlagRequest) (*pb.FeatureFlagResponse, error) {
	flag, err := s.service.CreateFeatureFlag(req.Name, req.Description)
	if err != nil {
		return nil, err
	}

	return &pb.FeatureFlagResponse{
		FeatureFlag: &pb.FeatureFlag{
			Id:          uint32(flag.ID),
			Name:        flag.Name,
			Description: flag.Description,
			IsActive:    flag.IsActive,
			CreatedAt:   flag.CreatedAt.String(),
			UpdatedAt:   flag.UpdatedAt.String(),
		},
	}, nil
}

func (s *FeatureFlagServiceServer) UpdateFeatureFlag(ctx context.Context, req *pb.UpdateFeatureFlagRequest) (*pb.FeatureFlagResponse, error) {
	flag, err := s.service.UpdateFeatureFlag(uint(req.Id), req.Name, req.Description, req.IsActive)
	if err != nil {
		return nil, err
	}

	return &pb.FeatureFlagResponse{
		FeatureFlag: &pb.FeatureFlag{
			Id:          uint32(flag.ID),
			Name:        flag.Name,
			Description: flag.Description,
			IsActive:    flag.IsActive,
			CreatedAt:   flag.CreatedAt.String(),
			UpdatedAt:   flag.UpdatedAt.String(),
		},
	}, nil
}

func (s *FeatureFlagServiceServer) DeleteFeatureFlag(ctx context.Context, req *pb.DeleteFeatureFlagRequest) (*pb.Empty, error) {
	err := s.service.DeleteFeatureFlag(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *FeatureFlagServiceServer) RollbackFeatureFlag(ctx context.Context, req *pb.RollbackFeatureFlagRequest) (*pb.FeatureFlagResponse, error) {
	flag, err := s.service.RollbackFeatureFlag(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.FeatureFlagResponse{
		FeatureFlag: &pb.FeatureFlag{
			Id:          uint32(flag.ID),
			Name:        flag.Name,
			Description: flag.Description,
			IsActive:    flag.IsActive,
			CreatedAt:   flag.CreatedAt.String(),
			UpdatedAt:   flag.UpdatedAt.String(),
		},
	}, nil
}
