package controllers

import (
	"fmt"
	"net/http"

	"github.com/OdairPianta/julia/config"
	"github.com/OdairPianta/julia/http/requests"
	"github.com/OdairPianta/julia/models"
	"github.com/OdairPianta/julia/repository"
	"github.com/OdairPianta/julia/services"
	"github.com/gin-gonic/gin"
)

// @Summary Get files
// @Description Get all files
// @Tags files
// @Accept  json
// @Produce  json
// @Success 200 {array} models.File
// @Failure 400 {object} models.APIMessage "Bad request"
// @Router /files [get]
// @Security Bearer
func FindFiles(c *gin.Context) {
	files, err := repository.FindFiles()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Arquivos não encontrados"})
		return
	}
	c.JSON(http.StatusOK, files)
}

// @Summary Create file
// @Description Create file by base64 or public url
// @Tags files
// @Accept  json
// @Produce  json
// @Param file body models.File true "File"
// @Success 200 {object} models.File
// @Failure 400 {object} models.APIMessage "Bad request"
// @Router /files [post]
// @Security Bearer
func CreateFile(context *gin.Context) {
	var input requests.CreateFileInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": services.GetErrosMessage(err)})
		return
	}

	file := models.File{
		Extension: input.Extension,
		Path:      input.Path,
		Name:      input.Name,
	}

	if input.PublicURL == "" {
		publicURL, err := models.SaveGryphonAPI(input.Base64, input.Extension, input.Path)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao salvar na API Gryphon: " + err.Error()})
			return
		}
		file.PublicURL = publicURL
	} else {
		file.PublicURL = input.PublicURL
	}
	if err := repository.CreateFile(&file); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Erro: %s", err.Error())})
		return
	}

	context.JSON(http.StatusOK, file)
}

// @Summary Delete file
// @Description Delete file by id
// @Tags files
// @Accept  json
// @Produce  json
// @Param id path int true "File ID"
// @Success 200 {object} models.File
// @Failure 400 {object} models.APIMessage "Bad request"
// @Router /files/{id} [delete]
// @Security Bearer
func DeleteFile(c *gin.Context) {
	var file models.File
	if err := config.DB.Where("id = ?", c.Param("id")).First(&file).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Arquivo não encontrado"})
		return
	}

	if err := repository.DeleteFile(&file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Erro: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, file)
}
