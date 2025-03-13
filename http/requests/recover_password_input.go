package requests

type RecoverPasswordInput struct {
	ResetPasswordToken string `json:"reset_password_token" binding:"required"`
	Password           string `json:"password" binding:"required"`
}
