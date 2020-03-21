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
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func genericPasswordExists(account string) (*bool, error) {
	commandStr := fmt.Sprintf("/usr/bin/security find-generic-password -a %s -s %s", account, service)
	args := strings.Split(commandStr, " ")
	command := exec.Command(args[0], args[1:]...)
	b, _ := command.CombinedOutput()
	result := strings.Contains(fmt.Sprintf("%s", b), "attributes:")
	return &result, nil
}

func deleteGenericPassword(account string) error {
	commandStr := fmt.Sprintf("/usr/bin/security delete-generic-password -a %s -s %s", account, service)
	args := strings.Split(commandStr, " ")
	command := exec.Command(args[0], args[1:]...)
	b, _ := command.CombinedOutput()
	result := strings.Contains(fmt.Sprintf("%s", b), "password has been deleted")
	if result != true {
		return errors.New(fmt.Sprintf("Password could not be deleted because it does not exist for %v", account ))
	}
	return nil
}

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configures the account in Security Keychain",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Invalid number of arguments. Please pass in the account and password.")
			os.Exit(1)
		}

		genericPasswordExists, err := genericPasswordExists(args[0])
		if err != nil {
			log.Printf("Getting generic password failed: %v", err)
			os.Exit(1)
		}

		if *genericPasswordExists == true {
			reader := bufio.NewReader(os.Stdin)
			var answer string
			fmt.Println("A generic password already exists for this account, enter y if you want to overwrite it? ")
			answer, _ = reader.ReadString('\n')
			answerTrimmed := strings.TrimSpace(answer)
			if answerTrimmed != "y" {
				fmt.Printf("Aborting adding the generic password for %s", args[0])
				os.Exit(0)
			}

			err := deleteGenericPassword(args[0])
			if err != nil {
				fmt.Printf("Failed deleting generic password for %s", args[0])
				os.Exit(1)
			}
		}

		commandStr := fmt.Sprintf("/usr/bin/security add-generic-password -a %s -s %s -w %s", args[0], service, args[1])
		commmandArgs := strings.Split(commandStr, " ")
		command := exec.Command(commmandArgs[0], commmandArgs[1:]...)
		_, err = command.CombinedOutput()
		if err != nil {
			log.Printf("Running security failed: %v", err)
		}

		fmt.Printf("Success adding generic password for %s", args[0])
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
