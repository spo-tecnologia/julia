package models

import (
	"gorm.io/gorm"
)

type SampleItem struct {
	DefaultModel
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	SampleString  string         `gorm:"size:255;notNull" validate:"required" json:"sample_string"`
	SampleModelID uint           `gorm:"notNull" validate:"required" json:"sample_model_id"`
}
