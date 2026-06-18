package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var RegisterCmd = &cobra.Command{
	Use:   "register <username> <password>",
	Short: "Register a new user",
	Args:  cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {

		username := strings.TrimSpace(args[0])
		password := strings.TrimSpace(args[1])

		if err := deps.AuthService.Register(
			username,
			password,
		); err != nil {
			return fmt.Errorf(
				"registration failed: %w",
				err,
			)
		}

		fmt.Println("✓ User registered successfully")

		return nil
	},
}
