package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"

	"github.com/spf13/cobra"
)

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with username and password",
	RunE: func(cmd *cobra.Command, args []string) error {

		if deps.CLIContext.IsLoggedIn() {
			fmt.Println("Already logged in")
			return nil
		}

		reader := bufio.NewReader(os.Stdin)

		// username
		fmt.Print("Username: ")

		username, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		username = strings.TrimSpace(username)

		// password
		fmt.Print("Password: ")

		passwordBytes, err := term.ReadPassword(
			int(os.Stdin.Fd()),
		)
		if err != nil {
			return err
		}

		fmt.Println()

		password := strings.TrimSpace(
			string(passwordBytes),
		)

		// authenticate
		user, err := deps.AuthService.Login(
			username,
			password,
		)
		if err != nil {
			return err
		}

		// MFA verification
		if user.MFAEnabled {

			if user.MFASecret == nil {
				return fmt.Errorf(
					"mfa enabled but secret missing",
				)
			}

			fmt.Print("OTP: ")

			otp, err := reader.ReadString('\n')
			if err != nil {
				return err
			}

			otp = strings.TrimSpace(otp)

			valid := deps.TOTPService.ValidateCode(
				otp,
				*user.MFASecret,
			)

			if !valid {
				return fmt.Errorf(
					"invalid otp code",
				)
			}
		}

		// create session
		session, err := deps.SessionManager.CreateSession(
			user.ID,
		)
		if err != nil {
			return err
		}

		// save into runtime context
		deps.CLIContext.Login(
			user,
			session,
		)

		fmt.Println()
		fmt.Println("✓ Login successful")
		fmt.Println()

		fmt.Println("User Details")
		fmt.Println("------------")
		fmt.Printf("Username: %s\n", user.Username)
		fmt.Printf("MFA Enabled: %v\n", user.MFAEnabled)
		fmt.Printf("Registered At: %s\n",
			user.CreatedAt.Format("2006-01-02 15:04:05"),
		)

		if user.LastLogin != nil {
			fmt.Printf("Last Login: %s\n",
				user.LastLogin.Format("2006-01-02 15:04:05"),
			)
		}

		fmt.Printf(
			"Session Expires: %s\n",
			session.ExpiresAt.Format(
				"2006-01-02 15:04:05",
			),
		)

		return nil
	},
}
