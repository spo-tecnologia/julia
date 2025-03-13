package models

import (
	"time"
)

type DefaultModel struct {
	ID        uint      `gorm:"primarykey" validate:"required" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
