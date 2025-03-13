package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OdairPianta/julia/models"
	"github.com/OdairPianta/julia/tests/factories"
	"github.com/OdairPianta/julia/tests/fakers"

	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	setupDatabase()
	router := routesSetup()
	_, token, err := initUser()
	assert.Nil(t, err)

	model, err := factories.CreateUser()
	assert.Nil(t, err)
	model.Email = fakers.Email()

	jsonModel, err := json.Marshal(model)
	assert.Nil(t, err)

	request, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonModel))
	request.Header.Set("Authorization", "Bearer "+token)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	fmt.Println(recorder.Body.String())

	assert.Equal(t, http.StatusOK, recorder.Code, "OK response is expected")

	var result map[string]string
	_ = json.Unmarshal(recorder.Body.Bytes(), &result)
	assert.NotNil(t, result, "response body must be a valid json")

	assert.NotEmpty(t, result["token"], "token must be not empty")
}

func TestUserFind(t *testing.T) {
	setupDatabase()
	router := routesSetup()
	_, token, err := initUser()
	assert.Nil(t, err)

	model, err := factories.CreateUser()
	assert.Nil(t, err)

	request, _ := http.NewRequest("GET", "/api/users/"+fmt.Sprint(model.ID), nil)
	request.Header.Set("Authorization", "Bearer "+token)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	fmt.Println(recorder.Body.String())

	assert.Equal(t, http.StatusOK, recorder.Code, "OK response is expected")

	var resultModel models.User
	_ = json.Unmarshal(recorder.Body.Bytes(), &resultModel)
	assert.NotNil(t, resultModel, "response body must be a valid json")
}

func TestUserUpdatee(t *testing.T) {
	setupDatabase()
	router := routesSetup()
	_, token, err := initUser()
	assert.Nil(t, err)

	model, err := factories.CreateUser()
	assert.Nil(t, err)
	model.Name = fakers.Name()

	jsonModel, err := json.Marshal(model)
	assert.Nil(t, err)

	request, _ := http.NewRequest("PUT", "/api/users/"+fmt.Sprint(model.ID), bytes.NewBuffer(jsonModel))
	request.Header.Set("Authorization", "Bearer "+token)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	fmt.Println(recorder.Body.String())

	assert.Equal(t, http.StatusOK, recorder.Code, "OK response is expected")

	var resultModel models.User
	_ = json.Unmarshal(recorder.Body.Bytes(), &resultModel)
	assert.NotNil(t, resultModel, "response body must be a valid json")
}

func TestUserUpdateFcmToken(t *testing.T) {
	setupDatabase()
	router := routesSetup()
	_, token, err := initUser()
	assert.Nil(t, err)

	model, err := factories.CreateUser()
	assert.Nil(t, err)

	body := []byte(`{"fcm_token": "new_fcm_token"}`)

	request, _ := http.NewRequest("PUT", "/api/users/"+fmt.Sprint(model.ID)+"/update_fcm_token", bytes.NewBuffer(body))
	request.Header.Set("Authorization", "Bearer "+token)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	fmt.Println(recorder.Body.String())

	assert.Equal(t, http.StatusOK, recorder.Code, "OK response is expected")

	var resultModel models.User
	_ = json.Unmarshal(recorder.Body.Bytes(), &resultModel)
	assert.NotNil(t, resultModel, "response body must be a valid json")
	//assert fcm token is equal new_fcm_token
	assert.Equal(t, "new_fcm_token", resultModel.FCMToken, "fcm token must be equal new_fcm_token")
}

func TestUserDelete(t *testing.T) {
	setupDatabase()
	router := routesSetup()
	_, token, err := initUser()
	assert.Nil(t, err)

	model, err := factories.CreateUser()
	assert.Nil(t, err)

	request, _ := http.NewRequest("DELETE", "/api/users/"+fmt.Sprint(model.ID), nil)
	request.Header.Set("Authorization", "Bearer "+token)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	fmt.Println(recorder.Body.String())

	assert.Equal(t, http.StatusOK, recorder.Code, "OK response is expected")

	var resultModel models.User
	_ = json.Unmarshal(recorder.Body.Bytes(), &resultModel)
	assert.NotNil(t, resultModel, "response body must be a valid json")
}

func TestUserSelect(t *testing.T) {
	setupDatabase()
	router := routesSetup()
	_, token, err := initUser()
	assert.Nil(t, err)

	model, err := factories.CreateUser()
	assert.Nil(t, err)

	request, _ := http.NewRequest("GET", fmt.Sprintf("/api/users_select?search=%s", model.Name), nil)
	request.Header.Set("Authorization", "Bearer "+token)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code, "Returned body: "+recorder.Body.String())
	var result []models.ItemSelect
	err = json.Unmarshal(recorder.Body.Bytes(), &result)
	assert.Nil(t, err, "Returned body: "+recorder.Body.String())
	assert.Greater(t, len(result), 0, "Returned body: "+recorder.Body.String())
}
