package services

import (
	"github.com/besyuzkirk/feature-flag-management/internal/domain/entities"
	"github.com/besyuzkirk/feature-flag-management/internal/domain/repositories"
)

type SegmentService interface {
	CreateSegment(name, description, criteria string) (*entities.Segment, error)
	UpdateSegment(id uint, name, description, criteria string) (*entities.Segment, error)
	DeleteSegment(id uint) error
	GetSegment(id uint) (*entities.Segment, error)
	ListSegments() ([]*entities.Segment, error)
}

type segmentService struct {
	repo repositories.SegmentRepository
}

func NewSegmentService(repo repositories.SegmentRepository) SegmentService {
	return &segmentService{repo}
}

func (s *segmentService) CreateSegment(name, description, criteria string) (*entities.Segment, error) {
	segment := &entities.Segment{
		Name:        name,
		Description: description,
		Criteria:    criteria,
	}

	if err := s.repo.CreateSegment(segment); err != nil {
		return nil, err
	}
	return segment, nil
}

func (s *segmentService) UpdateSegment(id uint, name, description, criteria string) (*entities.Segment, error) {
	segment, err := s.repo.GetSegment(id)
	if err != nil {
		return nil, err
	}

	segment.Name = name
	segment.Description = description
	segment.Criteria = criteria

	if err := s.repo.UpdateSegment(segment); err != nil {
		return nil, err
	}
	return segment, nil
}

func (s *segmentService) DeleteSegment(id uint) error {
	return s.repo.DeleteSegment(id)
}

func (s *segmentService) GetSegment(id uint) (*entities.Segment, error) {
	return s.repo.GetSegment(id)
}

func (s *segmentService) ListSegments() ([]*entities.Segment, error) {
	return s.repo.ListSegments()
}
