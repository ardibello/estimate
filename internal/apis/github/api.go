package github

import (
	"crypto/rsa"
	"net/http"
	"time"
)

type API struct {
	githubBaseURL        string
	githubClientID       string
	githubInstallationID string
	client               *http.Client
	githubPrivateKey     *rsa.PrivateKey
}

func NewAPI(
	githubBaseURL, githubClientID, githubInstallationID string,
	githubPrivateKey *rsa.PrivateKey,
) *API {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &API{
		githubBaseURL:        githubBaseURL,
		githubClientID:       githubClientID,
		client:               client,
		githubPrivateKey:     githubPrivateKey,
		githubInstallationID: githubInstallationID,
	}
}
