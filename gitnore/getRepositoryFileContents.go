package gitnore

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

const (
	gitignore = ".gitignore"
	prefix    = "@"
)

// GetRepositoryFileContents is used to retrieve information about a repository's file contents
func (c *Configuration) GetRepositoryFileContents(path string) error {
	file, _, resp, err := c.githubAPIClient.Repositories.GetContents(c.ctx, "github", "gitignore", path, nil)
	if err != nil {
		return errors.Wrap(err, "error occurred while retrieving the github/gitignore repository file contents")
	}

	if !(resp.StatusCode >= http.StatusOK) && !(resp.StatusCode <= http.StatusIMUsed) {
		return errors.Errorf("error occurred while retrieving file: %s", resp.Status)
	}

	fc, err := file.GetContent()
	if err != nil {
		return errors.Wrap(err, "error occurred while decoding file contents")
	}

	fmt.Println(fc)
	return nil
}
