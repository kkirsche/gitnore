package gitnore

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/github"
)

// GetRepositoryContents is used to retrieve information about a repository's contents
func GetRepositoryContents(ctx context.Context, client *github.Client, path string) {
	_, dir, _, err := client.Repositories.GetContents(ctx, "github", "gitignore", path, nil)
	if err != nil {
		fmt.Printf("error occurred while retrieving the github/gitignore repository contents: %s", err)
	}
	for _, file := range dir {
		if *file.Type == "dir" {
			GetRepositoryContents(ctx, client, *file.Path)
			continue
		}

		if *file.Type == "file" && !strings.Contains(*file.Name, ".gitignore") {
			continue
		}

		*file.Name = strings.TrimSuffix(*file.Name, ".gitignore")

		if path != "" {
			fmt.Printf("%s/%s\n", path, *file.Name)
			continue
		}
		fmt.Printf("%s\n", *file.Name)
	}
}
