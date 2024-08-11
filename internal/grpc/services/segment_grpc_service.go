package services

import (
	"context"

	"github.com/besyuzkirk/feature-flag-management/internal/domain/services"
	pb "github.com/besyuzkirk/feature-flag-management/internal/grpc/generated/segment"
)

type SegmentServiceServer struct {
	pb.UnimplementedSegmentServiceServer
	service services.SegmentService
}

func NewSegmentServiceServer(service services.SegmentService) *SegmentServiceServer {
	return &SegmentServiceServer{service: service}
}

func (s *SegmentServiceServer) GetSegment(ctx context.Context, req *pb.GetSegmentRequest) (*pb.SegmentResponse, error) {
	segment, err := s.service.GetSegment(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.SegmentResponse{
		Segment: &pb.Segment{
			Id:          uint32(segment.ID),
			Name:        segment.Name,
			Description: segment.Description,
			Criteria:    segment.Criteria,
			CreatedAt:   segment.CreatedAt.String(),
			UpdatedAt:   segment.UpdatedAt.String(),
		},
	}, nil
}

func (s *SegmentServiceServer) ListSegments(ctx context.Context, req *pb.Empty) (*pb.ListSegmentsResponse, error) {
	segments, err := s.service.ListSegments()
	if err != nil {
		return nil, err
	}

	var segmentResponses []*pb.Segment
	for _, segment := range segments {
		segmentResponses = append(segmentResponses, &pb.Segment{
			Id:          uint32(segment.ID),
			Name:        segment.Name,
			Description: segment.Description,
			Criteria:    segment.Criteria,
			CreatedAt:   segment.CreatedAt.String(),
			UpdatedAt:   segment.UpdatedAt.String(),
		})
	}

	return &pb.ListSegmentsResponse{Segments: segmentResponses}, nil
}

func (s *SegmentServiceServer) CreateSegment(ctx context.Context, req *pb.CreateSegmentRequest) (*pb.SegmentResponse, error) {
	segment, err := s.service.CreateSegment(req.Name, req.Description, req.Criteria)
	if err != nil {
		return nil, err
	}

	return &pb.SegmentResponse{
		Segment: &pb.Segment{
			Id:          uint32(segment.ID),
			Name:        segment.Name,
			Description: segment.Description,
			Criteria:    segment.Criteria,
			CreatedAt:   segment.CreatedAt.String(),
			UpdatedAt:   segment.UpdatedAt.String(),
		},
	}, nil
}

func (s *SegmentServiceServer) UpdateSegment(ctx context.Context, req *pb.UpdateSegmentRequest) (*pb.SegmentResponse, error) {
	segment, err := s.service.UpdateSegment(uint(req.Id), req.Name, req.Description, req.Criteria)
	if err != nil {
		return nil, err
	}

	return &pb.SegmentResponse{
		Segment: &pb.Segment{
			Id:          uint32(segment.ID),
			Name:        segment.Name,
			Description: segment.Description,
			Criteria:    segment.Criteria,
			CreatedAt:   segment.CreatedAt.String(),
			UpdatedAt:   segment.UpdatedAt.String(),
		},
	}, nil
}

func (s *SegmentServiceServer) DeleteSegment(ctx context.Context, req *pb.DeleteSegmentRequest) (*pb.Empty, error) {
	err := s.service.DeleteSegment(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
