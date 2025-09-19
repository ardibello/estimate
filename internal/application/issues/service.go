package issues

import "github.com/ardibello/estimate/internal/apis"

// App is a struct that holds the application/business layer dependencies.
type App struct {
	githubAPI apis.Github
}

func NewApp(githubAPI apis.Github) *App {
	return &App{
		githubAPI: githubAPI,
	}
}
