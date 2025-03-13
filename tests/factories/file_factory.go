package factories

import (
	"github.com/OdairPianta/julia/config"
	"github.com/OdairPianta/julia/models"
	"github.com/OdairPianta/julia/tests/fakers"
)

func CreateFile() (*models.File, error) {
	model := &models.File{
		Extension: "jpg",
		Path:      "uploads/users/1/",
		PublicURL: "https://example.com/image.jpg",
		Name:      fakers.RandomString(),
	}
	err := config.DB.Create(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}
