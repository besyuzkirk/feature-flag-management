package entities

import (
	"time"

	"gorm.io/gorm"
)

type RolloutStrategy struct {
	ID            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	FeatureFlagID uint           `gorm:"not null" json:"feature_flag_id"`
	Percentage    int            `gorm:"not null" json:"percentage"` // 0-100 arası bir değer
	Description   string         `json:"description"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (RolloutStrategy) TableName() string {
	return "rollout_strategies"
}
