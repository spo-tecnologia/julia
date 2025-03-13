package requests

type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required,email"`
}
