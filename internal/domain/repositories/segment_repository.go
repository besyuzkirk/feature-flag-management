package repositories

import (
	"errors"

	"github.com/besyuzkirk/feature-flag-management/internal/domain/entities"
	"gorm.io/gorm"
)

type SegmentRepository interface {
	CreateSegment(segment *entities.Segment) error
	UpdateSegment(segment *entities.Segment) error
	DeleteSegment(id uint) error
	GetSegment(id uint) (*entities.Segment, error)
	ListSegments() ([]*entities.Segment, error)
}

type segmentRepository struct {
	db *gorm.DB
}

func NewSegmentRepository(db *gorm.DB) SegmentRepository {
	return &segmentRepository{db}
}

func (r *segmentRepository) CreateSegment(segment *entities.Segment) error {
	return r.db.Create(segment).Error
}

func (r *segmentRepository) UpdateSegment(segment *entities.Segment) error {
	if r.db.Model(segment).Where("id = ?", segment.ID).Updates(segment).RowsAffected == 0 {
		return errors.New("no rows affected or segment not found")
	}
	return nil
}

func (r *segmentRepository) DeleteSegment(id uint) error {
	if r.db.Delete(&entities.Segment{}, id).RowsAffected == 0 {
		return errors.New("no rows affected or segment not found")
	}
	return nil
}

func (r *segmentRepository) GetSegment(id uint) (*entities.Segment, error) {
	var segment entities.Segment
	if err := r.db.First(&segment, id).Error; err != nil {
		return nil, err
	}
	return &segment, nil
}

func (r *segmentRepository) ListSegments() ([]*entities.Segment, error) {
	var segments []*entities.Segment
	if err := r.db.Find(&segments).Error; err != nil {
		return nil, err
	}
	return segments, nil
}
