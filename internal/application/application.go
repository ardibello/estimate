//go:generate go run go.uber.org/mock/mockgen -package=applicationtest -source=application.go -destination=applicationtest/application.go .

package application

import (
	"context"

	"github.com/ardibello/estimate/pkg/gen/openapi"
)

// EstimatesApp represents the business logic layer of the delegations API.
type EstimatesApp interface {
	ProcessNewIssue(ctx context.Context, newIssue *openapi.PostIssueRequest) error
}
