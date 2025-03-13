package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type SampleModel struct {
	DefaultModel
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	SampleString   string         `gorm:"size:255;notNull" validate:"required" json:"sample_string"`
	SampleUnique   string         `gorm:"size:255;notNull;unique" validate:"required" json:"sample_unique"`
	SampleDate     time.Time      `gorm:"notNull" validate:"required" json:"sample_date"`
	SampleNullable string         `gorm:"size:255" json:"sample_nullable"`
	SampleDouble   float64        `gorm:"notNull" validate:"required" json:"sample_double"`
	SampleBool     bool           `gorm:"notNull" validate:"required" json:"sample_bool"`
	SampleDetail   SampleDetail   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"sample_detail"`
	SampleDetailID uint           `gorm:"notNull" validate:"required" json:"sample_detail_id"`
	SampleItems    []SampleItem   `gorm:"foreignKey:SampleModelID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"sample_items"`
}

func (s SampleModel) MarshalJSON() ([]byte, error) {
	type Alias SampleModel
	return json.Marshal(&struct {
		Alias
		SampleDate string `json:"sample_date"`
	}{
		Alias:      (Alias)(s),
		SampleDate: s.SampleDate.Format("2006-01-02T15:04:05Z07:00"),
	})
}
