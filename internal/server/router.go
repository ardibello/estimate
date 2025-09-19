package server

import (
	"github.com/ardibello/estimate/pkg/gen/openapi"
	pkgmiddleware "github.com/ardibello/estimate/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(server openapi.ServerInterface, swagger *openapi3.T) (*echo.Echo, error) {
	e := echo.New()

	// Hide the default console output from echo
	e.HideBanner = true
	e.HidePort = true

	// middlewares
	e.Use(middleware.RequestID())
	e.Use(pkgmiddleware.OApiValidatorMiddleware(swagger))
	e.Use(pkgmiddleware.TimeoutMiddleware)
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.HTTPErrorHandler = pkgmiddleware.CustomHTTPErrorHandler

	// register routes
	openapi.RegisterHandlers(e, server)

	return e, nil
}
