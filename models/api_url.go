package models

type APIUrl struct {
	URL string `json:"url" validate:"required"`
}
