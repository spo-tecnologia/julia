package repository

import (
	"github.com/OdairPianta/julia/config"
	"github.com/OdairPianta/julia/http/requests"
	"github.com/OdairPianta/julia/models"
)

func FindUsers(search string, limit *int, offset *int) (*[]models.User, error) {
	var users []models.User
	query := config.DB
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}
	if limit != nil && *limit > 0 {
		query = query.Limit(*limit)
	}
	if offset != nil && *offset > 0 {
		query = query.Offset(*offset)
	}
	err := query.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func FindUserByID(ID string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("id = ?", ID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByCpf(cpf string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("cpf = ?", cpf).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *models.User) error {
	return config.DB.Create(&user).Error
}

func UpdateUser(user *models.User, input *requests.UpdateUserInput) error {
	return config.DB.Model(&user).Updates(input).Error
}

func UpdateFcmToken(user *models.User, fcmToken string) error {
	user.FCMToken = fcmToken
	return config.DB.Save(&user).Error
}

func DeleteUser(user *models.User) error {
	return config.DB.Delete(&user).Error
}

func FindUserSelects(search string, limit *int, offset *int) ([]models.ItemSelect, error) {
	users, err := FindUsers(search, limit, offset)
	if err != nil {
		return nil, err
	}
	var itemSelects []models.ItemSelect
	for _, user := range *users {
		itemSelects = append(itemSelects, models.ItemSelect{
			ID:   user.ID,
			Name: user.Name,
		})
	}
	return itemSelects, nil
}
