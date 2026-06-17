package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var Disable2FACmd = &cobra.Command{
	Use:   "disable-2fa <confirmation>",
	Short: "Disable MFA for current user",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {

		if !deps.CLIContext.IsLoggedIn() {
			return fmt.Errorf("please login first")
		}

		user := deps.CLIContext.CurrentUser

		if user == nil {
			return fmt.Errorf("user not found in context")
		}

		if !user.MFAEnabled {
			return fmt.Errorf("mfa is already disabled")
		}

		confirmation := strings.TrimSpace(
			strings.ToLower(args[0]),
		)

		if confirmation != "yes" {
			fmt.Println("operation cancelled")
			return nil
		}

		if err := deps.AuthService.Disable2FA(
			user.ID,
		); err != nil {
			return err
		}

		// Update runtime context
		user.MFAEnabled = false
		user.MFASecret = nil

		fmt.Println()
		fmt.Println("✓ MFA disabled successfully")

		return nil
	},
}
