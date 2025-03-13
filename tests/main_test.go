package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/OdairPianta/julia/config"
	"github.com/OdairPianta/julia/enums"
	"github.com/OdairPianta/julia/models"
	"github.com/OdairPianta/julia/routes"
	"github.com/OdairPianta/julia/services/token"
	"github.com/OdairPianta/julia/tests/fakers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
)

func setupDatabase() (*gorm.DB, func()) {
	config.InitDatabase()
	tx := config.DB.Begin()

	cleanup := func() {
		tx.Rollback()
	}

	return tx, cleanup
}

func routesSetup() *gin.Engine {
	r := gin.Default()
	routes.InitRoutes(r)
	return r
}

func initUser() (user models.User, stringToken string, err error) {
	user = models.User{
		Name:               fakers.Name(),
		Email:              fakers.Email(),
		Password:           fakers.Word(),
		EmailVerifiedAt:    time.Now(),
		RememberToken:      fakers.UUID(),
		FCMToken:           fakers.UUID(),
		Token:              fakers.UUID(),
		ResetPasswordToken: fakers.UUID(),
		Profile:            enums.UserProfileEnumAdministrator,
	}
	err = config.DB.Create(&user).Error
	if err != nil {
		fmt.Println(err)
		return user, "", err
	}

	token, errToken := token.GenerateToken(user.ID)
	if errToken != nil {
		fmt.Println(err)
		return user, "", errToken
	}
	return user, token, nil
}

func TestFindGitHeadFilesReturnAccessDeniedResponse(t *testing.T) {
	setupDatabase()
	router := routesSetup()

	request, _ := http.NewRequest("GET", "/.git/HEAD", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	fmt.Println(recorder.Body.String())

	assert.Equal(t, http.StatusUnauthorized, recorder.Code, "Unauthorized response is expected")

	var resultModel map[string]string
	_ = json.Unmarshal(recorder.Body.Bytes(), &resultModel)
	assert.NotNil(t, resultModel, "response body must be a valid json")
}

func TestFindGitConfigFilesReturnAccessDeniedResponse(t *testing.T) {
	setupDatabase()
	router := routesSetup()

	request, _ := http.NewRequest("GET", "/.git/config", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	fmt.Println(recorder.Body.String())

	assert.Equal(t, http.StatusUnauthorized, recorder.Code, "Unauthorized response is expected")

	var resultModel map[string]string
	_ = json.Unmarshal(recorder.Body.Bytes(), &resultModel)
	assert.NotNil(t, resultModel, "response body must be a valid json")
}

func TestFindEnvFilesReturnAccessDeniedResponse(t *testing.T) {
	setupDatabase()
	router := routesSetup()

	request, _ := http.NewRequest("GET", "/.env", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	fmt.Println(recorder.Body.String())

	assert.Equal(t, http.StatusUnauthorized, recorder.Code, "Unauthorized response is expected")

	var resultModel map[string]string
	_ = json.Unmarshal(recorder.Body.Bytes(), &resultModel)
	assert.NotNil(t, resultModel, "response body must be a valid json")
}
