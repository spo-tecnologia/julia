package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/OdairPianta/julia/http/requests"
	"github.com/OdairPianta/julia/models"
	"github.com/OdairPianta/julia/policies"
	"github.com/OdairPianta/julia/repository"
	"github.com/OdairPianta/julia/services"
	"github.com/OdairPianta/julia/services/token"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Get all users
// @Schemes
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Param search query string false "Search by name"
// @Success 200 {array} models.User "ok"
// @Failure 400 {object} models.APIMessage "Bad request"
// @Failure 404 {object} models.APIMessage "Not found"
// @Router /users [get]
// @Security Bearer
func FindUsers(context *gin.Context) {
	user := context.MustGet("user").(models.User)
	if !policies.NewUserPolicy(&user).ViewAny() {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Você não tem permissão!"})
		return
	}

	search := context.Query("search")
	limit, _ := strconv.Atoi(context.Query("limit"))
	offset, _ := strconv.Atoi(context.Query("offset"))

	users, err := repository.FindUsers(search, &limit, &offset)

	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Erro: %s", err.Error())})
		return
	}

	context.JSON(http.StatusOK, users)
}

// @Summary Get user
// @Schemes
// @Description Get user
// @Tags users
// @Accept json
// @Produce json
// @Param   id     path    int     true        "User ID"
// @Success 200 {object} models.User "ok"
// @Failure 400 {object} models.APIMessage "Bad request"
// @Failure 404 {object} models.APIMessage "Not found"
// @Router /users/{id} [get]
// @Security Bearer
func FindUser(context *gin.Context) {
	currentUser := context.MustGet("user").(models.User)
	if !policies.NewUserPolicy(&currentUser).ViewAny() {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Você não tem permissão!"})
		return
	}

	user, err := repository.FindUserByID(context.Param("id"))

	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusNotFound, gin.H{"message": "Usuário não encontrado!"})
		return
	}

	context.JSON(http.StatusOK, user)
}

// @Summary		Add an user
// @Description	add by json user
// @Tags			users
// @Accept		json
// @Produce		json
// @Param			user	body		requests.CreateUserInput	true	"Add user"
// @Success		200		{object}	models.User
// @Failure 400 {object} models.APIMessage "Bad request"
// @Failure 404 {object} models.APIMessage "Not found"
// @Router			/users [post]
// @Security Bearer
func CreateUser(context *gin.Context) {
	currentUser := context.MustGet("user").(models.User)
	if !policies.NewUserPolicy(&currentUser).Create() {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Você não tem permissão!"})
		return
	}
	// Validate input
	var input requests.CreateUserInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": services.GetErrosMessage(err)})
		return
	}

	//check email exists
	_, err := repository.FindUserByEmail(input.Email)
	if err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Email já cadastrado!"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao criptografar a senha!"})
		return
	}

	// Create user
	user := models.User{
		Name:            input.Name,
		Email:           input.Email,
		EmailVerifiedAt: time.Now(),
		Password:        string(hashedPassword),
		FCMToken:        input.FCMToken,
		Profile:         input.Profile,
		Phone:           input.Phone,
	}
	err = repository.CreateUser(&user)

	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Erro: %s", err.Error())})
		return
	}

	token, err := token.GenerateToken(user.ID)
	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao gerar o token do usuário"})
		return
	}
	user.Token = token

	context.JSON(http.StatusOK, user)
}

// @Summary		Update an user
// @Description	Update by json user
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"ID"
// @Param			user	body		requests.UpdateUserInput	true	"Update user"
// @Success		200		{object}	models.User
// @Failure 400 {object} models.APIMessage "Bad request"
// @Failure 404 {object} models.APIMessage "Not found"
// @Router			/users/{id} [put]
// @Security Bearer
func UpdateUser(context *gin.Context) {
	// Get model if exist
	user, err := repository.FindUserByID(context.Param("id"))
	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusNotFound, gin.H{"message": "Usuário não encontrado!"})
		return
	}

	currentUser := context.MustGet("user").(models.User)
	if !policies.NewUserPolicy(&currentUser).Update(user) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Você não tem permissão!"})
		return
	}

	// Validate input
	var input requests.UpdateUserInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": services.GetErrosMessage(err)})
		return
	}

	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			sentry.CaptureException(err)
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao criptografar a senha!"})
			return
		}
		input.Password = string(hashedPassword)
	}

	err = repository.UpdateUser(user, &input)

	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Erro: %s", err.Error())})
		return
	}

	context.JSON(http.StatusOK, user)
}

// @Summary		Update an user fcm token
// @Description	Update by json user
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"ID"
// @Param			user	body		requests.UpdateUserFcmTokenInput	true	"Update user fcm token"
// @Success		200		{object}	models.User
// @Failure 400 {object} models.APIMessage "Bad request"
// @Failure 404 {object} models.APIMessage "Not found"
// @Router			/users/{id}/update_fcm_token [put]
// @Security Bearer
func UpdateFcmToken(context *gin.Context) {
	// Get model if exist
	user, err := repository.FindUserByID(context.Param("id"))

	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusNotFound, gin.H{"message": "Usuário não encontrado!"})
		return
	}

	currentUser := context.MustGet("user").(models.User)
	if !policies.NewUserPolicy(&currentUser).Update(user) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Você não tem permissão!"})
		return
	}
	// Validate input
	var input requests.UpdateUserFcmTokenInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": services.GetErrosMessage(err)})
		return
	}

	err = repository.UpdateFcmToken(user, input.FCMToken)

	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Erro: %s", err.Error())})
		return
	}

	context.JSON(http.StatusOK, user)
}

// @Summary		Delete an user
// @Description	Delete by user ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"ID"	Format(int64)
// @Success		200	{object}	models.User
// @Failure 400 {object} models.APIMessage "Bad request"
// @Failure 404 {object} models.APIMessage "Not found"
// @Router			/users/{id} [delete]
// @Security Bearer
func DeleteUser(context *gin.Context) {
	// Get model if exist
	user, err := repository.FindUserByID(context.Param("id"))
	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusNotFound, gin.H{"message": "Usuário não encontrado!"})
		return
	}

	currentUser := context.MustGet("user").(models.User)
	if !policies.NewUserPolicy(&currentUser).Delete(user) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Você não tem permissão!"})
		return
	}

	err = repository.DeleteUser(user)
	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Erro: %s", err.Error())})
		return
	}

	context.JSON(http.StatusOK, user)
}

// @Summary Select users
// @Schemes
// @Description Select users
// @Tags users
// @Accept json
// @Produce json
// @Param search query string false "Search by name"
// @Param limit query int false "Limit the number of results"
// @Param offset query int false "Offset the results"
// @Success 200 {array} models.SampleModel "ok"
// @Failure 400 {object} models.APIMessage "Bad request"
// @Failure 404 {object} models.APIMessage "Not found"
// @Router /users_select [get]
// @Security Bearer
func SelectUsers(context *gin.Context) {
	user := context.MustGet("user").(models.User)
	if !policies.NewUserPolicy(&user).ViewAny() {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Você não tem permissão!"})
		return
	}

	search := context.Query("search")
	limit, _ := strconv.Atoi(context.Query("limit"))
	offset, _ := strconv.Atoi(context.Query("offset"))

	itemSelects, err := repository.FindUserSelects(search, &limit, &offset)
	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Erro: %s", err.Error())})
		return
	}

	context.JSON(http.StatusOK, itemSelects)
}
