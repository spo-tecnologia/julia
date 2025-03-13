package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type File struct {
	DefaultModel
	Extension string `gorm:"not null" validate:"required" json:"extension"`
	Path      string `gorm:"size:255;null" json:"path"`
	PublicURL string `gorm:"size:255;null" json:"public_url"`
	Name      string `gorm:"size:255;null" json:"name"`
	PostID    *uint  `gorm:"null" json:"post_id"`
}

func SaveGryphonAPI(base64, extension, filePath string) (string, error) {
	gryphonToken := os.Getenv("GRYPHON_API_TOKEN")
	gryphonURL := os.Getenv("GRYPHON_API_BASE_URL") + "/files/base64/create"

	payload := map[string]interface{}{
		"base64":    base64,
		"extension": extension,
		"file_path": filePath,
	}
	jsonData, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", gryphonURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+gryphonToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error: %v", response["message"])
	}

	publicURL, ok := response["public_url"].(string)
	if !ok {
		return "", errors.New("invalid response from API")
	}

	return publicURL, nil
}

func SaveImageGryphonAPI(base64, extension, filePath string, width, height int) (string, error) {
	gryphonToken := os.Getenv("GRYPHON_API_TOKEN")
	gryphonURL := os.Getenv("GRYPHON_API_BASE_URL") + "/images/base64/create"

	payload := map[string]interface{}{
		"base64":    base64,
		"extension": extension,
		"file_path": filePath,
		"width":     width,
		"height":    height,
	}
	jsonData, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", gryphonURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+gryphonToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error: %v", response["message"])
	}

	publicURL, ok := response["public_url"].(string)
	if !ok {
		return "", errors.New("invalid response from API")
	}

	return publicURL, nil
}
