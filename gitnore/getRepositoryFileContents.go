package gitnore

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/google/go-github/github"
)

const (
	gitignore = ".gitignore"
	prefix    = "@"
)

// GetRepositoryFileContents is used to retrieve information about a repository's file contents
func GetRepositoryFileContents(ctx context.Context, client *github.Client, path string) {
	file, _, resp, err := client.Repositories.GetContents(ctx, "github", "gitignore", path, nil)
	if err != nil {
		fmt.Printf("error occurred while retrieving the github/gitignore repository file contents: %s\n", err)
		return
	}

	if !(resp.StatusCode >= http.StatusOK) && !(resp.StatusCode <= http.StatusIMUsed) {
		fmt.Printf("error occurred while retrieving file: %s", resp.Status)
		return
	}

	c, err := file.GetContent()
	if err != nil {
		fmt.Printf("error occurred while decoding file contents: %s\n", err)
		return
	}
	fmt.Println(c)
}

// DownloadRepositoryFileContents is used to retrieve information about a repository's file contents
// and store it in the .gitignore file
func DownloadRepositoryFileContents(ctx context.Context, client *github.Client, path string) {
	file, _, resp, err := client.Repositories.GetContents(ctx, "github", "gitignore", path, nil)
	if err != nil {
		fmt.Printf("error occurred while retrieving the github/gitignore repository file contents: %s\n", err)
		return
	}

	if !(resp.StatusCode >= http.StatusOK) && !(resp.StatusCode <= http.StatusIMUsed) {
		fmt.Printf("error occurred while retrieving file: %s", resp.Status)
		return
	}

	c, err := file.GetContent()
	if err != nil {
		fmt.Printf("error occurred while decoding file contents: %s\n", err)
		return
	}

	if fileExists(gitignore) {
		b, err := ioutil.ReadFile(gitignore)
		if err != nil {
			fmt.Printf("error occurred while reading gitignore file: %s\n", err)
			return
		}

		bs := string(b)
		if strings.Contains(bs, fmtCommentLine(file.GetPath())) && strings.Contains(bs, fmtCommentLine(file.GetHTMLURL())) {
			fmt.Printf("file type '%s' already exists in gitignore file\n", file.GetPath())
			return
		}

		f, err := os.OpenFile(gitignore, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("error occurred while operning gitignore file: %s\n", err)
			return
		}
		defer f.Close()

		items := []string{
			"\n",
			file.GetPath(),
			file.GetHTMLURL(),
		}

		for _, item := range items {
			f.WriteString(fmtCommentLine(item))
		}
		f.WriteString(c)
	} else {
		f, err := os.Create(gitignore)
		if err != nil {
			fmt.Printf("error occurred while creating gitignore file: %s\n", err)
			return
		}
		defer f.Close()

		items := []string{
			file.GetPath(),
			file.GetHTMLURL(),
		}

		for _, item := range items {
			f.WriteString(fmtCommentLine(item))
		}
		f.WriteString(c)
	}
}

func fmtCommentLine(s string) string {
	if s == "\n" {
		return s
	}

	return fmt.Sprintf("#%s %s\n", prefix, s)
}

// Exists reports whether the named file or directory exists.
func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
