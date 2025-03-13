package requests

type UpdateUserFcmTokenInput struct {
	FCMToken string `json:"fcm_token" binding:"required"`
}
