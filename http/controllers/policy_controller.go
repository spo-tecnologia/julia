package controllers

import (
	"net/http"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func FindPolicyPrivacy(context *gin.Context) {
	context.Header("Content-Type", "text/html; charset=utf-8")
	doc, err := os.ReadFile("docs/policy_privacy.html")
	if err != nil {
		sentry.CaptureException(err)
		context.String(http.StatusInternalServerError, "Error: %s", err.Error())
		return
	}
	context.String(http.StatusOK, string(doc))
}

func FindPolicyDeleteUserData(context *gin.Context) {
	context.Header("Content-Type", "text/html; charset=utf-8")
	doc, err := os.ReadFile("docs/policy_delete_user_data.html")
	if err != nil {
		sentry.CaptureException(err)
		context.String(http.StatusInternalServerError, "Error: %s", err.Error())
		return
	}
	context.String(http.StatusOK, string(doc))
}
