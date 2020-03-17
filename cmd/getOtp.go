/*
Copyright © 2020 Jason Ganub <jasonganub@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"strings"
)

// getOtpCmd represents the getOtp command
var getOtpCmd = &cobra.Command{
	Use:   "getOtp",
	Short: "Gets the OTP from the account in Security Keychain",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Invalid number of arguments. Please pass in the account")
			os.Exit(1)
		}

		commandStr := fmt.Sprintf("/usr/bin/security find-generic-password -a %s -s %s -w", args[0], service)
		args = strings.Split(commandStr, " ")
		command := exec.Command(args[0], args[1:]...)
		otpKey, err := command.CombinedOutput()
		if err != nil {
			log.Printf("Getting OTP Key failed: %v", err)
		}

		otpKeyStr := strings.Replace(fmt.Sprintf("%s", otpKey), "\n", "", 1)
		commandStr = fmt.Sprintf("/usr/local/bin/oathtool --totp -b %s", otpKeyStr)
		args = strings.Split(commandStr, " ")
		command = exec.Command(args[0], args[1:]...)
		otp, err := command.CombinedOutput()
		if err != nil {
			log.Printf("Getting OTP failed: %v", err)
		}

		fmt.Printf("%s", otp)
	},
}

func init() {
	rootCmd.AddCommand(getOtpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getOtpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getOtpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
