package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Disable2FACmd = &cobra.Command{
	Use:   "disable-2fa",
	Short: "Disable 2fa for current user",
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

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Are you sure you want to disable MFA? (yes/no): ")

		confirmation, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		confirmation = strings.TrimSpace(strings.ToLower(confirmation))

		if confirmation != "yes" {
			fmt.Errorf("operation cancelled")
			return nil
		}

		if err := deps.AuthService.Disable2FA(user.ID); err != nil {
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
