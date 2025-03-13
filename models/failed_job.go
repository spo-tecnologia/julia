package models

import (
	"time"
)

type FailedJob struct {
	DefaultModel
	JobID     uint      `gorm:"not null" validate:"required" json:"job_id"`
	Exception string    `gorm:"type:text;not null" validate:"required" json:"exception"`
	FailedAt  time.Time `gorm:"not null" validate:"required" json:"failed_at"`
}
