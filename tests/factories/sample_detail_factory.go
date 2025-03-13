// tests/factories/sample_factory.go

package factories

import (
	"github.com/OdairPianta/julia/config"
	"github.com/OdairPianta/julia/models"
	"github.com/OdairPianta/julia/tests/fakers"
)

func CreateSampleDetail() (*models.SampleDetail, error) {
	model := &models.SampleDetail{
		SampleString: fakers.Word(),
	}
	err := config.DB.Create(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}
