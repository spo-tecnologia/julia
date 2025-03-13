package models

import (
	"gorm.io/gorm"
)

type SampleDetail struct {
	DefaultModel
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	SampleString string         `gorm:"size:255;notNull" validate:"required" json:"sample_string"`
}
