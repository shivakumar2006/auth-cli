package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var HelpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show available commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println()
		fmt.Println("Available Commands")
		fmt.Println("==================")

		if !deps.CLIContext.IsLoggedIn() {
			fmt.Println("register     - Create a new user")
			fmt.Println("login        - Login to your account")
			fmt.Println("help         - Show available commands")
			fmt.Println("exit         - Exit application")

			return nil
		}

		fmt.Println("whoami       - Show current user details")
		fmt.Println("enable-2fa   - Enable MFA")
		fmt.Println("disable-2fa  - Disable MFA")
		fmt.Println("logout       - Logout current session")
		fmt.Println("help         - Show available commands")
		fmt.Println("exit         - Exit application")

		return nil
	},
}
