package models

type ItemSelect struct {
	ID   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
