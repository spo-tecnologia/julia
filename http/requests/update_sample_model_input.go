package requests

import "time"

type UpdateSampleModelInput struct {
	Name           string    `json:"name" binding:"required"`
	SampleString   string    `json:"sample_string" binding:"required"`
	SampleUnique   string    `json:"sample_unique" binding:"required"`
	SampleDate     time.Time `json:"sample_date" binding:"required" time_format:"2006-01-02T15:04:05Z07:00"`
	SampleNullable string    `json:"sample_nullable"`
	SampleDouble   float64   `json:"sample_double" binding:"required"`
	SampleDetailID uint      `json:"sample_detail_id" binding:"required,exists=sample_details.id"`
	OrderNumber    int       `json:"order_number" binding:"required"`
}
