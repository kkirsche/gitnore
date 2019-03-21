package gitnore

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/google/go-github/github"
)

// GetRepositoryFileContents is used to retrieve information about a repository's file contents
func GetRepositoryFileContents(ctx context.Context, client *github.Client, path string) {
	file, _, _, err := client.Repositories.GetContents(ctx, "github", "gitignore", path, nil)
	if err != nil {
		fmt.Printf("error occurred while retrieving the github/gitignore repository file contents: %s", err)
	}

	c, err := base64.StdEncoding.DecodeString(*file.Content)
	if err != nil {
		fmt.Printf("error occurred while decoding file contents: %s", err)
	}
	fmt.Println(string(c))
}
