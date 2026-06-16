package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var WhoAmICmd = &cobra.Command{
	Use:   "whoami",
	Short: "Show current user information",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !deps.CLIContext.IsLoggedIn() {
			return fmt.Errorf("please login first")
		}

		user := deps.CLIContext.CurrentUser
		session := deps.CLIContext.CurrentSession

		if user == nil {
			return fmt.Errorf("user not found in context")
		}

		fmt.Println()
		fmt.Println("Current User")
		fmt.Println("------------")

		fmt.Printf("Username: %s\n", user.Username)

		fmt.Printf(
			"Registered At: %s\n",
			user.CreatedAt.Format("2006-01-02 15:04:05"),
		)

		fmt.Printf(
			"MFA Enabled: %v\n",
			user.MFAEnabled,
		)

		if user.LastLogin != nil {
			fmt.Printf(
				"Last Login: %s\n",
				user.LastLogin.Format(
					"2006-01-02 15:04:05",
				),
			)
		} else {
			fmt.Println("Last Login: Never")
		}

		if session != nil {
			fmt.Printf(
				"Session Expires: %s\n",
				session.ExpiresAt.Format(
					"2006-01-02 15:04:05",
				),
			)
		}

		fmt.Println()

		return nil
	},
}
