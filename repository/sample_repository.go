package repository

import (
	"github.com/OdairPianta/julia/config"
	"github.com/OdairPianta/julia/models"
)

func FindSamples(search string, limit *int, offset *int) ([]models.SampleModel, error) {
	var sampleModels []models.SampleModel
	query := config.DB.Preload("SampleDetail").Preload("SampleItems")
	if search != "" {
		query = query.Where("id LIKE ? OR sample_string LIKE ? OR sample_unique LIKE ? OR sample_nullable LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}
	if limit != nil && *limit > 0 {
		query = query.Limit(*limit)
	}
	if offset != nil && *offset > 0 {
		query = query.Offset(*offset)
	}
	err := query.Find(&sampleModels).Error
	if err != nil {
		return nil, err
	}
	return sampleModels, nil
}

func FindSampleByID(ID string) (*models.SampleModel, error) {
	var model models.SampleModel
	err := config.DB.Preload("SampleDetail").Preload("SampleItems").Where("id = ?", ID).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func CreateSampleModel(sampleModel *models.SampleModel) error {
	return config.DB.Create(&sampleModel).Error
}

func UpdateSampleModel(model *models.SampleModel, updatedSampleModel *models.SampleModel) error {
	return config.DB.Model(&model).Updates(updatedSampleModel).Error
}

func DeleteSampleModel(model *models.SampleModel) error {
	return config.DB.Delete(&model).Error
}

func FindSampleItemSelects(search string, limit *int, offset *int) ([]models.ItemSelect, error) {
	sampleModels, err := FindSamples(search, limit, offset)
	if err != nil {
		return nil, err
	}
	var itemSelects []models.ItemSelect
	for _, sampleModel := range sampleModels {
		itemSelects = append(itemSelects, models.ItemSelect{
			ID:   sampleModel.ID,
			Name: sampleModel.SampleString,
		})
	}
	return itemSelects, nil
}
