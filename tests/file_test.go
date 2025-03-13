package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/OdairPianta/julia/config"
	"github.com/OdairPianta/julia/models"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestFileFindAll(t *testing.T) {
	setupDatabase()
	router := routesSetup()
	_, token, err := initUser()
	assert.Nil(t, err)

	fmt.Println(`>> GET: /api/files`)

	request, _ := http.NewRequest("GET", "/api/files", nil)
	request.Header.Set("Authorization", "Bearer "+token)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code, "OK response is expected")

	var files []models.File
	_ = json.Unmarshal(recorder.Body.Bytes(), &files)
	assert.NotNil(t, files, "response body must be a valid json")
}

func TestFileCreate(t *testing.T) {
	setupDatabase()
	router := routesSetup()
	_, token, err := initUser()
	assert.Nil(t, err)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	gryphonURL := os.Getenv("GRYPHON_API_BASE_URL") + "/files/base64/create"
	httpmock.RegisterResponder("POST", gryphonURL,
		httpmock.NewStringResponder(200, `{"public_url": "https://example.com/image.jpg"}`))

	input := map[string]interface{}{
		"base_64":   "base64string",
		"extension": "jpg",
		"path":      "uploads/users/1/",
		"name":      "profile_picture",
	}
	jsonInput, _ := json.Marshal(input)

	request, _ := http.NewRequest("POST", "/api/files", bytes.NewBuffer(jsonInput))
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	fmt.Println(recorder.Body.String())

	assert.Equal(t, http.StatusOK, recorder.Code, "OK response is expected")

	var file models.File
	_ = json.Unmarshal(recorder.Body.Bytes(), &file)
	assert.NotNil(t, file, "response body must be a valid json")
	assert.Equal(t, "https://example.com/image.jpg", file.PublicURL, "Public URL should match the mocked response")
}

func TestFileCreateWithUrl(t *testing.T) {
	setupDatabase()
	router := routesSetup()
	_, token, err := initUser()
	assert.Nil(t, err)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	input := map[string]interface{}{
		"base_64":    nil,
		"extension":  "png",
		"path":       "uploads/users/1/",
		"name":       "profile_picture",
		"public_url": "https://example.com/image.jpg",
	}
	jsonInput, _ := json.Marshal(input)

	request, _ := http.NewRequest("POST", "/api/files", bytes.NewBuffer(jsonInput))
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	fmt.Println(recorder.Body.String())

	assert.Equal(t, http.StatusOK, recorder.Code, "Body: "+recorder.Body.String())

	var file models.File
	_ = json.Unmarshal(recorder.Body.Bytes(), &file)
	assert.NotNil(t, file, "Body: "+recorder.Body.String())
	assert.Equal(t, "https://example.com/image.jpg", file.PublicURL, "Body: "+recorder.Body.String())
}

func TestFileDelete(t *testing.T) {
	setupDatabase()
	router := routesSetup()
	_, token, err := initUser()
	assert.Nil(t, err)

	file := models.File{
		Extension: "jpg",
		Path:      "uploads/users/1/",
		PublicURL: "https://example.com/image.jpg",
		Name:      "profile_picture",
	}
	config.DB.Create(&file)

	request, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/files/%d", file.ID), nil)
	request.Header.Set("Authorization", "Bearer "+token)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code, "OK response is expected")
}
