package models

type APIMessage struct {
	Error string `json:"message" validate:"required"`
}
