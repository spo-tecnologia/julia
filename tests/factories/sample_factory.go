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
		Name:           fakers.Word(),
		SampleString:   fakers.Word(),
		SampleUnique:   fakers.UUID(),
		SampleDate:     time.Now(),
		SampleNullable: fakers.Word(),
		SampleDouble:   rand.Float64(),
		SampleDetailID: sampleDetail.ID,
		SampleDetail:   sampleDetail,
		OrderNumber:    1,
	}

	err = config.DB.Create(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}
