package requests

import (
	"github.com/OdairPianta/julia/enums"
)

type CreateUserInput struct {
	Name          string                `json:"name" binding:"required"`
	Email         string                `json:"email" binding:"required,email,not_exists=users.email"`
	Password      string                `json:"password" binding:"required"`
	RememberToken string                `json:"remember_token"`
	FCMToken      string                `json:"fcm_token"`
	Profile       enums.UserProfileEnum `json:"profile" binding:"required"`
	Phone         string                `json:"phone" binding:"phone_number"`
}
