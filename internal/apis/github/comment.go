package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var defaultMissingEstimateComment = map[string]string{
	"body": "Please add a time estimate to this issue.",
}

func (a *API) AuthenticateAndWriteComment(fullRepoName string, issueNumber int) error {
	jwtToken, err := a.creteJWT()
	if err != nil {
		return err
	}

	accessToken, err := a.getInstallationAccessToken(jwtToken)
	if err != nil {
		return err
	}

	err = a.writeComment(fullRepoName, accessToken, issueNumber)
	if err != nil {
		return err
	}

	return nil
}

func (a *API) writeComment(fullRepoName, token string, issueNumber int) error {
	commentURL := fmt.Sprintf(
		"%s/repos/%s/issues/%d/comments",
		a.githubBaseURL,
		fullRepoName,
		issueNumber,
	)

	bodyBytes, _ := json.Marshal(defaultMissingEstimateComment)

	req, err := http.NewRequest(http.MethodPost, commentURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-Github-Api-Version", githubAPIVersion)

	resp, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make http request or receive resonse: %w", err)
	}
	defer resp.Body.Close()

	return nil
}
