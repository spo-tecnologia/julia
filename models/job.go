package models

type JobStatus string

const (
	JobStatusPending   JobStatus = "pending"
	JobStatusCompleted JobStatus = "completed"
)

type Job struct {
	DefaultModel
	Type    string    `gorm:"size:255;not null" validate:"required" json:"type"`
	Payload string    `gorm:"type:text;not null" validate:"required" json:"payload"`
	Status  JobStatus `gorm:"default:'pending';not null" validate:"required" json:"status"`
}
