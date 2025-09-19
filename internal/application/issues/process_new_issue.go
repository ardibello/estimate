package issues

import (
	"context"

	"github.com/ardibello/estimate/pkg/gen/openapi"
)

func (s *App) ProcessNewIssue(ctx context.Context, newIssue *openapi.PostIssueRequest) error {
	if ContainsEstimate(newIssue.Issue.Body) {
		return nil
	}

	err := s.githubAPI.AuthenticateAndWriteComment(
		newIssue.Repository.FullName,
		newIssue.Issue.Number,
	)
	if err != nil {
		return err
	}

	return nil
}
