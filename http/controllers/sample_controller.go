package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/OdairPianta/julia/config"
	"github.com/OdairPianta/julia/exports"
	"github.com/OdairPianta/julia/http/requests"
	"github.com/OdairPianta/julia/models"
	"github.com/OdairPianta/julia/policies"
	"github.com/OdairPianta/julia/repository"
	"github.com/OdairPianta/julia/services"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

// @Summary Get samples
// @Schemes
// @Description Get samples
// @Tags sample_models
// @Accept json
// @Produce json
// @Param search query string false "Search by id or sample_string or sample_unique or sample_nullable"
// @Success 200 {array} models.SampleModel "ok"
// @Failure 400 {object} models.APIMessage "Bad request"
// @Failure 404 {object} models.APIMessage "Not found"
// @Router /samples [get]
// @Security Bearer
func FindSamples(context *gin.Context) {
	user := context.MustGet("user").(models.User)
	if !policies.NewSamplePolicy(&user).ViewAny() {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Você não tem permissão!"})
		return
	}

	search := context.Query("search")
	limit, _ := strconv.Atoi(context.Query("limit"))
	offset, _ := strconv.Atoi(context.Query("offset"))

	sampleModels, err := repository.FindSamples(search, &limit, &offset)
	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Erro: %s", err.Error())})
		return
	}

	context.JSON(http.StatusOK, sampleModels)
}

// @Summary Get sample_model
// @Schemes
// @Description Get sample_model
// @Tags sample_models
// @Accept json
// @Produce json
// @Param   id     path    int     true        "SampleModel ID"
// @Success 200 {object} models.SampleModel "ok"
// @Failure 400 {object} models.APIMessage "Bad request"
// @Failure 404 {object} models.APIMessage "Not found"
// @Router /sample_models/{id} [get]
// @Security Bearer
func FindSample(context *gin.Context) {
	model, err := repository.FindSampleByID(context.Param("id"))
	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Erro: %s", err.Error())})
		return
	}

	user := context.MustGet("user").(models.User)
	if !policies.NewSamplePolicy(&user).View(model) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Você não tem permissão!"})
		return
	}

	context.JSON(http.StatusOK, model)
}

// @Summary		Add an sample_model
// @Description	add by json sample_model
// @Tags			sample_models
// @Accept		json
// @Produce		json
// @Param			sample_model	body		requests.CreateSampleModelInput	true	"Add sample_model"
// @Success		200		{object}	models.SampleModel
// @Failure 400 {object} models.APIMessage "Bad request"
// @Failure 404 {object} models.APIMessage "Not found"
// @Router			/sample_models [post]
// @Security Bearer
func CreateSampleModel(context *gin.Context) {
	user := context.MustGet("user").(models.User)
	if !policies.NewSamplePolicy(&user).Create() {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Você não tem permissão!"})
		return
	}

	var input requests.CreateSampleModelInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": services.GetErrosMessage(err)})
		return
	}

	var existsSampleModel models.SampleModel
	err := config.DB.Where("sample_unique = ?", input.SampleUnique).First(&existsSampleModel).Error
	if err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Sample Model já cadastrado!"})
		return
	}

	sample_model := models.SampleModel{
		Name:           input.Name,
		SampleString:   input.SampleString,
		SampleUnique:   input.SampleUnique,
		SampleDate:     time.Now(),
		SampleNullable: input.SampleNullable,
		SampleDouble:   0.0,
		SampleDetailID: 1,
	}
	err = repository.CreateSampleModel(&sample_model)

	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Erro: %s", err.Error())})
		return
	}

	context.JSON(http.StatusOK, sample_model)
}

// @Summary		Update an sample_model
// @Description	Update by json sample_model
// @Tags			sample_models
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"ID"
// @Param			sample_model	body		requests.UpdateSampleModelInput	true	"Update sample_model"
// @Success		200		{object}	models.SampleModel
// @Failure 400 {object} models.APIMessage "Bad request"
// @Failure 404 {object} models.APIMessage "Not found"
// @Router			/sample_models/{id} [put]
// @Security Bearer
func UpdateSampleModel(context *gin.Context) {
	model, err := repository.FindSampleByID(context.Param("id"))
	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusNotFound, gin.H{"message": "Sample Model não encontrado!"})
		return
	}

	user := context.MustGet("user").(models.User)
	if !policies.NewSamplePolicy(&user).Update(model) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Você não tem permissão!"})
		return
	}

	var input requests.UpdateSampleModelInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": services.GetErrosMessage(err)})
		return
	}

	updatedSampleModel := models.SampleModel{
		Name:           input.Name,
		SampleString:   input.SampleString,
		SampleUnique:   input.SampleUnique,
		SampleDate:     input.SampleDate,
		SampleNullable: input.SampleNullable,
		SampleDouble:   input.SampleDouble,
		SampleDetailID: input.SampleDetailID,
		OrderNumber:    input.OrderNumber,
	}

	err = repository.UpdateSampleModel(model, &updatedSampleModel)

	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Erro: %s", err.Error())})
		return
	}

	context.JSON(http.StatusOK, model)
}

