package requests

import "time"

type UpdateSampleModelInput struct {
	SampleString   string    `json:"sample_string" binding:"required"`
	SampleUnique   string    `json:"sample_unique" binding:"required"`
	SampleDate     time.Time `json:"sample_date" binding:"required" time_format:"2006-01-02T15:04:05Z07:00"`
	SampleNullable string    `json:"sample_nullable"`
	SampleDouble   float64   `json:"sample_double" binding:"required"`
	SampleBool     bool      `json:"sample_bool" binding:"boolean"`
	SampleDetailID uint      `json:"sample_detail_id" binding:"required,exists=sample_details.id"`
}
