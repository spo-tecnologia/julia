package factories

import (
	"time"

	"github.com/OdairPianta/julia/config"
	"github.com/OdairPianta/julia/enums"
	"github.com/OdairPianta/julia/models"
	"github.com/OdairPianta/julia/tests/fakers"
)

func CreateUser() (*models.User, error) {
	model := &models.User{
		Name:               fakers.RandomString(),
		Email:              fakers.Email(),
		Password:           fakers.Word(),
		EmailVerifiedAt:    time.Now(),
		RememberToken:      fakers.UUID(),
		FCMToken:           fakers.UUID(),
		Token:              fakers.UUID(),
		ResetPasswordToken: fakers.UUID(),
		Profile:            enums.UserProfileEnumAdministrator,
		Phone:              "+5511988776655",
	}

	err := config.DB.Create(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}
