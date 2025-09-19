package config

type Config struct {
	Port                 string `env:"PORT"                            envDefault:"8080"`
	GithubApiURL         string `env:"GITHUB_API_URL"                  envDefault:"https://api.github.com"`
	GithubPrivateKey     string `env:"GITHUB_PRIVATE_KEY,required"`
	GithubClientID       string `env:"GITHUB_CLIENT_ID,required"`
	GithubInstallationID string `env:"GITHUB_INSTALLATION_ID,required"`
}
