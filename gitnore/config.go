package gitnore

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
)

// Configuration represents the gitnore configuration
type Configuration struct {
	// Token is
	GithubAccessToken string `mapstructure:"token" yaml:"token" json:"token" toml:"token"`
	githubAPIClient   *github.Client
	ctx               context.Context
}

// validate is used to check whether the current configuration is valid or not
func (c *Configuration) validate() error {
	if c.GithubAccessToken == "" {
		return errors.New("github access token is a required configuration value")
	}

	return nil
}
