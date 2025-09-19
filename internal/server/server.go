package server

import (
	"github.com/ardibello/estimate/internal/application"
)

type EstimatesAPI struct {
	app application.EstimatesApp // app layer (business logic)
}

func NewEstimatesAPI(appLayer application.EstimatesApp) *EstimatesAPI {
	return &EstimatesAPI{
		app: appLayer,
	}
}
