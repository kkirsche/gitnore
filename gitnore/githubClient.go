package gitnore

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

// NewGithubAPIClient is used to create a new API client for use with the Github API
func (c *Configuration) NewGithubAPIClient() error {
	err := c.validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate configuration before creating Github API client")
	}

	c.ctx = context.Background()

	sts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.GithubAccessToken},
	)
	tc := oauth2.NewClient(c.ctx, sts)
	c.githubAPIClient = github.NewClient(tc)

	return nil
}
