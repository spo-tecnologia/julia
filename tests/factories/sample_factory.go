// tests/factories/sample_factory.go

package factories

import (
	"math/rand/v2"
	"time"

	"github.com/OdairPianta/julia/config"
	"github.com/OdairPianta/julia/models"
	"github.com/OdairPianta/julia/tests/fakers"
)

func CreateSampleModel() (*models.SampleModel, error) {
	sampleDetail, err := CreateSampleDetail()
	if err != nil {
		return nil, err
	}

	model := &models.SampleModel{
		SampleString:   fakers.Word(),
		SampleUnique:   fakers.UUID(),
		SampleDate:     time.Now(),
		SampleNullable: fakers.Word(),
		SampleDouble:   rand.Float64(),
		SampleDetailID: sampleDetail.ID,
	}

	err = config.DB.Create(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}
