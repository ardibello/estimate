package server_test

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ardibello/estimate/internal/application/applicationtest"
	"github.com/ardibello/estimate/internal/server"
	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
)

type newEchoContextParams struct {
	method  string
	target  string
	payload string
}

// newEchoContext returns an echo.Context.
func newEchoContext(params *newEchoContextParams) (echo.Context, *httptest.ResponseRecorder) {
	var (
		target = "/"
		method string
		body   io.Reader
	)

	if params != nil {
		if params.method != "" {
			method = params.method
		}

		if params.target != "" {
			target = params.target
		}

		if params.payload != "" {
			body = strings.NewReader(params.payload)
		}
	}

	e := echo.New()
	req := httptest.NewRequest(method, target, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	return e.NewContext(req, rec), rec
}

// func newMockService(t *testing.T) (openapi.ServerInterface, application.DelegationsApp) {.
func newMockService(t *testing.T) (*server.EstimatesAPI, applicationtest.MockEstimatesApp) {
	ctrl := gomock.NewController(t)
	app := applicationtest.NewMockEstimatesApp(ctrl)
	svc := server.NewEstimatesAPI(app)

	return svc, *app
}
