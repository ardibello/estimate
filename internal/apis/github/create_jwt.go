package github

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (a *API) creteJWT() (string, error) {
	// create and sign jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iat": jwt.NewNumericDate(time.Now().Add(-60 * time.Second)),
		"exp": jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		"iss": a.githubClientID,
	})

	signedToken, err := token.SignedString(a.githubPrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to create signed jwt: %w", err)
	}

	return signedToken, nil
}
