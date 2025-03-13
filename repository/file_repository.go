package repository

import (
	"github.com/OdairPianta/julia/config"
	"github.com/OdairPianta/julia/models"
)

func FindFiles() (*[]models.File, error) {
	var files []models.File
	err := config.DB.Find(&files).Error
	if err != nil {
		return nil, err
	}
	return &files, nil
}

func CreateFile(file *models.File) error {
	return config.DB.Create(file).Error
}

func DeleteFile(file *models.File) error {
	return config.DB.Delete(file).Error
}
