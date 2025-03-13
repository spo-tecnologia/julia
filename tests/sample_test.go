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

func TestSampleModelFindAll(t *testing.T) {
	setupDatabase()
	router := routesSetup()
	_, token, err := initUser()
	assert.Nil(t, err)

	model, err := factories.CreateSampleModel()
	assert.Nil(t, err)

	request, _ := http.NewRequest("GET", "/api/sample_models?&search="+fmt.Sprintf("%d", model.ID), nil)
	request.Header.Set("Authorization", "Bearer "+token)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code, "Returned body: "+recorder.Body.String())
	var resultModels []models.SampleModel
	_ = json.Unmarshal(recorder.Body.Bytes(), &resultModels)
	assert.NotNil(t, resultModels, "Returned body: "+recorder.Body.String())
	assert.Greater(t, len(resultModels), 0, "Returned body: "+recorder.Body.String())
}

func TestSampleModelCreate(t *testing.T) {
	setupDatabase()
	router := routesSetup()
	_, token, err := initUser()
	assert.Nil(t, err)

	model, err := factories.CreateSampleModel()
	assert.Nil(t, err)
	model.SampleUnique = fakers.Word()

	jsonModel, err := json.Marshal(model)
	assert.Nil(t, err)

	request, _ := http.NewRequest("POST", "/api/sample_models", bytes.NewBuffer(jsonModel))
	request.Header.Set("Authorization", "Bearer "+token)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code, "Returned body: "+recorder.Body.String())

	var result models.SampleModel
	_ = json.Unmarshal(recorder.Body.Bytes(), &result)
	assert.NotNil(t, result, "Returned body: "+recorder.Body.String())
}

func TestSampleModelFind(t *testing.T) {
	setupDatabase()
	router := routesSetup()
	_, token, err := initUser()
	assert.Nil(t, err)

	model, err := factories.CreateSampleModel()
	assert.Nil(t, err)

	request, _ := http.NewRequest("GET", "/api/sample_models/"+fmt.Sprint(model.ID), nil)
	request.Header.Set("Authorization", "Bearer "+token)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	fmt.Println(recorder.Body.String())

	assert.Equal(t, http.StatusOK, recorder.Code, "Returned body: "+recorder.Body.String())

	var resultModel models.SampleModel
	_ = json.Unmarshal(recorder.Body.Bytes(), &resultModel)
	assert.NotNil(t, resultModel, "Returned body: "+recorder.Body.String())
}

func TestSampleModelUpdate(t *testing.T) {
	setupDatabase()
	router := routesSetup()
	_, token, err := initUser()
	assert.Nil(t, err)

	model, err := factories.CreateSampleModel()
	assert.Nil(t, err)

	model.SampleString = fakers.Name()
	jsonModel, err := json.Marshal(model)
	if err != nil {
		fmt.Println(err)
		return
	}

	request, _ := http.NewRequest("PUT", "/api/sample_models/"+fmt.Sprint(model.ID), bytes.NewBuffer(jsonModel))
	request.Header.Set("Authorization", "Bearer "+token)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code, "Returned body: "+recorder.Body.String())

	var resultModel models.SampleModel
	_ = json.Unmarshal(recorder.Body.Bytes(), &resultModel)
	assert.NotNil(t, resultModel, "Returned body: "+recorder.Body.String())
}

func TestSampleModelDelete(t *testing.T) {
	setupDatabase()
	router := routesSetup()
	_, token, err := initUser()
	assert.Nil(t, err)

	model, err := factories.CreateSampleModel()
	assert.Nil(t, err)

	request, _ := http.NewRequest("DELETE", "/api/sample_models/"+fmt.Sprint(model.ID), nil)
	request.Header.Set("Authorization", "Bearer "+token)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	fmt.Println(recorder.Body.String())

	assert.Equal(t, http.StatusOK, recorder.Code, "Returned body: "+recorder.Body.String())

	var resultModel models.SampleModel
	_ = json.Unmarshal(recorder.Body.Bytes(), &resultModel)
	assert.NotNil(t, resultModel, "Returned body: "+recorder.Body.String())
}

func TestSampleModelSelect(t *testing.T) {
	setupDatabase()
	router := routesSetup()
	_, token, err := initUser()
	assert.Nil(t, err)

	model, err := factories.CreateSampleModel()
	assert.Nil(t, err)

	request, _ := http.NewRequest("GET", "/api/sample_models_select?&limit=10&search="+fmt.Sprintf("%d", model.ID), nil)
	request.Header.Set("Authorization", "Bearer "+token)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code, "Returned body: "+recorder.Body.String())
	var resultModels []models.SampleModel
	_ = json.Unmarshal(recorder.Body.Bytes(), &resultModels)
	assert.NotNil(t, resultModels, "Returned body: "+recorder.Body.String())
	assert.Greater(t, len(resultModels), 0, "Returned body: "+recorder.Body.String())
}
