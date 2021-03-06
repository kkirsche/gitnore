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
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize gitnore's configuration directory and file",
	Long: `Initialize gitnore's configuration directory and file. If a
Github personal access token has been provided via the token flag, write this
to the configuration file, otherwise prompt the user for their personal access
token.

In the event that a configuration file already exists at this location, init
will not overwrite the file.

Usage:
$ gitnore init -t="my personal access token"

or

$ gitnore init
Enter Personal Access Token:
`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Printf("failed to create configuration path with error: %s\n", err)
			os.Exit(1)
		}

		exists := fileExists(cfgFile)
		if exists {
			fmt.Printf("configuration file %s already exists, use writeConfig command to overwrite\n", cfgFile)
			os.Exit(1)
		}

		_, err = os.OpenFile(cfgFile, os.O_RDONLY|os.O_CREATE, 0755)
		if err != nil {
			fmt.Printf("failed to create configuration file with error: %s", err)
			os.Exit(1)
		}

		token := viper.GetString("token")
		if token == "" {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter Personal Access Token: ")
			token, err = reader.ReadString('\n')
			if err != nil {
				fmt.Printf("failed to read personal access token from os.Stdin with error: %s\n", err)
				os.Exit(1)
			}
		}
		token = strings.TrimSpace(token)

		viper.SetDefault("token", token)
		err = viper.WriteConfig()
		if err != nil {
			fmt.Printf("failed to write configuration with error: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

// fileExists reports whether the named file or directory exists.
func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
