package models

import "time"

type ActivityType int

const (
	Read   = 1
	Create = 2
	Update = 3
	Delete = 4
)

type ActivityLog struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	Activity  ActivityType
	Entity    string
	EntityID  uint
	Changes   string
	UserID    uint
}

type LoggableModel interface {
	ShouldLog() bool
}
