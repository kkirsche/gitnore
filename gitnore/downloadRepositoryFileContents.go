package gitnore

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// DownloadRepositoryFileContents is used to retrieve information about a repository's file contents
// and store it in the .gitignore file
func (c *Configuration) DownloadRepositoryFileContents(path string) error {
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

	if fileExists(gitignore) {
		b, err := ioutil.ReadFile(gitignore)
		if err != nil {
			return errors.Wrap(err, "error occurred while reading gitignore file")
		}

		bs := string(b)
		if strings.Contains(bs, fmtCommentLine(file.GetPath())) && strings.Contains(bs, fmtCommentLine(file.GetHTMLURL())) {
			return errors.Errorf("file type '%s' already exists in gitignore file", file.GetPath())
		}

		f, err := os.OpenFile(gitignore, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return errors.Wrap(err, "error occurred while operning gitignore file")
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
		f.WriteString(fc)
	} else {
		f, err := os.Create(gitignore)
		if err != nil {
			return errors.Wrap(err, "error occurred while creating gitignore file")
		}
		defer f.Close()

		items := []string{
			file.GetPath(),
			file.GetHTMLURL(),
		}

		for _, item := range items {
			f.WriteString(fmtCommentLine(item))
		}
		f.WriteString(fc)
	}
	return nil
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
