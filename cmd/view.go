// Copyright © 2019 Kevin Kirsche <kev.kirsche@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// previewCmd represents the view command
var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Preview the content of a gitignore file",
	Long: `Preview the content of a remote gitignore file from the
Github gitignore repository. This command takes one or more positional
arguments representing the path of the file which you would like to preview.

Note: This command's arguments are case-sensitive

Usage:
$ gitnore preview Go

$ gitnore preview Global/macOS
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := config.NewGithubAPIClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, arg := range args {
			if !strings.Contains(arg, ".gitignore") {
				arg = fmt.Sprintf("%s.gitignore", arg)
			}

			err := config.GetRepositoryFileContents(arg)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(previewCmd)
}
