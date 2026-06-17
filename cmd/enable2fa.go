package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Enable2FACmd = &cobra.Command{
	Use:   "enable-2fa",
	Short: "Start MFA setup process",

	RunE: func(cmd *cobra.Command, args []string) error {

		if !deps.CLIContext.IsLoggedIn() {
			return fmt.Errorf("please login first")
		}

		user := deps.CLIContext.CurrentUser

		if user == nil {
			return fmt.Errorf("user not found")
		}

		if user.MFAEnabled {
			return fmt.Errorf("mfa already enabled")
		}

		secret, otpURL, err := deps.TOTPService.GenerateSecret(
			user.Username,
		)
		if err != nil {
			return err
		}

		qrPath, err := deps.TOTPService.GenerateQRCode(
			user.Username,
			otpURL,
		)
		if err != nil {
			return err
		}

		deps.CLIContext.SetPendingMFASecret(
			secret,
		)

		fmt.Println()
		fmt.Println("===================================")
		fmt.Println("2FA Setup")
		fmt.Println("===================================")

		fmt.Printf("QR Code : %s\n", qrPath)
		fmt.Printf("Secret  : %s\n", secret)

		fmt.Println()
		fmt.Println("Scan QR code in Authenticator app")
		fmt.Println()

		fmt.Println(
			"Next step:",
		)

		fmt.Println(
			"verify-2fa <otp>",
		)

		return nil
	},
}
