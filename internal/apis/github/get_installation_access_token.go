package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const githubAPIVersion = "2022-11-28"

func (a *API) getInstallationAccessToken(jwtToken string) (string, error) {
	url := fmt.Sprintf(
		"%s/app/installations/%s/access_tokens",
		a.githubBaseURL,
		a.githubInstallationID,
	)

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+jwtToken)
	req.Header.Set("X-Github-Api-Version", githubAPIVersion)

	resp, err := a.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make http request or receive resonse: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Token string `json:"token"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Token, nil
}
