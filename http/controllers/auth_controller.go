package controllers

import (
	"net/http"

	"github.com/OdairPianta/julia/config"
	"github.com/OdairPianta/julia/http/requests"
	"github.com/OdairPianta/julia/models"
	"github.com/OdairPianta/julia/notifications"
	"github.com/OdairPianta/julia/repository"
	"github.com/OdairPianta/julia/services"
	"github.com/OdairPianta/julia/services/token"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Summary		Login an user
// @Description	Login by json user
// @Tags			users
// @Accept		json
// @Produce		json
// @Param			user	body		requests.LoginInput	true	"Add user"
// @Success		200		{object}	models.User
// @Failure 400 {object} models.APIMessage "Bad request"
// @Failure 404 {object} models.APIMessage "Not found"
// @Router			/login [post]
func Login(context *gin.Context) {

	var input requests.LoginInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": services.GetErrosMessage(err)})
		return
	}

	user := models.User{}

	user.Email = input.Email
	user.Password = input.Password

	token, loginErr := token.LoginCheck(user.Email, user.Password)

	if loginErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Email ou senha incorretos!"})
		return
	}
	registeredUser, errGetUser := repository.FindUserByEmail(input.Email)
	if errGetUser != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Usuário não encontrado"})
		return
	}
	registeredUser.Token = token

	context.JSON(http.StatusOK, registeredUser)
}

// @Summary		Forgot password
// @Description	Forgot password by json user
// @Tags			users
// @Accept		json
// @Produce		json
// @Param			user	body		requests.ForgotPasswordInput	true	"Forgot password"
// @Success		200		{object}	string
// @Failure 400 {object} models.APIMessage "Bad request"
// @Failure 404 {object} models.APIMessage "Not found"
// @Router			/forgot_password [post]
func ForgotPassword(context *gin.Context) {
	var input requests.ForgotPasswordInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": services.GetErrosMessage(err)})
		return
	}

	user, err := repository.FindUserByEmail(input.Email)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Usuário não encontrado"})
		return
	}

	token, err := token.GenerateToken(user.ID)
	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao gerar o token do usuário"})
		return
	}
	user.ResetPasswordToken = token

	err = repository.UpdateUser(user, &requests.UpdateUserInput{
		ResetPasswordToken: user.ResetPasswordToken,
	})
	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao atualizar usuário"})
		return
	}

	notification := notifications.ForgotPasswordNotification{User: *user}
	err = notification.Send()

	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao enviar email!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Email enviado com sucesso!"})
}

// @Summary		Recover password
// @Description	Recover password by json user
// @Tags			users
// @Accept		json
// @Produce		json
// @Param			user	body		requests.RecoverPasswordInput	true	"Recover password"
// @Success		200		{object}	models.User
// @Failure 400 {object} models.APIMessage "Bad request"
// @Failure 404 {object} models.APIMessage "Not found"
// @Router			/recover_password [post]
func RecoverPassword(context *gin.Context) {
	var input requests.RecoverPasswordInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": services.GetErrosMessage(err)})
		return
	}

	user := models.User{}

	user.ResetPasswordToken = input.ResetPasswordToken
	user.Password = input.Password

	err := config.DB.Where("reset_password_token = ?", user.ResetPasswordToken).First(&user).Error
	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusNotFound, gin.H{"message": "Token não encontrado"})
		return
	}

	hashedPassword, errEncryptPass := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if errEncryptPass != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao criptografar senha!"})
		return
	}

	user.Password = string(hashedPassword)
	user.ResetPasswordToken = ""

	err = config.DB.Model(&user).Updates(map[string]interface{}{"password": user.Password, "reset_password_token": user.ResetPasswordToken}).Error
	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao atualizar usuário"})
		return
	}

	token, loginErr := token.LoginCheck(user.Email, input.Password)

	if loginErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "O email ou a senha estão incorretos."})
		return
	}

	user.Token = token

	context.JSON(http.StatusOK, user)
}
