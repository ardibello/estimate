package main

import (
	"errors"
	"io/fs"
	"log/slog"

	"github.com/ardibello/estimate/internal/apis/github"
	"github.com/ardibello/estimate/internal/application/issues"
	"github.com/ardibello/estimate/internal/config"
	"github.com/ardibello/estimate/internal/server"
	"github.com/ardibello/estimate/pkg/gen/openapi"
	"github.com/ardibello/estimate/pkg/logger"
	"github.com/caarlos0/env/v11"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func main() {
	// setup logger
	logger.Init()

	// load env vars from .env files
	err := godotenv.Load()
	if err != nil {
		var pathErr *fs.PathError
		if !errors.As(err, &pathErr) {
			logger.Fatal("failed to read env file", err)

			return
		}
	}

	// load config
	cfg, err := env.ParseAs[config.Config]()
	if err != nil {
		logger.Fatal("failed to load config", err)

		return
	}

	swagger, err := openapi.GetSwagger()
	if err != nil {
		logger.Fatal("error loading swagger spec", err)

		return
	}

	// read GitHub issued private key
	githubPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(cfg.GithubPrivateKey))
	if err != nil {
		logger.Fatal("error loading github issued private key", err)

		return
	}

	// initialize service
	svc := server.NewEstimatesAPI(
		issues.NewApp(
			github.NewAPI(
				cfg.GithubApiURL,
				cfg.GithubClientID,
				cfg.GithubInstallationID,
				githubPrivateKey,
			),
		),
	)

	// initialize router
	router, err := server.NewRouter(svc, swagger)
	if err != nil {
		logger.Fatal("failed to initialize router", err)

		return
	}

	// start the http server
	logger.Info("starting server", slog.String("port", cfg.Port))

	if err = router.Start(":" + cfg.Port); err != nil {
		logger.Fatal("failed to start http server", err)
	}
}
