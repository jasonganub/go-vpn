/*
Copyright © 2020 Jason Ganub <jasonganub@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

const service = "go-vpn"

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configures the Google Authenticator account",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Invalid number of arguments. Please pass in the account and password.")
			os.Exit(1)
		}
		commandStr := fmt.Sprintf("/usr/bin/security add-generic-password -a %s -s %s -w %s", args[0], service, args[1])
		args = strings.Split(commandStr, " ")
		command := exec.Command(args[0], args[1:]...)
		b, err := command.CombinedOutput()
		if err != nil {
			log.Printf("Running security failed: %v", err)
		}

		fmt.Printf("%s\n", b)
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
