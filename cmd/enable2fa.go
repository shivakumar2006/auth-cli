package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Enable2FACmd = &cobra.Command{
	Use:   "enable-2fa",
	Short: "Enable 2FA for current user",
	RunE: func(cmd *cobra.Command, args []string) error {

		if !deps.CLIContext.IsLoggedIn() {
			return fmt.Errorf("please login first")
		}

		user := deps.CLIContext.CurrentUser

		if user == nil {
			return fmt.Errorf("user not found")
		}

		// generate secret
		secret, otpURL, err := deps.TOTPService.GenerateSecret(
			user.Username,
		)
		if err != nil {
			return err
		}

		// generate QR code
		qr, err := deps.TOTPService.GenerateQRCode(
			user.Username,
			otpURL,
		)
		if err != nil {
			return err
		}

		fmt.Println()
		fmt.Println("===================================")
		fmt.Println("2FA Setup")
		fmt.Println("===================================")

		fmt.Printf("QR Code: %s\n", qr)
		fmt.Printf("Secret : %s\n", secret)

		fmt.Println()
		fmt.Println("Scan the QR code using:")
		fmt.Println("- Google Authenticator")
		fmt.Println("- Microsoft Authenticator")
		fmt.Println("- Authy")

		fmt.Println()
		fmt.Print("Enter OTP to verify setup: ")

		reader := bufio.NewReader(os.Stdin)

		code, err := reader.ReadString('\n')
		if err != nil {
			return nil
		}

		code = strings.TrimSpace(code)

		valid := deps.TOTPService.ValidateCode(code, secret)

		if !valid {
			return fmt.Errorf("invalid otp code")
		}

		// save secret in db
		if err := deps.AuthService.Enable2FA(user.ID, secret); err != nil {
			return err
		}

		// update runtime context
		user.MFAEnabled = true
		user.MFASecret = &secret

		fmt.Println()
		fmt.Println("✓ MFA enabled successfully")

		return nil
	},
}
