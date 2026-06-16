package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var RegisteredCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	RunE: func(cmd *cobra.Command, args []string) error {

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Username: ")
		username, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		username = strings.TrimSpace(username)

		fmt.Print("Password: ")
		passwordBytes, err := term.ReadPassword(
			int(os.Stdin.Fd()),
		)
		fmt.Println()

		password := strings.TrimSpace(string(passwordBytes))

		if err := deps.AuthService.Register(username, password); err != nil {
			return fmt.Errorf("registration failed: %w", err)
		}

		fmt.Println()
		fmt.Println("✓ User registered successfully")

		return nil
	},
}
