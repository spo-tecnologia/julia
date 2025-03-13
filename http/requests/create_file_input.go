package requests

type CreateFileInput struct {
	Base64    string `json:"base_64" binding:"required_without=PublicURL"`
	Extension string `json:"extension" binding:"required"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	PublicURL string `json:"public_url" binding:"required_without=Base64"`
}
