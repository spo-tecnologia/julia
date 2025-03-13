package requests

import (
	"github.com/OdairPianta/julia/enums"
)

type UpdateUserInput struct {
	Name               string                `json:"name" binding:"required"`
	Email              string                `json:"email" binding:"required,email"`
	Password           string                `json:"password"`
	RememberToken      string                `json:"remember_token"`
	FCMToken           string                `json:"fcm_token"`
	ResetPasswordToken string                `json:"reset_password_token"`
	Profile            enums.UserProfileEnum `json:"profile" binding:"required"`
	Phone              string                `json:"phone" binding:"phone_number"`
}
