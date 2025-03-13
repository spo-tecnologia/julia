package controllers

import (
	"net/http"

	"github.com/OdairPianta/julia/models"
	"github.com/gin-gonic/gin"
)

// @BasePath /api

// @Summary Get api status
// @Schemes
// @Description Get api status
// @Tags status
// @Accept json
// @Produce json
// @Success 200 {object} models.APIStatus "ok"
// @Failure 400 {object} models.APIMessage "Bad request"
// @Failure 404 {object} models.APIMessage "Not found"
// @Router /status [get]
func FindStatus(context *gin.Context) {
	var apiStatus = models.APIStatus{
		Status: "ok",
	}

	context.JSON(http.StatusOK, apiStatus)
}
