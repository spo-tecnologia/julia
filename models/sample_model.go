package models

import (
	"time"

	"gorm.io/gorm"
)

type SampleModel struct {
	DefaultModel
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name           string         `gorm:"size:255;notNull" validate:"required" json:"name"`
	SampleString   string         `gorm:"size:255;notNull" validate:"required" json:"sample_string"`
	SampleUnique   string         `gorm:"size:255;notNull;unique" validate:"required" json:"sample_unique"`
	SampleDate     time.Time      `gorm:"notNull" validate:"required" json:"sample_date"`
	SampleNullable string         `gorm:"size:255" json:"sample_nullable"`
	SampleDouble   float64        `gorm:"notNull" validate:"required" json:"sample_double"`
	SampleDetail   *SampleDetail  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"sample_detail"`
	SampleDetailID uint           `gorm:"notNull" validate:"required" json:"sample_detail_id"`
	SampleItems    []*SampleItem  `gorm:"foreignKey:SampleModelID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"sample_items"`
	OrderNumber    int            `gorm:"default:0" json:"order_number"`
}
