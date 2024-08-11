package entities

import (
	"time"

	"gorm.io/gorm"
)

type FeatureFlagHistory struct {
	ID            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	FeatureFlagID uint           `gorm:"not null" json:"feature_flag_id"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	IsActive      bool           `json:"is_active"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (FeatureFlagHistory) TableName() string {
	return "feature_flag_histories"
}