// @Summary		Delete an sample_model
// @Description	Delete by sample_model ID
// @Tags			sample_models
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"ID"	Format(int64)
// @Success		200	{object}	models.SampleModel
// @Failure 400 {object} models.APIMessage "Bad request"
// @Failure 404 {object} models.APIMessage "Not found"
// @Router			/sample_models/{id} [delete]
// @Security Bearer
func DeleteSampleModel(context *gin.Context) {
	model, err := repository.FindSampleByID(context.Param("id"))
	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusNotFound, gin.H{"message": "Sample Model não encontrado!"})
		return
	}

	user := context.MustGet("user").(models.User)
	if !policies.NewSamplePolicy(&user).Delete(model) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Você não tem permissão!"})
		return
	}

	err = repository.DeleteSampleModel(model)
	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Erro: %s", err.Error())})
		return
	}

	context.JSON(http.StatusOK, model)
}

// @Summary Select samples
// @Schemes
// @Description Select samples
// @Tags sample_models
// @Accept json
// @Produce json
// @Param search query string false "Search by sample_string"
// @Param limit query int false "Limit the number of results"
// @Param offset query int false "Offset the results"
// @Success 200 {array} models.Assessment "ok"
// @Failure 400 {object} models.APIMessage "Bad request"
// @Failure 404 {object} models.APIMessage "Not found"
// @Router /sample_models_select [get]
// @Security Bearer
func SelectSamples(context *gin.Context) {
	user := context.MustGet("user").(models.User)
	if !policies.NewSamplePolicy(&user).ViewAny() {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Você não tem permissão!"})
		return
	}

	search := context.Query("search")
	limit, _ := strconv.Atoi(context.Query("limit"))
	offset, _ := strconv.Atoi(context.Query("offset"))

	itemSelects, err := repository.FindSampleItemSelects(search, &limit, &offset)
	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Erro: %s", err.Error())})
		return
	}

	context.JSON(http.StatusOK, itemSelects)
}

// @Summary Duplicate Sample
// @Description Duplicate an existing sample.
// @Tags sample_models
// @Accept json
// @Produce json
// @Param id path int true "Sample ID"
// @Success 200 {object} models.SampleModel "ok"
// @Failure 400 {object} models.APIMessage "Bad request"
// @Failure 401 {object} models.APIMessage "Unauthorized"
// @Router /sample_models/{id}/duplicate [post]
// @Security Bearer
func DuplicateSampleModel(c *gin.Context) {
	sample, err := repository.FindSampleByID(c.Param("id"))
	if err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusNotFound, gin.H{"message": "Exemplo não encontrada!"})
		return
	}

	user := c.MustGet("user").(models.User)
	if !policies.NewSamplePolicy(&user).Update(sample) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Você não tem permissão!"})
		return
	}
	sample.ID = 0
	sample.Name = sample.Name + " - (Cópia)"
	sample.SampleUnique = sample.SampleUnique + " - (Cópia)"
	sample.CreatedAt = time.Now()
	sample.UpdatedAt = time.Now()

	err = repository.CreateSampleModel(sample)
	if err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Erro ao duplicar sample model: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, sample)
}

// @Summary Export samples
// @Schemes
// @Description Export samples
// @Tags sample_models
// @Accept json
// @Produce json
// @Param search query string false "Search by id or sample_string or sample_unique or sample_nullable"
// @Success 200 {array} models.APIUrl "url"
// @Failure 400 {object} models.APIMessage "Bad request"
// @Failure 404 {object} models.APIMessage "Not found"
// @Router /samples/export [get]
// @Security Bearer
func ExportSamples(context *gin.Context) {
	user := context.MustGet("user").(models.User)
	if !policies.NewSamplePolicy(&user).ViewAny() {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Você não tem permissão!"})
		return
	}

	search := context.Query("search")
	limit, _ := strconv.Atoi(context.Query("limit"))
	offset, _ := strconv.Atoi(context.Query("offset"))

	sampleModels, err := repository.FindSamples(search, &limit, &offset)
	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Erro: %s", err.Error())})
		return
	}

	fileName, err := exports.ExportSampleModels(sampleModels)
	if err != nil {
		sentry.CaptureException(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Erro ao exportar: %s", err.Error())})
		return
	}
	appUrl := os.Getenv("APP_URL")
	fileURL := fmt.Sprintf("%s/api/storage/temp/%s", appUrl, filepath.Base(fileName))

	response := models.APIUrl{URL: fileURL}
	context.JSON(http.StatusOK, response)
}
