package server

import (
	"log/slog"

	"github.com/ardibello/estimate/pkg/gen/openapi"
	"github.com/ardibello/estimate/pkg/logger"
	"github.com/labstack/echo/v4"
)

func (s *EstimatesAPI) PostIssues(c echo.Context) error {
	newIssuePayload := new(openapi.PostIssueRequest)

	err := c.Bind(newIssuePayload)
	if err != nil {
		return err
	}

	logger.Info("PostIssues request",
		slog.String("action", string(newIssuePayload.Action)),
		slog.String("issue.body", newIssuePayload.Issue.Body),
	)

	err = s.app.ProcessNewIssue(c.Request().Context(), newIssuePayload)
	if err != nil {
		return err
	}

	return nil
}
