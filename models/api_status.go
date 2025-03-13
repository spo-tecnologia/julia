package models

type APIStatus struct {
	Status string `json:"status" validate:"required"`
}
