package main

import (
	"github.com/OdairPianta/julia/config"
	"github.com/OdairPianta/julia/jobs"
	"github.com/OdairPianta/julia/routes"
	"github.com/OdairPianta/julia/services"
)

// swagger embed files

// @title           Julia API
// @version         0.7.4
// @description     API of Julia 2

// @contact.name   API Support
// @contact.url    https://spotec.app/contato/
// @contact.email  contact@spotec.app

// @host      api.julia.spotec.app
// @BasePath  /api

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	config.InitDatabase()
	go jobs.StartClient()
	services.InitSentry()
	routes.HandleRequest()
}
