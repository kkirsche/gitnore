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

// writeConfigCmd represents the writeConfig command
var writeConfigCmd = &cobra.Command{
	Use:   "writeConfig",
	Short: "Write a new configuration file (unsafe)",
	Long: `The writeConfig command allows you to write a new configuration file
without having to manually remove an existing configuration file. As such, this
command will overwrite any existing configuration you may have. As such, this is
not considered a "Safe" command.

Personal access token can only be provided via a prompt, the -t / --token flag
is not supported at this time.

Usage:
$ gitnore writeConfig
Enter Personal Access Token:
`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll(path, 0644)
		if err != nil {
			fmt.Printf("failed to create configuration path with error: %s\n", err)
			os.Exit(1)
		}
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Personal Access Token: ")
		token, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error occurred while reading personal access token: %s\n", err)
			os.Exit(1)
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
	rootCmd.AddCommand(writeConfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// writeConfigCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// writeConfigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
