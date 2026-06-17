package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var Verify2FACmd = &cobra.Command{
	Use:   "verify-2fa <otp>",
	Short: "Verify and enable MFA",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {

		if !deps.CLIContext.IsLoggedIn() {
			return fmt.Errorf("please login first")
		}

		user := deps.CLIContext.CurrentUser

		if user == nil {
			return fmt.Errorf("user not found")
		}

		secret := deps.CLIContext.PendingMFASecret

		if secret == "" {
			return fmt.Errorf(
				"run enable-2fa first",
			)
		}

		code := strings.TrimSpace(
			args[0],
		)

		valid := deps.TOTPService.ValidateCode(
			code,
			secret,
		)

		if !valid {
			return fmt.Errorf(
				"invalid otp code",
			)
		}

		if err := deps.AuthService.Enable2FA(
			user.ID,
			secret,
		); err != nil {
			return err
		}

		user.MFAEnabled = true
		user.MFASecret = &secret

		deps.CLIContext.ClearPendingMFASecret()

		fmt.Println()
		fmt.Println(
			"✓ MFA enabled successfully",
		)

		return nil
	},
}
