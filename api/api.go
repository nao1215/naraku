// Package api control naraku api.
package api

import (
	"github.com/labstack/echo/v4"
	// Import docs for Swagger documentation generation
	"github.com/go-spectest/naraku/api/di"
	_ "github.com/go-spectest/naraku/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title			naraku API Swagger
//	@version		0.0.1
//	@description	This is a naraku API server swagger.
//	@termsOfService	https://github.com/go-spectest/naraku

//	@contact.name	Naohiro CHIKAMATSU (Author)
//	@contact.url	https://github.com/go-spectest/naraku/issues
//	@contact.email	n.chika156@gmail

//	@license.name	MIT License
//	@license.url	https://github.com/go-spectest/naraku/blob/main/LICENSE

//	@host		https://nao1215.github.io/naraku/html/index.html
//	@BasePath	/v1

// Run start server.
func Run() error {
	return NewAPI().Start(":8080")
}

// API is a structure that aggregates the necessary information for API execution.
type API struct {
	// echo framework. it's manage api handlers.
	*echo.Echo
	// Naraku is a struct that contains the settings for the naraku.
	naraku *di.Naraku
	// healthController is a controller for /health API.
	healthController *HealthController
	// osController is a controller for /os API.
	osController *OSController
	// ulidController is a controller for /ulid API.
	ulidController *ULIDController
}

// NewAPI return api struct.
func NewAPI() *API {
	api := new(API)
	api.Echo = echo.New()
	api.naraku = di.NewNaraku()
	api.setControllers()
	api.route()
	return api
}

// setControllers set controllers.
// This method is called before routing.
func (a *API) setControllers() {
	a.healthController = NewHealthController()
	a.osController = NewOSController()
	a.ulidController = NewULIDController()
}

// route set routing.
// This method is called after setControllers.
func (a *API) route() {
	a.GET("/v1/health", a.healthController.health)
	a.GET("/v1/os/:distribution", a.osController.handleGetOS)
	a.GET("/v1/ulid", a.ulidController.generate)
	a.GET("/swagger/*", echoSwagger.WrapHandler)
}
