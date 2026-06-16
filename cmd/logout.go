package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var LogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from the current session",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !deps.CLIContext.IsLoggedIn() {
			fmt.Println("No active session")
			return nil
		}

		if deps.CLIContext.CurrentSession == nil {
			return fmt.Errorf("session not found")
		}

		err := deps.SessionManager.DestroySession(
			deps.CLIContext.CurrentSession.Token,
		)
		if err != nil {
			return nil
		}

		deps.CLIContext.Logout()

		fmt.Println("✓ Logged out successfully")

		return nil
	},
}
