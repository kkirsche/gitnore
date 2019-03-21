package gitnore

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// GetRepositoryContents is used to retrieve information about a repository's contents
func (c *Configuration) GetRepositoryContents(path string) error {
	_, dir, _, err := c.githubAPIClient.Repositories.GetContents(c.ctx, "github", "gitignore", path, nil)
	if err != nil {
		return errors.Wrap(err, "error occurred while retrieving the github/gitignore repository contents")
	}

	for _, file := range dir {
		if file.GetType() == "dir" {
			c.GetRepositoryContents(file.GetPath())
			continue
		}

		if file.GetType() == "file" && !strings.Contains(file.GetName(), ".gitignore") {
			continue
		}

		*file.Name = strings.TrimSuffix(*file.Name, ".gitignore")

		if path != "" {
			fmt.Printf("%s/%s\n", path, file.GetName())
			continue
		}

		fmt.Printf("%s\n", file.GetName())
	}

	return nil
}
